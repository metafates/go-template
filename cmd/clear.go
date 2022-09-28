package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/go-template/color"
	"github.com/metafates/go-template/constant"
	"github.com/metafates/go-template/filesystem"
	"github.com/metafates/go-template/icon"
	"github.com/metafates/go-template/util"
	"github.com/metafates/go-template/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var clearTargets = []lo.Tuple2[string, func() error]{
	{"cache", func() error {
		return filesystem.Api().RemoveAll(where.Cache())
	}},
	{"logs", func() error {
		return filesystem.Api().RemoveAll(where.Logs())
	}},
}

func init() {
	rootCmd.AddCommand(clearCmd)
	for _, n := range clearTargets {
		clearCmd.Flags().BoolP(n.A, string(n.A[0]), false, "clear "+n.A)
	}
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears sidelined files produced by the " + constant.App,
	Run: func(cmd *cobra.Command, args []string) {
		successStyle := lipgloss.NewStyle().Foreground(color.Green).Render
		var didSomething bool
		for _, n := range clearTargets {
			if lo.Must(cmd.Flags().GetBool(n.A)) {
				handleErr(n.B())
				fmt.Printf("%s %s cleared\n", successStyle(icon.Check), util.Capitalize(n.A))
				didSomething = true
			}
		}

		if !didSomething {
			_ = cmd.Help()
		}
	},
}
