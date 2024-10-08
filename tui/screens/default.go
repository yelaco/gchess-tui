package screens

import tea "github.com/charmbracelet/bubbletea"

type DefaultModel struct{}

func (m DefaultModel) Init() tea.Cmd {
	return nil
}

func (m DefaultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m DefaultModel) View() string {
	return "No model found"
}
