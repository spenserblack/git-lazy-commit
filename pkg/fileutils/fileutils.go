// Package file provides utilities for working with files.
package fileutils

import (
	"path/filepath"
	"strings"
)

// SharedDirectory returns the shared directory of the given paths.
// If there is no shared directory, "/" is returned.
func SharedDirectory(paths []string) string {
	if len(paths) == 0 {
		return ""
	}
	if len(paths) == 1 {
		return filepath.Dir(paths[0])
	}
	dirs := make([][]string, len(paths))
	for i, path := range paths {
		cleaned := filepath.ToSlash(path)
		dirs[i] = strings.Split(cleaned, "/")
	}

	shared := ""
	SharedDirLoop:
		for i := 0; i < len(dirs[0]); i++ {
			dir := dirs[0][i]
			for _, d := range dirs[1:] {
				if i >= len(d) || d[i] != dir {
					break SharedDirLoop
				}
			}
			shared = filepath.Join(shared, dir)
		}

	if shared != "" {
		return shared
	}

	return "/"
}
