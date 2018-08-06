package hub

import "encoding/json"

type ParsedReq struct {
	Method  string            `json:"method"`
	Path    string            `json:"path"`
	Headers map[string]string `json:"headers"`
	Params  map[string]string `json:"-"` // pulled from the Path: '?key=value'
	Data    interface{}       `json:"data"`
}

func LoadRequest(bin []byte) (Req, error) {
	req := ParsedReq{}
	err := json.Unmarshal(bin, &req)
	if err != nil {
		return Req{}, err
	}
	req.Params = ParseQueryParams(req.Path)
	req.Path = ParsePath(req.Path)
	return Req{req: req}, nil
}
