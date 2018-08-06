package chatting

type Chatter interface {
	Id() int
	Send(msg []byte) error
	Broadcast(msg []byte) error
}
