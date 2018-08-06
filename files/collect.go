package files

import (
	"os"
	"path/filepath"
)

// Collect walks the directory starting at root and collects the names of files.
func Collect(root string) ([]string, error) {
	files := make([]string, 0)

	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // skip errors
		}
		if info.IsDir() {
			return nil // skip dirs
		}
		if info.Name() == root {
			return nil // skip root folder
		}
		files = append(files, path) // collect files
		return nil
	}

	err := filepath.Walk(root, walk)
	if err != nil {
		return nil, err
	}

	return files, nil
}
