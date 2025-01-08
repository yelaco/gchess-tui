package match_condition

import (
	tea "github.com/charmbracelet/bubbletea"
	domains "github.com/yelaco/gchess-tui/pkg/ui/domains"
)

type MatchingAiConditionModel struct {
	user domains.User
}

func NewMatchingAiConditionModel() MatchingAiConditionModel {
	return MatchingAiConditionModel{}
}

func (m MatchingAiConditionModel) Init() tea.Cmd {
	return nil
}

func (m MatchingAiConditionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MatchingAiConditionModel) View() string {
	return "Matching AI"
}
