package cmd

import (
	"github.com/metafates/go-template/color"
	"github.com/metafates/go-template/style"
	"html/template"
	"os"
	"runtime"

	"github.com/metafates/go-template/constant"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.SetOut(os.Stdout)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the " + constant.App,
	Run: func(cmd *cobra.Command, args []string) {
		versionInfo := struct {
			Version  string
			OS       string
			Arch     string
			App      string
			Compiler string
		}{
			Version:  constant.Version,
			App:      constant.App,
			OS:       runtime.GOOS,
			Arch:     runtime.GOARCH,
			Compiler: runtime.Compiler,
		}

		t, err := template.New("version").Funcs(map[string]any{
			"faint":   style.Faint,
			"bold":    style.Bold,
			"magenta": style.Fg(color.Purple),
		}).Parse(`{{ magenta "▇▇▇" }} {{ magenta .App }} 

  {{ faint "Version" }}   {{ bold .Version }}
  {{ faint "Platform" }}  {{ bold .OS }}/{{ bold .Arch }}
  {{ faint "Compiler" }}  {{ bold .Compiler }}
`)
		handleErr(err)
		handleErr(t.Execute(cmd.OutOrStdout(), versionInfo))
	},
}
