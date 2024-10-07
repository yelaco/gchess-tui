package matching

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/user"
)

type MatchingAiConditionModel struct {
	user user.User
}

func NewMatchingAiConditionModel(user user.User) MatchingAiConditionModel {
	return MatchingAiConditionModel{
		user: user,
	}
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
