package screens

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
)

var Dump *os.File

type rootScreenModel struct {
	model tea.Model
}

func RootScreen() rootScreenModel {
	return rootScreenModel{
		model: NewLoginScreenModel(),
	}
}

func (m rootScreenModel) Init() tea.Cmd {
	return m.model.Init() // rest methods are just wrappers for the model's methods
}

func (m rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if Dump != nil {
		spew.Fdump(Dump, msg)
	}

	return m.model.Update(msg)
}

func (m rootScreenModel) View() string {
	return m.model.View()
}

// this is the switcher which will switch between screens
func (m rootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m, tea.Sequence(
		m.model.Init(),
		tea.Batch(tea.ClearScreen, tea.WindowSize()),
	) // must return .Init() to initialize the screen (and here the magic happens)
}
