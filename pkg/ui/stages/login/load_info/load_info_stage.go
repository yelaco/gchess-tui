package load_info

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yelaco/gchess-tui/assets"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
	loginstages "github.com/yelaco/gchess-tui/pkg/ui/stages/login"
	"github.com/yelaco/gchess-tui/pkg/utils"
)

var (
	loadInfoHelpStyle = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("250"))
	loadInfoMainStyle = lipgloss.NewStyle().Align(lipgloss.Center)
)

type result struct {
	name     string
	duration time.Duration
	info     interface{}
}

type LoadInfoStageModel struct {
	user    domains.User
	spinner spinner.Model
	jobs    []tea.Cmd
	results []result
	done    bool
	failed  bool
}

const showLastResults = 5

func NewLoadInfoStageModel(user domains.User) LoadInfoStageModel {
	sp := spinner.New()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#4e7837"))

	jobs := []tea.Cmd{
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
		if m.failed {
			return m, tea.Quit
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.done {
				return m, tea.Batch(tea.ClearScreen, loginstages.CompleteLogin(m.user))
			}
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case loginstages.LoadInfoFinishedMsg:
		if m.failed {
			return m, nil
		}
		res := result{name: assets.GetCheckMark() + " " + msg.Name, duration: msg.Duration, info: msg.Result}
		if len(m.results) < showLastResults {
			m.results = append(m.results, res)
		} else {
			m.results = append(m.results[1:], res)
		}
		if len(m.results) == len(m.jobs) {
			m.done = true
		}
	case loginstages.LoadInfoFailedMsg:
		res := result{name: assets.GetXMark() + " " + msg.Name, duration: msg.Duration}
		if len(m.results) < showLastResults {
			m.results = append(m.results, res)
		} else {
			m.results = append(m.results[1:], res)
		}
		m.failed = true
	}

	return m, nil
}

func (m LoadInfoStageModel) View() string {
	var b strings.Builder
	var results string

	if m.done {
		b.WriteString("\n" + "🎉 " + "Loading information...\n\n")
		b.WriteString("%s")
		b.WriteString("\nDone! Press enter to continue\n")
	} else if m.failed {
		b.WriteString("\n" + "🚨 " + "Loading information\n\n")
		b.WriteString("%s")
		b.WriteString("\nFailed! Press any key to quit\n")
	} else {
		b.WriteString("\n" + m.spinner.View() + " Loading information...\n\n")
		b.WriteString("%s")
		b.WriteString(loadInfoHelpStyle.Render("\nPress ctrl+c or q to quit\n"))
	}

	for _, res := range m.results {
		if res.info != nil {
			results += fmt.Sprintf("%s: finished in %s\n", res.name, res.duration)
		} else {
			results += fmt.Sprintf("%s: failed in %s\n", res.name, res.duration)
		}
	}
	results += utils.GenerateBlankLine(showLastResults - len(m.results))

	return loadInfoMainStyle.Render(fmt.Sprintf(b.String(), results))
}
