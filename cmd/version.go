package cmd

import (
	"fmt"
	"github.com/metafates/go-template/constant"
	"github.com/spf13/cobra"
	"runtime"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the " + constant.App,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s version %s %s/%s\n", constant.App, constant.Version, runtime.GOOS, runtime.GOARCH)
	},
}
