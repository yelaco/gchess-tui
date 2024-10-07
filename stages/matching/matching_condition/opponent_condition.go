package matching

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/user"
)

type MatchingOpponentConditionModel struct {
	user user.User
}

func NewMatchingOpponentConditionModel(user user.User) MatchingOpponentConditionModel {
	return MatchingOpponentConditionModel{
		user: user,
	}
}

func (m MatchingOpponentConditionModel) Init() tea.Cmd {
	return nil
}

func (m MatchingOpponentConditionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MatchingOpponentConditionModel) View() string {
	return "Matching opponent hahahahahahahahahahahahah\n\n\n\n\n\n"
}
