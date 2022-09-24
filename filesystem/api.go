package filesystem

import "github.com/spf13/afero"

// Api returns the filesystem api
func Api() afero.Afero {
	return wrapper
}
