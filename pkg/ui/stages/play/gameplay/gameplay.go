package gameplay

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/pkg/app"
	"github.com/yelaco/gchess-tui/pkg/ui/stages/play"
	"github.com/yelaco/gchess-tui/pkg/utils"
)

const (
	startingBoard = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	startingX     = 6
	startingY     = 3
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

type position struct{ x, y int }

type GamePlayStageModel struct {
	Board            [][]string
	startPos         *position
	endPos           *position
	waitForSelection bool
}

func NewGamePlayStageModel() GamePlayStageModel {
	m := GamePlayStageModel{
		startPos: &position{x: startingX, y: startingY},
	}

	m.setBoard(startingBoard)
	return m
}

func (m *GamePlayStageModel) setBoard(boardFen string) {
	if m.Board == nil {
		m.Board = make([][]string, 8)
		for i := range 8 {
			m.Board[i] = make([]string, 8)
		}
	}
	rows := strings.Split(boardFen, "/")
	for x, row := range rows {
		var y int
		for _, ch := range row {
			switch ch {
			case 'r':
				m.Board[x][y] = "♖"
			case 'n':
				m.Board[x][y] = "♘"
			case 'b':
				m.Board[x][y] = "♗"
			case 'q':
				m.Board[x][y] = "♕"
			case 'k':
				m.Board[x][y] = "♔"
			case 'p':
				m.Board[x][y] = "♙"
			case 'R':
				m.Board[x][y] = "♜"
			case 'N':
				m.Board[x][y] = "♞"
			case 'B':
				m.Board[x][y] = "♝"
			case 'Q':
				m.Board[x][y] = "♛"
			case 'K':
				m.Board[x][y] = "♚"
			case 'P':
				m.Board[x][y] = "♟"
			default:
				numSpaces, err := strconv.Atoi(string(ch))
				if err != nil {
					panic(err)
				}
				for range numSpaces {
					m.Board[x][y] = "."
					y++
				}
				continue
			}
			y++
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
			if (m.endPos != nil && i == m.endPos.x && j == m.endPos.y) ||
				(m.startPos != nil && i == m.startPos.x && j == m.startPos.y) {
				renderedRows = append(renderedRows, " ", focusedCellStyle.Render(m.Board[i][j]))
				continue
			}
			renderedRows = append(renderedRows, " ", blurredCellStyle.Render(m.Board[i][j]))

		}
		renderedRows = append(renderedRows, " ")
		renderedBoard = append(renderedBoard, lipgloss.JoinHorizontal(lipgloss.Center, renderedRows...))
	}

	return boardStyle.Render(lipgloss.JoinVertical(lipgloss.Center, renderedBoard...))
}

func (m GamePlayStageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()
		switch s {
		case "esc":
			m.endPos = nil
			m.waitForSelection = false
		case "enter":
			// Allow new move if it is user's turn
			if app.GetMatch().GameState.IsWhiteTurn == app.GetMatch().PlayerState.IsWhiteSide {
				if m.waitForSelection {
					return m, sendMove(utils.BoardToFen(m.Board), m.CurrentMove())
				}
				if m.endPos == nil {
					m.waitForSelection = true
					m.endPos = &position{m.startPos.x, m.startPos.y}
				}
			}
		case "up", "down", "left", "right", "j", "k", "h", "l":
			m.moveCursor(s)
		}
	case play.GameUpdatedMsg:
		m.NextState()
	}

	return m, cmd
}

func (m *GamePlayStageModel) moveCursor(direction string) {
	pos := m.startPos
	if m.waitForSelection {
		pos = m.endPos
	}
	if pos == nil {
		panic("nil position")
	}
	switch direction {
	case "up", "k":
		if pos.x > 0 {
			pos.x--
		}
	case "down", "j":
		if pos.x < 7 {
			pos.x++
		}
	case "left", "h":
		if pos.y > 0 {
			pos.y--
		}
	case "right", "l":
		if pos.y < 7 {
			pos.y++
		}
	}
}

func (m GamePlayStageModel) CurrentMove() string {
	if m.endPos == nil {
		return ""
	}
	return fmt.Sprintf("%c%d%c%d", 'a'+m.startPos.y, 8-m.startPos.x, 'a'+m.endPos.y, 8-m.endPos.x)
}

func (m *GamePlayStageModel) NextState() {
	m.setBoard(app.GetMatch().GameState.BoardFen)
	{
		m.startPos.x = startingX
		m.startPos.y = startingY
		m.endPos = nil
	}
	m.waitForSelection = false
}
