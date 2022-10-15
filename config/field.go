package config

import (
	"fmt"
	"github.com/metafates/go-template/constant"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
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
		lipgloss.NewStyle().Faint(true).Render(f.Description),
		lipgloss.NewStyle().Foreground(color.Purple).Render(f.Key),
		lipgloss.NewStyle().Foreground(color.Yellow).Render(reflect.TypeOf(f.Value).String()),
		lipgloss.NewStyle().Foreground(color.Cyan).Render(fmt.Sprintf("%v", viper.Get(f.Key))),
	)
}

func (f Field) Env() string {
	return strings.ToUpper(constant.App + "_" + EnvKeyReplacer.Replace(f.Key))
}
