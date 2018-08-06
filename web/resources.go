package web

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServeFileOnRoute(source, route string) {
	http.HandleFunc(ServeFile(source, route))
}

func ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// Resource represents the idea of a file coming from a directory, as in
// dir/source-name and routed to the given prefix/source-name using the source
// name for find the file on disk and for registering a route.
type Resource struct {
	FromDir     string
	SourceName  string
	RoutePrefix string
}

// Routing for a Resource creates the pair (dir/name, prefix/name).
func (a Resource) Routing() (source, route string) {
	home := os.Getenv("HOME")
	// replace ~ with the home dir
	dir := strings.Replace(a.FromDir, "~", home, 1)
	source = filepath.Join(dir, a.SourceName)
	route = filepath.Join(a.RoutePrefix, a.SourceName)

	return source, route
}

// FileAsset is the simplest form of an Asset.
type FileAsset struct {
	File  string
	Route string
}

// Routing specifically the source file to the route explicitly.
func (x FileAsset) Routing() (source, route string) {
	return x.File, x.Route
}

// Asset is a static file.  The file is served up via a route and that route
// can be from any source, hence the results of Routing provide those two
// values.
type Asset interface {
	Routing() (source, route string)
}

// ServeAssets serves those assets as provided.
func ServeAssets(assets ...Asset) {
	for _, a := range assets {
		ServeFileOnRoute(a.Routing())
	}
}
