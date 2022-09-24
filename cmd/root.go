package cmd

import (
	"github.com/metafates/go-template/constant"
	"github.com/spf13/cobra"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   strings.ToLower(constant.App),
	Short: "App description",
}

func Execute() {
	_ = rootCmd.Execute()
}
