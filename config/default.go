package config

import (
	"github.com/metafates/go-template/key"
	"github.com/spf13/viper"
)

// fields is the config fields with their default values and descriptions
var fields = []*Field{
	// LOGS
	{
		key.LogsWrite,
		true,
		"Write logs to file",
	},
	{
		key.LogsLevel,
		"info",
		`Logs level.
Available options are: (from less to most verbose)
fatal, error, warn, info, debug`,
	},
	{
		key.LogsReportCaller,
		false,
		"Whether the logger should report the caller location.",
	},
	// END LOGS
}

func setDefaults() {
	Default = make(map[string]*Field, len(fields))
	for _, f := range fields {
		Default[f.Key] = f
		viper.SetDefault(f.Key, f.DefaultValue)
		viper.MustBindEnv(f.Key)
	}
}

var Default map[string]*Field
