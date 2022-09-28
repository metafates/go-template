package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/metafates/go-template/color"
	"github.com/metafates/go-template/constant"
	"github.com/metafates/go-template/filesystem"
	"github.com/metafates/go-template/icon"
	"github.com/metafates/go-template/log"
	"github.com/metafates/go-template/where"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     strings.ToLower(constant.App),
	Short:   "App description",
	Version: constant.Version,
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:       rootCmd,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		Commands:      cc.HiYellow + cc.Bold,
		Example:       cc.Italic,
		ExecName:      cc.Bold,
		Flags:         cc.Bold,
		FlagsDataType: cc.Italic + cc.HiBlue,
	})

	// Clears temp files on each run.
	// It should not affect startup time since it's being run in parallel.
	go func() {
		_ = filesystem.Api().RemoveAll(where.Temp())
	}()

	_ = rootCmd.Execute()
}

func handleErr(err error) {
	if err != nil {
		log.Error(err)
		_, _ = fmt.Fprintf(
			os.Stderr,
			"%s %s\n",
			lipgloss.NewStyle().Foreground(color.Red).Render(icon.Cross()),
			strings.Trim(err.Error(), " \n"))
		os.Exit(1)
	}
}
