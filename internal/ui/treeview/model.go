package treeview

import (
  "fmt"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

type Item struct {
  Title    string
  Children []*Item
}

type Model struct {
  Items    []*Item
  Cursor   int
  Selected map[int]struct{}
}

func New() Model {
  return Model{
    Items: []*Item{
      {Title: "Collection 1", Children: []*Item{{Title: "Request 1"}, {Title: "Request 2"}}},
      {Title: "Collection 2", Children: []*Item{{Title: "Request 3"}}},
    },
    Selected: make(map[int]struct{}),
  }
}

func (m Model) Init() tea.Cmd {
  return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "up", "k":
      if m.Cursor > 0 {
        m.Cursor--
      }
    case "down", "j":
      m.Cursor++
    }
  }
  return m, nil
}

func (m Model) View() string {
  s := "Tree View\n\n"
  for i, item := range m.Items {
    cursor := " "
    if m.Cursor == i {
      cursor = ">"
    }
    s += fmt.Sprintf("%s %s\n", cursor, item.Title)
    for _, child := range item.Children {
      s += fmt.Sprintf("  - %s\n", child.Title)
    }
  }
  return lipgloss.NewStyle().Margin(1, 0, 0, 2).Render(s)
}
