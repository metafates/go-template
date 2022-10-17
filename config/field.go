package config

import (
	"fmt"
	"github.com/metafates/go-template/constant"
	"github.com/metafates/go-template/style"
	"reflect"
	"strings"

	"github.com/metafates/go-template/color"
	"github.com/spf13/viper"
)

type Field struct {
	Key         string
	Value       any
	Description string
}

func (f Field) Pretty() string {
	return fmt.Sprintf(
		`%s
%s: %s = %s
`,
		style.Faint(f.Description),
		style.Fg(color.Purple)(f.Key),
		style.Fg(color.Yellow)(reflect.TypeOf(f.Value).String()),
		style.Fg(color.Cyan)(fmt.Sprintf("%v", viper.Get(f.Key))),
	)
}

func (f Field) Env() string {
	return strings.ToUpper(constant.App + "_" + EnvKeyReplacer.Replace(f.Key))
}
