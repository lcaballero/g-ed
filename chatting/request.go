package chatting

type Request interface {
	Data() interface{}
	Method() string
	Path() string
	Headers() map[string]string
	Params() map[string]string
	ToJson() []byte
}
