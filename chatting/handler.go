package chatting

import "fmt"

type Handler struct {
}

type Message struct {
	Text    string `json:"text"`
	Session string `json:"session"`
}

func (h Handler) Handle(c Chatter, request Request) {

	switch request.Path() {
	case "/load":
	case "/join/room-name":
		fmt.Println("joining 'room-name'")
		c.Send([]byte(`{"type": "joined-room"}`))
	case "/speak":
		fmt.Println("broadcasting to all")
		c.Broadcast([]byte(`{"type":"broadcast", "message": "hello"}`))
	case "/message":

	default:
		fmt.Println(string(request.ToJson()))
	}
}
