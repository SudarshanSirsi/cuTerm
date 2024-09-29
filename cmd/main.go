package main

import (
  "fmt"
  "os"

  tea "github.com/charmbracelet/bubbletea"
  "github.com/SudarshanSirsi/cuTerm/internal/ui"
)

func main() {
  p := tea.NewProgram(ui.InitialModel(), tea.WithAltScreen())
  if err := p.Start(); err != nil {
    fmt.Printf("Error: %v", err)
    os.Exit(1)
  }
} 
