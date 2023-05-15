package config

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"reflect"
	"strconv"
	"strings"
	"text/template"

	"github.com/metafates/go-template/app"
	"github.com/samber/lo"
	"github.com/spf13/viper"

	"github.com/metafates/go-template/color"
)

type Field struct {
	Key          string
	DefaultValue any
	Description  string
}

// typeName returns the type of the field without reflection
func (f *Field) typeName() string {
	switch f.DefaultValue.(type) {
	case string:
		return "string"
	case int:
		return "int"
	case bool:
		return "bool"
	case []string:
		return "[]string"
	case []int:
		return "[]int"
	default:
		return "unknown"
	}
}

var prettyTemplate = lo.Must(template.New("pretty").Funcs(template.FuncMap{
	"faint":  lipgloss.NewStyle().Faint(true).Render,
	"bold":   lipgloss.NewStyle().Bold(true).Render,
	"purple": lipgloss.NewStyle().Foreground(color.Purple).Render,
	"blue":   lipgloss.NewStyle().Foreground(color.Blue).Render,
	"cyan":   lipgloss.NewStyle().Foreground(color.Cyan).Render,
	"value":  func(k string) any { return viper.Get(k) },
	"hl": func(v any) string {
		switch value := v.(type) {
		case bool:
			b := strconv.FormatBool(value)
			if value {
				return lipgloss.NewStyle().Foreground(color.Green).Render(b)
			}

			return lipgloss.NewStyle().Foreground(color.Red).Render(b)
		case string:
			return lipgloss.NewStyle().Foreground(color.Yellow).Render(value)
		default:
			return fmt.Sprint(value)
		}
	},
	"typename": func(v any) string { return reflect.TypeOf(v).String() },
}).Parse(`{{ faint .Description }}
{{ blue "Key:" }}     {{ purple .Key }}
{{ blue "Env:" }}     {{ .Env }}
{{ blue "Value:" }}   {{ hl (value .Key) }}
{{ blue "Default:" }} {{ hl (.DefaultValue) }}
{{ blue "Type:" }}    {{ typename .DefaultValue }}`))

func (f *Field) Pretty() string {
	var b strings.Builder

	lo.Must0(prettyTemplate.Execute(&b, f))

	return b.String()
}

func (f *Field) Env() string {
	env := strings.ToUpper(EnvKeyReplacer.Replace(f.Key))
	appPrefix := strings.ToUpper(app.Name + "_")

	if strings.HasPrefix(env, appPrefix) {
		return env
	}

	return appPrefix + env
}
