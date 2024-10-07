package home

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/assets"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Align(lipgloss.Left, lipgloss.Top)
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Align(lipgloss.Left, lipgloss.Top)
)

const (
	PlayOption         = "Start a match"
	SolvePuzzleOption  = "Solve puzzle"
	ViewMatchOption    = "View match"
	MatchHistoryOption = "Match history"
	SocialOption       = "Social"
	LogOutOption       = "Logout"
)

type MenuStageModel struct {
	options  []string
	choice   int
	selected bool
	quitting bool
}

func NewMenuStageModel() MenuStageModel {
	return MenuStageModel{
		options: []string{
			PlayOption,
			SolvePuzzleOption,
			ViewMatchOption,
			MatchHistoryOption,
			SocialOption,
			LogOutOption,
		},
	}
}

func (m MenuStageModel) Init() tea.Cmd {
	return nil
}

func (m MenuStageModel) View() string {
	options := make([]string, 0, len(m.options)+2)

	for i, option := range m.options {
		if i == m.choice {
			options = append(options, focusedStyle.Render(option))
		} else {
			options = append(options, blurredStyle.Render(option))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Center, "", assets.GetLogo(), lipgloss.JoinVertical(lipgloss.Center, options...), "")
}

func (m MenuStageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			return m, m.Select()
		case "tab", "shift+tab", "up", "down":
			m.moveCursor(msg)
		}
	}

	return m, cmd
}

func (m *MenuStageModel) moveCursor(msg tea.KeyMsg) {
	switch msg.String() {
	case "up", "shift+tab":
		if m.choice > 0 {
			m.choice--
		} else {
			m.choice = len(m.options) - 1
		}
	case "down", "tab":
		if m.choice < len(m.options)-1 {
			m.choice++
		} else {
			m.choice = 0
		}
	}
}

func (m MenuStageModel) Select() tea.Cmd {
	return func() tea.Msg {
		switch m.options[m.choice] {
		case PlayOption:
			return PlayMsg{}
		case SolvePuzzleOption:
			return SolvePuzzleMsg{}
		case ViewMatchOption:
			return ViewMatchMsg{}
		case MatchHistoryOption:
			return MatchHistoryMsg{}
		case SocialOption:
			return SocialMsg{}
		case LogOutOption:
			return LogOutMsg{}
		default:
			return nil
		}
	}
}
