package login

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/domains/user"
)

// Submit info
type LoginInfoMsg struct {
	username string
	password string
}

func submitLoginInfo(info LoginInfoMsg) tea.Cmd {
	return func() tea.Msg {
		return info
	}
}

// Login
type LoginResultMsg struct {
	user user.User
}

func doLogin(info LoginInfoMsg) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second)
		return LoginResultMsg{
			user: user.User{
				Username:  info.username,
				Email:     "test@gmail.com",
				Rating:    100,
				CreatedAt: time.Now(),
			},
		}
	}
}

// Cancel
type LoginCancelMsg struct{}

func cancelLogin() tea.Cmd {
	return func() tea.Msg {
		return LoginCancelMsg{}
	}
}

// Load info
type jobFinishedMsg struct {
	name     string
	duration time.Duration
	result   interface{}
}

type jobFailedMsg struct {
	name     string
	duration time.Duration
}

func succeedStubJob() tea.Cmd {
	return func() tea.Msg {
		duration := time.Duration(rand.Intn(10)+1) * time.Second
		time.Sleep(duration)
		return jobFinishedMsg{
			name:     "Stub job",
			duration: duration,
			result:   "Done",
		}
	}
}

func failedStubJob() tea.Cmd {
	return func() tea.Msg {
		duration := time.Duration(rand.Intn(10)+1) * time.Second
		time.Sleep(duration)
		return jobFailedMsg{
			name:     "Stub job",
			duration: duration,
		}
	}
}

// Complete
type LoginCompleteMsg struct {
	User user.User
}

func completeLogin(user user.User) tea.Cmd {
	return func() tea.Msg {
		return LoginCompleteMsg{
			User: user,
		}
	}
}
