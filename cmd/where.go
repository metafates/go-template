package cmd

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/go-template/color"
	"github.com/metafates/go-template/constant"
	"github.com/metafates/go-template/util"
	"github.com/metafates/go-template/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
)

var wherePaths = []lo.Tuple2[string, func() string]{
	{"config", where.Config},
	{"logs", where.Logs},
}

func init() {
	rootCmd.AddCommand(whereCmd)

	for _, n := range wherePaths {
		whereCmd.Flags().BoolP(n.A, string(n.A[0]), false, n.A+" path")
	}

	whereCmd.MarkFlagsMutuallyExclusive(lo.Map(wherePaths, func(t lo.Tuple2[string, func() string], _ int) string {
		return t.A
	})...)

	whereCmd.SetOut(os.Stdout)
}

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Show the paths for a files related to the " + constant.App,
	Run: func(cmd *cobra.Command, args []string) {
		headerStyle := lipgloss.NewStyle().Bold(true).Foreground(color.HiPurple).Render
		yellowStyle := lipgloss.NewStyle().Foreground(color.Yellow).Render

		for _, n := range wherePaths {
			if lo.Must(cmd.Flags().GetBool(n.A)) {
				cmd.Println(n.B())
				return
			}
		}

		for i, n := range wherePaths {
			cmd.Printf("%s %s\n", headerStyle(util.Capitalize(n.A)+"?"), yellowStyle("--"+n.A))
			cmd.Println(n.B())

			if i < len(wherePaths)-1 {
				cmd.Println()
			}
		}
	},
}
