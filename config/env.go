package config

import (
	"github.com/metafates/go-template/constant"
	"github.com/spf13/viper"
	"strings"
)

var (
	EnvKeyReplacer = strings.NewReplacer(".", "_")
	EnvExposed     []string
)

// setEnvs sets the environment variables
func setEnvs() {
	viper.SetEnvPrefix(constant.App)
	viper.SetEnvKeyReplacer(EnvKeyReplacer)

	for _, env := range EnvExposed {
		viper.MustBindEnv(env)
	}
}
