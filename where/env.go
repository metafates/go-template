package where

import (
	"strings"

	"github.com/metafates/go-template/app"
)

// EnvConfigPath is the environment variable name for the config path
var EnvConfigPath = strings.ToUpper(app.Name) + "_CONFIG_PATH"
