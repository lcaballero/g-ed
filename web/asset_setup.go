package web

import (
	"path/filepath"
)

func SetupAssets(root string, dirs ...string) error {
	for _, d := range dirs {
		base := filepath.Base(d)
		dirPath := filepath.Dir(d)

		switch base {
		case "index.html":
			ServeAssets(FileAsset{File: d, Route: "/"})
		default:
			res := Resource{
				FromDir:     dirPath,
				SourceName:  base,
				RoutePrefix: dirPath[len(root):],
			}
			if res.RoutePrefix == "" {
				res.RoutePrefix = "/"
			}
			ServeAssets(res)
		}
	}

	return nil
}
