package hub

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"time"
)

type Connection struct {
	conn   *websocket.Conn
	closed bool
}

func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		conn: conn,
	}
}

func (c *Connection) Setup() {
	c.conn.SetReadLimit(MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(PongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(PongWait))
		return nil
	})
}

func (c *Connection) Close() error {
	if c.closed {
		return nil
	}
	err := c.conn.Close()
	c.closed = true
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) IsClosed() bool {
	return c.closed
}

func (c *Connection) WriteMessage(msg []byte) {
	c.conn.WriteMessage(websocket.CloseMessage, msg)
}

func (c *Connection) ReadMessage() (messageType int, p []byte, err error) {
	return c.conn.ReadMessage()
}

func (c *Connection) SetWriteDeadline() error {
	return c.conn.SetWriteDeadline(c.nextDeadline())
}

func (c *Connection) NextWriter() (io.WriteCloser, error) {
	w, err := c.conn.NextWriter(websocket.TextMessage)
	return w, err
}

func (c *Connection) SendPing() error {
	c.SetWriteDeadline()
	err := c.conn.WriteMessage(websocket.PingMessage, empty)
	if err != nil {
		fmt.Println("write message error", err)
		return err
	}
	return nil
}

func (c *Connection) nextDeadline() time.Time {
	deadline := time.Now().Add(WriteWait)
	return deadline
}
