package config

import (
	"github.com/metafates/go-template/constant"
	"github.com/spf13/viper"
)

type field struct {
	key         string
	value       any
	description string
}

var fields = []field{
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
	Default = make(map[string]any)
	for _, f := range fields {
		Default[f.key] = f.value
		viper.SetDefault(f.key, f.value)
		EnvExposed = append(EnvExposed, f.key)
	}
}

var Default map[string]any
