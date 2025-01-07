package load_info

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	loginstages "github.com/yelaco/gchess-tui/pkg/ui/stages/login"
)

func succeedStubJob() tea.Cmd {
	return func() tea.Msg {
		duration := time.Duration(rand.Intn(10)+1) * time.Second
		time.Sleep(duration)
		return loginstages.LoadInfoFinishedMsg{
			Name:     "Stub job",
			Duration: duration,
			Result:   "Done",
		}
	}
}

func failedStubJob() tea.Cmd {
	return func() tea.Msg {
		duration := time.Duration(rand.Intn(10)+1) * time.Second
		time.Sleep(duration)
		return loginstages.LoadInfoFailedMsg{
			Name:     "Stub job",
			Duration: duration,
		}
	}
}
