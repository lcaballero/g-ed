package hub

import (
	"strings"
)

func ParsePath(qs string) string {
	index := strings.Index(qs, "?")
	path := qs
	if index >= 0 {
		path = qs[:index]
	}
	return path
}

func ParseQueryParams(qs string) map[string]string {
	index := strings.Index(qs, "?")
	params := ""
	if index >= 0 {
		params = qs[index+1:]
	}

	pairs := strings.Split(params, "&")
	m := make(map[string]string)
	for _, kvp := range pairs {
		nv := strings.Split(kvp, "=")
		if len(nv) == Pair {
			m[nv[0]] = nv[1]
		}
	}
	return m
}
