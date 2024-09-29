package ui

import (
  "github.com/charmbracelet/bubbles/help"
  "github.com/charmbracelet/bubbles/key"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
  "github.com/SudarshanSirsi/cuTerm/internal/ui/treeview"
  "github.com/SudarshanSirsi/cuTerm/internal/ui/requestview"
  "github.com/SudarshanSirsi/cuTerm/internal/ui/responseview"
)

type Model struct {
  treeView     treeview.Model
  requestView  requestview.Model
  responseView responseview.Model
  help         help.Model
  keys         keyMap
  activeView   string
}

type keyMap struct {
  Up    key.Binding
  Down  key.Binding
  Left  key.Binding
  Right key.Binding
  Help  key.Binding
  Quit  key.Binding
}

func (km keyMap) FullHelp() [][]key.Binding {
  return [][]key.Binding{
    {km.Up, km.Down, km.Left, km.Right},
    {km.Help, km.Quit},
  }
}

// Implementing ShortHelp() to satisfy help.KeyMap interface
func (km keyMap) ShortHelp() []key.Binding {
  return []key.Binding{
    km.Up, km.Down, km.Help, km.Quit,
  }
}

func InitialModel() Model {
  return Model{
    treeView:     treeview.New(),
    requestView:  requestview.New(),
    responseView: responseview.New(),
    help:         help.New(),
    keys:         initKeyMap(),
    activeView:   "tree",
  }
}

func initKeyMap() keyMap {
  return keyMap{
    Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "move up")),
    Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "move down")),
    Left:  key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "move left")),
    Right: key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "move right")),
    Help:  key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "toggle help")),
    Quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
  }
}

func (m Model) Init() tea.Cmd {
  return nil
}


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmds []tea.Cmd

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch {
    case key.Matches(msg, m.keys.Quit):
      return m, tea.Quit
    case key.Matches(msg, m.keys.Help):
      m.help.ShowAll = !m.help.ShowAll
    case key.Matches(msg, m.keys.Left):
      m.activeView = "tree"
    case key.Matches(msg, m.keys.Right):
      if m.activeView == "tree" {
        m.activeView = "request"
      } else if m.activeView == "request" {
        m.activeView = "response"
      }
    }
  }

  // Use tea.Batch to handle multiple cmds, instead of declaring and not using
  switch m.activeView {
  case "tree":
    treeModel, cmd := m.treeView.Update(msg)
    m.treeView = treeModel.(treeview.Model)
    cmds = append(cmds, cmd)
  case "request":
    requestModel, cmd := m.requestView.Update(msg)
    m.requestView = requestModel.(requestview.Model)
    cmds = append(cmds, cmd)
  case "response":
    responseModel, cmd := m.responseView.Update(msg)
    m.responseView = responseModel.(responseview.Model)
    cmds = append(cmds, cmd)
  }

  return m, tea.Batch(cmds...)
}


func (m Model) View() string {
  treeView := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("62")).
    Padding(1).
    Render(m.treeView.View())

  requestView := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("62")).
    Padding(1).
    Render(m.requestView.View())

  responseView := lipgloss.NewStyle().
    Border(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("62")).
    Padding(1).
    Render(m.responseView.View())

  mainView := lipgloss.JoinHorizontal(
    lipgloss.Left,
    lipgloss.NewStyle().Width(40).Render(treeView),
    lipgloss.NewStyle().Width(80).Render(
      lipgloss.JoinVertical(
        lipgloss.Left,
        lipgloss.NewStyle().Height(15).Render(requestView),
        responseView,
        ),
      ),
    )

  return lipgloss.JoinVertical(
    lipgloss.Left,
    mainView,
    m.help.View(m.keys),
    )
}
