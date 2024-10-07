package login

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/domains/user"
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
		failedStubJob(),
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
	var head string
	var results string
	var tail string

	if m.done {
		if m.FailCount > 0 {
			head = "\n" + "❌ " + "Loading information...\n\n"
			tail = "\nFailed! Press any key to quit\n"
		} else {
			head = "\n" + "✅ " + "Loading information...\n\n"
			tail = "\nDone! Press enter to continue\n"
		}
	} else {
		head = "\n" + m.spinner.View() + " Loading information...\n\n"
		tail = loadInfoHelpStyle.Render("\nPress ctrl+c to quit\n")
	}

	for _, res := range m.results {
		results += fmt.Sprintf("%s: finished in %s\n", res.name, res.duration)
	}
	results += generateBlankLine(showLastResults - len(m.results))
	results += fmt.Sprintf("\nSucceed jobs: %d/%d", m.SuccessCount, len(m.jobs))

	return loadInfoMainStyle.Render(head + results + tail)
}

func generateBlankLine(count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat("\n", count)
}
