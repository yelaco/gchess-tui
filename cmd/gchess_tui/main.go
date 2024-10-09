package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/tui/screens"
)

func main() {
	if _, err := tea.NewProgram(screens.RootScreen(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal("could not start program: ", err)
	}
}
