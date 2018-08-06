package hub

type Relay struct {
	client *Client
	hub    *Hub
}

func (r Relay) Id() int {
	return r.client.Id()
}

func (r Relay) Broadcast(msg []byte) error {
	return nil
}

func (r Relay) Send(msg []byte) error {
	return r.client.Send(msg)
}
