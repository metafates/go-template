package config

import (
	"github.com/metafates/go-template/constant"
	"github.com/metafates/go-template/filesystem"
	"github.com/metafates/go-template/where"
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName(constant.App)
	viper.SetConfigType(constant.ConfigFormat)
	viper.SetFs(filesystem.Api())
	viper.AddConfigPath(where.Config())

	setDefaults()
	setEnvs()

	err := viper.ReadInConfig()

	switch err.(type) {
	case viper.ConfigFileNotFoundError:
		// Use defaults then
		return nil
	default:
		return err
	}
}
