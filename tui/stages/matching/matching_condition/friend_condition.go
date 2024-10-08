package matching

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/dtos"
)

type MatchingFriendConditionModel struct {
	user dtos.User
}

func NewMatchingFriendConditionModel(user dtos.User) MatchingFriendConditionModel {
	return MatchingFriendConditionModel{
		user: user,
	}
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
