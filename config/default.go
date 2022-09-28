package config

import (
	"github.com/metafates/go-template/constant"
	"github.com/spf13/viper"
)

var fields = []Field{
	// LOGS
	{
		constant.LogsWrite,
		false,
		"Write logs to file",
	},
	{
		constant.LogsLevel,
		"info",
		`Logs level.
Available options are: (from less to most verbose)
panic, fatal, error, warn, info, debug, trace`,
	},
	// END LOGS
}

func setDefaults() {
	Default = make(map[string]Field)
	for _, f := range fields {
		Default[f.Key] = f
		viper.SetDefault(f.Key, f.Value)
		viper.MustBindEnv(f.Key)
	}
}

var Default map[string]Field
