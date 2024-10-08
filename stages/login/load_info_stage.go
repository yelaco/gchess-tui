package login

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/user"
	"github.com/yelaco/gchess-tui/util"
)

var (
	loadInfoHelpStyle = blurredStyle
	loadInfoMainStyle = lipgloss.NewStyle().MarginLeft(1)
)

type result struct {
	name     string
	duration time.Duration
	info     interface{}
}

type LoadInfoStageModel struct {
	user         user.User
	spinner      spinner.Model
	jobs         []tea.Cmd
	results      []result
	SuccessCount int
	FailCount    int
	done         bool
}

const showLastResults = 5

func NewLoadInfoStageModel(user user.User) LoadInfoStageModel {
	sp := spinner.New()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("206"))

	jobs := []tea.Cmd{
		succeedStubJob(),
		succeedStubJob(),
		succeedStubJob(),
		succeedStubJob(),
		succeedStubJob(),
		succeedStubJob(),
		succeedStubJob(),
		succeedStubJob(),
		succeedStubJob(),
	}
	return LoadInfoStageModel{
		user:    user,
		spinner: sp,
		results: make([]result, 0, showLastResults),
		jobs:    jobs,
	}
}

func (m LoadInfoStageModel) Init() tea.Cmd {
	jobBatch := tea.Batch(m.jobs...)
	return tea.Batch(m.spinner.Tick, jobBatch)
}

func (m LoadInfoStageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.FailCount > 0 && m.done {
			return m, tea.Quit
		}
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.done {
				return m, tea.Batch(tea.ClearScreen, completeLogin(m.user))
			}
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case jobFinishedMsg:
		res := result{name: "✅ " + msg.name, duration: msg.duration, info: msg.result}
		if len(m.results) < showLastResults {
			m.results = append(m.results, res)
		} else {
			m.results = append(m.results[1:], res)
		}
		m.SuccessCount += 1
		if m.SuccessCount+m.FailCount == len(m.jobs) {
			m.done = true
		}
	case jobFailedMsg:
		res := result{name: "❌ " + msg.name, duration: msg.duration}
		if len(m.results) < showLastResults {
			m.results = append(m.results, res)
		} else {
			m.results = append(m.results[1:], res)
		}
		m.FailCount += 1
		if m.SuccessCount+m.FailCount == len(m.jobs) {
			m.done = true
		}

	}

	return m, nil
}

func (m LoadInfoStageModel) View() string {
	var b strings.Builder
	var results string

	if m.done {
		if m.FailCount > 0 {
			b.WriteString("\n" + "❌ " + "Loading information...\n\n")
			b.WriteString("%s")
			b.WriteString("\nFailed! Press any key to quit\n")
		} else {
			b.WriteString("\n" + "✅ " + "Loading information...\n\n")
			b.WriteString("%s")
			b.WriteString("\nDone! Press enter to continue\n")
		}
	} else {
		b.WriteString("\n" + m.spinner.View() + " Loading information...\n\n")
		b.WriteString("%s")
		b.WriteString(loadInfoHelpStyle.Render("\nPress ctrl+c to quit\n"))
	}

	for _, res := range m.results {
		results += fmt.Sprintf("%s: finished in %s\n", res.name, res.duration)
	}
	results += util.GenerateBlankLine(showLastResults - len(m.results))
	results += fmt.Sprintf("\nSucceed jobs: %d/%d", m.SuccessCount, len(m.jobs))

	return loadInfoMainStyle.Render(fmt.Sprintf(b.String(), results))
}
