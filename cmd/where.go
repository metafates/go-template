package cmd

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/go-template/color"
	"github.com/metafates/go-template/constant"
	"github.com/metafates/go-template/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(whereCmd)

	whereCmd.Flags().BoolP("config", "c", false, "configuration path")
	whereCmd.Flags().BoolP("logs", "l", false, "logs path")
	whereCmd.MarkFlagsMutuallyExclusive("config", "logs")
	whereCmd.SetOut(os.Stdout)
}

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Show the paths for a files related to the " + constant.App,
	Run: func(cmd *cobra.Command, args []string) {

		headerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(color.HiPurple)).Render
		yellowStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(color.Yellow)).Render

		whereConfig := lo.Must(cmd.Flags().GetBool("config"))
		whereLogs := lo.Must(cmd.Flags().GetBool("logs"))

		title := func(do bool, what, arg string) {
			if do {
				cmd.Printf("%s %s\n", headerStyle(what+"?"), yellowStyle(arg))
			}
		}

		printConfigPath := func(header bool) {
			title(header, "Configuration", "--config")
			cmd.Println(where.Config())
		}

		printLogsPath := func(header bool) {
			title(header, "Logs", "--logs")
			cmd.Println(where.Logs())
		}

		switch {
		case whereConfig:
			printConfigPath(false)
		case whereLogs:
			printLogsPath(false)
		default:
			printConfigPath(true)
			cmd.Println()
			printLogsPath(true)
			cmd.Println()
		}
	},
}
