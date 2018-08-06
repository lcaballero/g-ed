package hub

import (
	"encoding/json"
)

const Pair = 2

type Req struct {
	req ParsedReq
}

func (r Req) Data() interface{} {
	return r.req.Data
}

func (r Req) Method() string {
	return r.req.Method
}

func (r Req) Path() string {
	return r.req.Path
}

func (r Req) Headers() map[string]string {
	return r.req.Headers
}

func (r Req) Params() map[string]string {
	return r.req.Params
}

func (r Req) ToJson() []byte {
	bin, err := json.MarshalIndent(&r.req, "", "  ")
	if err != nil {
		return []byte{}
	}
	return bin
}
