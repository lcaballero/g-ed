package web

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func ServeFile(path, route string) (string, Handler) {
	return route, func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		bin, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		ext := filepath.Ext(path)

		switch ext {
		case ".css":
			w.Header().Add("content-type", "text/css")
		case ".js":
			w.Header().Add("content-type", "application/javascript")
		}

		w.Write(bin)
	}
}
