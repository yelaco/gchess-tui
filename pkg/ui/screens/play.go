package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/pkg/ui/stages/play/gameplay"
	"github.com/yelaco/gchess-tui/pkg/ui/theme"
)

// TODO: implement play screen
type PlayScreenModel struct {
	stage         tea.Model
	width, height int
}

func NewPlayScreenModel() PlayScreenModel {
	return PlayScreenModel{
		stage: gameplay.NewGamePlayStageModel(),
	}
}

func (m PlayScreenModel) Init() tea.Cmd {
	return m.stage.Init()
}

func (m PlayScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			m.stage, cmd = m.stage.Update(msg)
			return m, cmd
		}
	default:
		m.stage, cmd = m.stage.Update(msg)
	}

	return m, cmd
}

func (m PlayScreenModel) View() string {
	header := theme.HeaderStyle.Width(m.width).Render("Play")
	footer := theme.FooterStyle.Width(m.width).Render("")
	content := theme.ContentStyle.
		Width(m.width).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer)).
		Render(m.stage.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
