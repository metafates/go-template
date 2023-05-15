package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/go-template/color"
)

var (
	Success = lipgloss.NewStyle().Foreground(color.Green).Render
	Failure = lipgloss.NewStyle().Foreground(color.Red).Render
	Warning = lipgloss.NewStyle().Foreground(color.Yellow).Render
)
