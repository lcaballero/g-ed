package hub

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lcaballero/g-ed/chatting"
	"sync"
)

var ErrSendToClosedClient = errors.New("error sending to closed client")
var ErrSendToStoppedClient = errors.New("error sending to closed client")
var ErrBlockedSend = errors.New("error send channel blocked")
var empty = []byte{}

const (
	WriteWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	PongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	PingPeriod     = (PongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	MaxMessageSize = 1024 * 1024         // Maximum message size allowed from peer.
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	id         int             // used to register and unregister from the hub.
	hub        *Hub            // the hub holding clients.
	conn       *Connection     // the websocket connection.
	send       chan []byte     // buffered channel of outbound messages.
	stopped    bool            // true if client has fully shutdown
	stop       chan bool       // filled when pump funcs are exited
	killed     bool            // true if shutdown is in progress
	killSignal chan bool       // a broadcast channel when channel is closed
	kg         *sync.WaitGroup // WaitGroup blocking while client is running
}

// NewClient creates a Client.
func NewClient(
	hub *Hub,
	upgrader websocket.Upgrader,
	w http.ResponseWriter, r *http.Request) (*Client, error) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		id:         NewId(),
		hub:        hub,
		conn:       NewConnection(conn),
		send:       make(chan []byte, 256),
		killSignal: make(chan bool),
		stop:       make(chan bool, 3), // 3 because 3 pumps
		kg:         &sync.WaitGroup{},
	}
	client.kg.Add(cap(client.stop))

	fmt.Printf("new client id: %d\n", client.id)

	return client, nil
}

// Id provides the unique Id of the client.
func (c *Client) Id() int {
	return c.id
}

// Send attempts to ship the message to the web client via the write pump.
func (c *Client) Send(msg []byte) error {
	if c.stopped {
		return ErrSendToStoppedClient
	}
	if c.conn.IsClosed() {
		return ErrSendToClosedClient
	}

	select {
	case c.send <- msg:
		return nil
	default:
		return ErrBlockedSend // message dropped
	}
}

// ServeWs handles websocket requests from the peer.
func (client *Client) ServeWs() {
	fmt.Println("creating new client")
	go client.writePump()
	go client.readPump()
	go client.tic()
}

// Close shuts down the internal read/write loops.  Once all of them are
// shutdown the underlying web-socket is closed.  While the client is shutdown
// sending messages to it will return errors from Send(...).
func (c *Client) Close() error {
	if c.stopped || c.killed {
		fmt.Println("client is stopped or is being stopped")
		return nil
	}

	fmt.Println("closing client")
	c.killed = true
	n := cap(c.stop)
	close(c.killSignal) // broadcast kill, causing loop exits
	c.hub.Unregister(c)

	for {
		select {
		case <-c.stop:
			fmt.Println("c.wait ed", n)
			if n == 0 {
				err := c.closeWebSocket()
				c.stopped = true
				fmt.Println("close()")
				return err
			}
			n--
		}
	}

	return errors.New("unreachable")
}

// Wait blocks and waits for the web-socket connection to be closed.
// Once unblocked the client should be closed.
func (c *Client) Wait() {
	c.kg.Wait()
	fmt.Println("client done waiting")
}

// closeWebSocket closes the underlying connection and returns the error.
func (c *Client) closeWebSocket() error {
	return c.conn.Close()
}

// done calls Done on the underlying wait group setup to wait for each
// read/write pump to exit.  A channel is filled with signals from these each
// loop so that once close is called it can correctly shutdown.  This is required
// because the client needs to shutdown from outside calls, but also when a
// close signal from the web-client is issued.
func (c *Client) done() {
	c.kg.Done()
	c.stop <- true
	c.Close()
}

// tic starts a write pump that pushes heartbeat like updates to the client.
func (c *Client) tic() {
	tic := time.NewTicker(1 * time.Second)
	fmt.Println("starting ticker pump")

	defer func() {
		fmt.Println("done ticking")
		tic.Stop()
		c.done()
	}()

	r := 1
	res := map[string]interface{}{"type": "ping"}
	ping, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-c.killSignal:
			return

		case <-tic.C:
			fmt.Println("ticker fired")
			c.send <- ping
			r++
		}
	}
}

// isUnexpectedClose checks to see if the error is CloseGoingAway.
func (c *Client) isUnexpectClose(err error) bool {
	return websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway)
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		fmt.Println("done read pump")
		c.done()
	}()

	fmt.Println("starting read pump")
	c.conn.Setup()
	h := chatting.Handler{}

	for {
		select {
		case <-c.killSignal:
			return

		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				if c.isUnexpectClose(err) {
					log.Printf("error: %v", err)
				} else {
					fmt.Printf("read message: %s\n", err)
				}
				return
			}

			message = bytes.TrimSpace(message)
			request, err := LoadRequest(message)

			if err != nil {
				fmt.Printf("bad message: '%s'\n'%s'", string(message), err)
			} else {
				h.Handle(Relay{client: c, hub: c.hub}, request)
			}
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		fmt.Println("done write pump")
		ticker.Stop()
		c.done()
	}()

	fmt.Println("starting write pump")

	for {
		select {

		case <-c.killSignal:
			return

		case message, ok := <-c.send:
			err := c.conn.SetWriteDeadline()

			if !ok || err != nil {
				fmt.Println("send channel closed", ok, err)
				c.conn.WriteMessage(empty)
				return
			}

			w, err := c.conn.NextWriter()
			if err != nil {
				fmt.Println("next writer error:", err)
				return
			}

			w.Write(message)

			if err := w.Close(); err != nil {
				fmt.Println("error closing writer??", err)
				return
			}

		case <-ticker.C:
			c.conn.SendPing()
		}
	}
}
