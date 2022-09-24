package filesystem

import "github.com/spf13/afero"

// SetOsFs sets the filesystem to the os filesystem
func SetOsFs() {
	if wrapper.Fs == nil || wrapper.Fs.Name() != "os" {
		wrapper.Fs = afero.NewOsFs()
	}
}

// SetMemMapFs sets the filesystem to the memory mapped filesystem
// Use this if you want to use the filesystem in a sandbox
func SetMemMapFs() {
	if wrapper.Fs == nil || wrapper.Fs.Name() != "memmap" {
		wrapper.Fs = afero.NewMemMapFs()
	}
}
