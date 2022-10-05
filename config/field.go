package config

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/go-template/color"
	"reflect"
	"strings"
)

type Field struct {
	Key         string
	Value       any
	Description string
}

func (f Field) Pretty() string {
	return fmt.Sprintf(
		`%s
%s: %s %s
`,
		lipgloss.NewStyle().Faint(true).Render(f.Description),
		lipgloss.NewStyle().Foreground(color.Purple).Render(f.Key),
		lipgloss.NewStyle().Foreground(color.Yellow).Render(reflect.TypeOf(f.Value).String()),
		lipgloss.NewStyle().Faint(true).Render("default: "+fmt.Sprintf("%v", f.Value)),
	)
}

func (f Field) Env() string {
	return strings.ToUpper(EnvPrefix + EnvKeyReplacer.Replace(f.Key))
}
