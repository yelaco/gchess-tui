package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/dtos"
)

// TODO: implement play screen
type PlayScreenModel struct {
	matchInfo dtos.MatchInfo
	model     tea.Model
}

func NewPlayScreenModel(dtos.MatchInfo) PlayScreenModel {
	return PlayScreenModel{
		model: DefaultModel{},
	}
}

func (m PlayScreenModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m PlayScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m PlayScreenModel) View() string {
	return m.model.View()
}
