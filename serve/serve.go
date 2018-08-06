package serve

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lcaballero/g-ed/files"
	"github.com/lcaballero/g-ed/hub"
	"github.com/lcaballero/g-ed/params"
	"github.com/lcaballero/g-ed/web"
	"net/http"
)

// Serve takes the command-line parameters form which it loads ServeParams and
// uses the Root from which to serve assets.
func Serve(val params.ValContext) error {
	parms := params.Load(val)

	dirs, err := files.Collect(parms.Root)
	if err != nil {
		return err
	}

	fmt.Println(dirs)

	err = web.SetupAssets(parms.Root, dirs...)
	if err != nil {
		return err
	}

	hub := hub.NewHub()
	hub.Run()

	ip := ":3000"
	fmt.Printf("binding web server to port %s\n", ip)

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024 * 1024,
		WriteBufferSize: 1024 * 1024,
		CheckOrigin: func(*http.Request) bool {
			return true
		},
	}

	web.HandleFunc("/websocket", webSocket(upgrader, hub))

	err = web.ListenAndServe(ip, nil)
	if err != nil {
		return err
	}

	return nil
}

func webSocket(
	upgrader websocket.Upgrader,
	clientHub *hub.Hub) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("handling websocket")

		client, err := hub.NewClient(clientHub, upgrader, res, req)
		if err != nil {
			return
		}
		clientHub.Register(client)

		client.ServeWs()
		client.Wait()
		fmt.Println("closing web socket")
		client.Close()
	}
}
