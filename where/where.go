package where

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/metafates/go-template/app"
)

func home() string {
	home, err := os.UserHomeDir()
	if err == nil {
		return home
	}

	return "."
}

// Config path
// Will create the directory if it doesn't exist
func Config() string {
	var path string

	if customDir, present := os.LookupEnv(EnvConfigPath); present {
		return mkdir(customDir)
	}

	var userConfigDir string

	if runtime.GOOS == "darwin" {
		userConfigDir = filepath.Join(home(), ".config")
	} else {
		var err error
		userConfigDir, err = os.UserConfigDir()
		if err != nil {
			userConfigDir = filepath.Join(home(), ".config")
		}
	}

	path = filepath.Join(userConfigDir, app.Name)
	return mkdir(path)
}

// Logs path
// Will create the directory if it doesn't exist
func Logs() string {
	return mkdir(filepath.Join(Cache(), "logs"))
}

// Cache path
// Will create the directory if it doesn't exist
func Cache() string {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		userCacheDir = "."
	}

	cacheDir := filepath.Join(userCacheDir, app.Name)
	return mkdir(cacheDir)
}

// Temp path
// Will create the directory if it doesn't exist
func Temp() string {
	tempDir := filepath.Join(os.TempDir(), app.Name)
	return mkdir(tempDir)
}
