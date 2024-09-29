package ui

import "github.com/charmbracelet/lipgloss"

var (
  SelectedStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("170")).
    Bold(true)

  RegularStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("240"))
)
