package responseview

import (
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

type Model struct {
  StatusCode int
  Body       string
}

func New() Model {
  return Model{
    StatusCode: 200,
    Body:       "Response will appear here",
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
    "Response View\n\n" +
    "Status: " + lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Render(m.Status()) + "\n" +
    "Body: " + m.Body,
    )
}

func (m Model) Status() string {
  return fmt.Sprintf("%d OK", m.StatusCode)
}
