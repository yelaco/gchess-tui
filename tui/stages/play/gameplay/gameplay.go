package gameplay

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/dtos"
	"github.com/yelaco/gchess-tui/tui/stages/play"
	"github.com/yelaco/gchess-tui/util"
)

const (
	startingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
)

var (
	boardBorder      = lipgloss.ThickBorder()
	blurredCellStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.HiddenBorder()).
				Width(3).
				Height(1).
				Align(lipgloss.Center, lipgloss.Center)
	focusedCellStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#69923E")).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("#69923E")).
				Width(3).
				Height(1).
				Align(lipgloss.Center, lipgloss.Center)
	boardStyle = lipgloss.NewStyle().BorderStyle(boardBorder).Align(lipgloss.Center, lipgloss.Center)
)

type index struct{ x, y int }

type GamePlayStageModel struct {
	matchInfo         dtos.MatchInfo
	Board             [][]string
	focusIndex        index
	selectedIndex     *index
	availableMoves    []index
	moveChoice        int
	waitingGameUpdate bool
}

func NewGamePlayStageModel(matchInfo dtos.MatchInfo) GamePlayStageModel {
	m := GamePlayStageModel{
		matchInfo:      matchInfo,
		focusIndex:     index{x: 6, y: 3},
		availableMoves: nil,
	}

	m.setBoard(matchInfo.Fen)
	return m
}

func (m *GamePlayStageModel) setBoard(fen string) {
	m.Board = make([][]string, 8)
	for i := 0; i < 8; i++ {
		m.Board[i] = make([]string, 8)
		for j := 0; j < 8; j++ {
			m.Board[i][j] = "•"
		}
	}
}

func (m GamePlayStageModel) Init() tea.Cmd {
	return nil
}

func (m GamePlayStageModel) View() string {
	renderedBoard := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		renderedRows := make([]string, 0, 8)
		for j := 0; j < 8; j++ {
			if i == m.focusIndex.x && j == m.focusIndex.y {
				renderedRows = append(renderedRows, " ", focusedCellStyle.Render(m.Board[i][j]))
			} else if m.selectedIndex != nil && i == m.selectedIndex.x && j == m.selectedIndex.y {
				renderedRows = append(renderedRows, " ", focusedCellStyle.Render(m.Board[i][j]))
			} else {
				renderedRows = append(renderedRows, " ", blurredCellStyle.Render(m.Board[i][j]))
			}
		}
		renderedRows = append(renderedRows, " ")
		renderedBoard = append(renderedBoard, lipgloss.JoinHorizontal(lipgloss.Center, renderedRows...))
	}

	return boardStyle.Render(lipgloss.JoinVertical(lipgloss.Center, renderedBoard...))
}

func (m GamePlayStageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.waitingGameUpdate {
			return m, nil
		}
		s := msg.String()
		switch s {
		case "tab":
			if m.availableMoves == nil {
				m.updateAvailableMoves()
			} else {
				m.pickMoveChoice()
			}
		case "enter":
			if m.selectedIndex == nil {
				return m, nil
			}
			m.waitingGameUpdate = true
			return m, sendMove(util.BoardToFen(m.Board), m.CurrentMove())
		case "up", "down", "left", "right", "j", "k", "h", "l":
			m.moveCursor(s)
		}
	case play.GameUpdateMsg:
		m.NextState(msg)
	}

	return m, nil
}

func (m *GamePlayStageModel) moveCursor(direction string) {
	m.selectedIndex = nil
	m.availableMoves = nil
	switch direction {
	case "up", "k":
		if m.focusIndex.x > 0 {
			m.focusIndex.x--
		}
	case "down", "j":
		if m.focusIndex.x < 7 {
			m.focusIndex.x++
		}
	case "left", "h":
		if m.focusIndex.y > 0 {
			m.focusIndex.y--
		}
	case "right", "l":
		if m.focusIndex.y < 7 {
			m.focusIndex.y++
		}
	}
}

func (m *GamePlayStageModel) pickMoveChoice() {
	if m.availableMoves != nil {
		m.moveChoice = (m.moveChoice + 1) % len(m.availableMoves)
		m.selectedIndex = &m.availableMoves[m.moveChoice]
	}
}

func (m *GamePlayStageModel) updateAvailableMoves() {
	m.availableMoves = []index{{x: 4, y: 3}, {x: 4, y: 4}}
	m.moveChoice = 0
	m.selectedIndex = &m.availableMoves[m.moveChoice]
}

func (m GamePlayStageModel) CurrentMove() string {
	if m.selectedIndex == nil {
		return ""
	}
	return fmt.Sprintf("%c%d%c%d", 'a'+m.focusIndex.y, 8-m.focusIndex.x, 'a'+m.selectedIndex.y, 8-m.selectedIndex.x)
}

func (m *GamePlayStageModel) NextState(update play.GameUpdateMsg) {
	m.setBoard(update.Fen)
	m.focusIndex = index{}
	m.selectedIndex = nil
	m.availableMoves = nil
	m.waitingGameUpdate = false
}
