package config

import (
	"fmt"
	"github.com/metafates/go-template/constant"
	"github.com/metafates/go-template/style"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"reflect"
	"strings"
	"text/template"

	"github.com/metafates/go-template/color"
)

type Field struct {
	Key          string
	DefaultValue any
	Description  string
}

var prettyTemplate = lo.Must(template.New("pretty").Funcs(template.FuncMap{
	"faint":    style.Faint,
	"bold":     style.Bold,
	"purple":   style.Fg(color.Purple),
	"yellow":   style.Fg(color.Yellow),
	"cyan":     style.Fg(color.Cyan),
	"value":    func(k string) string { return fmt.Sprint(viper.Get(k)) },
	"typename": func(v any) string { return reflect.TypeOf(v).String() },
}).Parse(`{{ faint .Description }}
{{ yellow "Key:" }}     {{ purple .Key }}
{{ yellow "Env:" }}     {{ .Env }}
{{ yellow "Value:" }}   {{ value .Key }}
{{ yellow "Default:" }} {{ .DefaultValue }}
{{ yellow "Type:" }}    {{ typename .DefaultValue }}`))

func (f *Field) Pretty() string {
	var b strings.Builder

	lo.Must0(prettyTemplate.Execute(&b, f))

	return b.String()
}

func (f *Field) Env() string {
	env := strings.ToUpper(EnvKeyReplacer.Replace(f.Key))
	appPrefix := strings.ToUpper(constant.App + "_")

	if strings.HasPrefix(env, appPrefix) {
		return env
	}

	return appPrefix + env
}
