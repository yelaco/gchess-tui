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
	stage         tea.Model
	width, height int
}

func NewPlayScreenModel(matchInfo dtos.MatchInfo) PlayScreenModel {
	return PlayScreenModel{
		matchInfo: matchInfo,
		stage:     gameplay.NewGamePlayStageModel(matchInfo),
	}
}

func (m PlayScreenModel) Init() tea.Cmd {
	return m.stage.Init()
}

func (m PlayScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tui.DumpMsgLog(reflect.TypeOf(m).Name(), msg)
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
