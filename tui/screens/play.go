package screens

import (
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/dtos"
	"github.com/yelaco/gchess-tui/tui"
	"github.com/yelaco/gchess-tui/tui/stages/play/gameplay"
	"github.com/yelaco/gchess-tui/tui/theme"
)

// TODO: implement play screen
type PlayScreenModel struct {
	matchInfo     dtos.MatchInfo
	model         tea.Model
	width, height int
}

func NewPlayScreenModel(matchInfo dtos.MatchInfo) PlayScreenModel {
	return PlayScreenModel{
		matchInfo: matchInfo,
		model:     gameplay.NewGamePlayStageModel(matchInfo),
	}
}

func (m PlayScreenModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m PlayScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tui.DumpMsgLog(reflect.TypeOf(m).Name(), msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.model, cmd = m.model.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m PlayScreenModel) View() string {
	header := theme.HeaderStyle.Width(m.width).Render("Play")
	footer := theme.FooterStyle.Width(m.width).Render("")
	content := theme.ContentStyle.
		Width(m.width).
		Height(m.height - lipgloss.Height(header) - lipgloss.Height(footer)).
		Render(m.model.View())

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
