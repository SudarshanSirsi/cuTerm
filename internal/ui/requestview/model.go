package requestview

import (
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

type Model struct {
  Method string
  URL    string
}

func New() Model {
  return Model{
    Method: "GET",
    URL:    "https://api.example.com",
  }
}

func (m Model) Init() tea.Cmd {
  return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return m, nil
}

func (m Model) View() string {
  return lipgloss.NewStyle().Margin(1).Render(
    "Request View\n\n" +
    "Method: " + m.Method + "\n" +
    "URL: " + m.URL,
    )
}
