package match_condition

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

type MatchingFriendConditionModel struct {
	user domains.User
}

func NewMatchingFriendConditionModel() MatchingFriendConditionModel {
	return MatchingFriendConditionModel{}
}

func (m MatchingFriendConditionModel) Init() tea.Cmd {
	return nil
}

func (m MatchingFriendConditionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MatchingFriendConditionModel) View() string {
	return "Matching Friend"
}
