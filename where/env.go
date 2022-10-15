package where

import (
	"github.com/metafates/go-template/constant"
	"strings"
)

// EnvConfigPath is the environment variable name for the config path
var EnvConfigPath = strings.ToUpper(constant.App) + "_CONFIG_PATH"
