package where

import (
	"github.com/metafates/go-template/filesystem"
	"github.com/samber/lo"
	"os"
)

// mkdir creates a directory and all parent directories if they don't exist
// will return the path of the directory
func mkdir(path string) string {
	lo.Must0(filesystem.Api().MkdirAll(path, os.ModePerm))
	return path
}
