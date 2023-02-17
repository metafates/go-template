package cmd

import (
	"fmt"
	"os"
	"strings"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/metafates/go-template/app"
	"github.com/metafates/go-template/color"
	"github.com/metafates/go-template/filesystem"
	"github.com/metafates/go-template/icon"
	"github.com/metafates/go-template/log"
	"github.com/metafates/go-template/style"
	"github.com/metafates/go-template/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().BoolP("version", "v", false, app.Name+" version")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   strings.ToLower(app.Name),
	Short: app.DescriptionShort,
	Long:  app.DescriptionLong,
	Run: func(cmd *cobra.Command, args []string) {
		if lo.Must(cmd.Flags().GetBool("version")) {
			versionCmd.Run(versionCmd, args)
		}
	},
}

func Execute() {
	// Setup colored cobra
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})

	// Clears temp files on each run.
	// It should not affect startup time since it's being run as goroutine.
	go func() {
		_ = filesystem.Api().RemoveAll(where.Temp())
	}()

	_ = rootCmd.Execute()
}

// handleErr will stop program execution and log error to the stderr
// if err is not nil
func handleErr(err error) {
	if err == nil {
		return
	}

	log.Error(err)
	_, _ = fmt.Fprintf(
		os.Stderr,
		"%s %s\n",
		style.Fg(color.Red)(icon.Cross),
		strings.Trim(err.Error(), " \n"),
	)
	os.Exit(1)
}
