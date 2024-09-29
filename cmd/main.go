package main

import (
  "fmt"
  "os"

  tea "github.com/charmbracelet/bubbletea"
  "github.com/SudarshanSirsi/cuTerm/internal/ui"
)

func main() {
  Program := tea.NewProgram(ui.InitialModel(), tea.WithAltScreen())
  if _, err := Program.Run(); err != nil {
    fmt.Printf("Error: %v", err)
    os.Exit(1)
  }
} 
