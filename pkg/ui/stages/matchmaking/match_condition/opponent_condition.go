package match_condition

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

type MatchingOpponentConditionModel struct {
	condition domains.MatchCondition
}

func NewMatchingOpponentConditionModel() MatchingOpponentConditionModel {
	return MatchingOpponentConditionModel{}
}

func (m MatchingOpponentConditionModel) Init() tea.Cmd {
	return nil
}

func (m MatchingOpponentConditionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "P":
			return m, confirmMatchingCondition(m.condition)
		}
	}
	return m, nil
}

func (m MatchingOpponentConditionModel) View() string {
	return "Matching opponent"
}
