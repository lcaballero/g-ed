package hub

type Result struct {
	Name string `json:"name"`
	Op   string `json:"op"`
	Time int64  `json:"time"`
}
