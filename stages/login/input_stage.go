package login

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#4e7837"))
	blurredStyle = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("250"))
	cursorStyle  = focusedStyle
	noStyle      = lipgloss.NewStyle()
	borderStyle  = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	// helpStyle           = blurredStyle
	// cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedLoginButton = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#4e7837")).
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#4e7837")).
				Render("Login")

	blurredLoginButton = lipgloss.NewStyle().
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				Render("Login")
)

type InputStageModel struct {
	inputs    []textinput.Model
	submitted bool
	cursor    int
}

func NewInputStageModel() InputStageModel {
	m := InputStageModel{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Username"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.Width = 20
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.PromptStyle = blurredStyle
			t.TextStyle = blurredStyle
			t.EchoCharacter = '•'
			t.Width = 20
		}

		m.inputs[i] = t
	}

	return m
}

func (m InputStageModel) Init() tea.Cmd {
	return nil
}

func (m InputStageModel) View() string {
	var b strings.Builder

	b.WriteString("Enter information\n")
	for i := range m.inputs {
		b.WriteString(borderStyle.Render(m.inputs[i].View()) + "\n")
	}
	button := &blurredLoginButton
	if m.cursor == len(m.inputs) {
		button = &focusedLoginButton
	}
	b.WriteString(*button)

	return b.String()
}

func (m InputStageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			authStageModel := NewAuthStageModel(LoginInfoMsg{
				m.inputs[0].Value(),
				m.inputs[1].Value(),
			})
			return authStageModel, tea.Batch(tea.ClearScreen, authStageModel.Init())
		case "tab", "shift+tab", "up", "down":
			// this stage ended
			if m.submitted {
				return m, nil
			}

			// move cursor position
			m.moveCursor(msg)

			// update input state
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.cursor {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = blurredStyle
				m.inputs[i].TextStyle = blurredStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m InputStageModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *InputStageModel) moveCursor(msg tea.KeyMsg) {
	switch msg.String() {
	case "up", "shift+tab":
		if m.cursor > 0 {
			m.cursor--
		} else {
			m.cursor = len(m.inputs) - 1
		}
	case "down", "tab":
		if m.cursor < len(m.inputs) {
			m.cursor++
		} else {
			m.cursor = 0
		}
	}
}
