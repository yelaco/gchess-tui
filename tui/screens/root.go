package screens

import (
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/tui"
)

type rootScreenModel struct {
	screen tea.Model
}

func RootScreen() rootScreenModel {
	return rootScreenModel{
		screen: NewLoginScreenModel(),
	}
}

func (m rootScreenModel) Init() tea.Cmd {
	return m.screen.Init() // rest methods are just wrappers for the model's methods
}

func (m rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tui.DumpMsgLog(reflect.TypeOf(m).Name(), msg)

	return m.screen.Update(msg)
}

func (m rootScreenModel) View() string {
	return m.screen.View()
}

// this is the switcher which will switch between screens
func (m rootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.screen = model
	return m, tea.Sequence(
		m.screen.Init(),
		tea.Batch(tea.ClearScreen, tea.WindowSize()),
	) // must return .Init() to initialize the screen (and here the magic happens)
}
