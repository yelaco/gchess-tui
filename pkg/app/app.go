package app

import (
	"fmt"
	"log"
	"os"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
	"github.com/yelaco/gchess-tui/configs"
	"github.com/yelaco/gchess-tui/pkg/gchessclient"
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

var mu = &sync.RWMutex{}

// app struct    stores application state
type app struct {
	msgDump *os.File
	appDump *os.File

	Client gchessclient.Client
	Config configs.Config
	User   domains.User
	Match  domains.Match
}

var singleApp *app

// getApp function    returns the same singleton app and initialize it if not done already
func getApp() *app {
	if singleApp == nil {
		if singleApp == nil {
			singleApp = newApp()
		}
	}

	return singleApp
}

// newApp function    initialize app
func newApp() *app {
	config, err := configs.LoadConfig("./.infra")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	var msgDump *os.File
	var appDump *os.File
	if config.Debug {
		var err error
		msgDump, err = os.OpenFile("messages.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal("cannot open dump file: ", err)
		}
		appDump, err = os.OpenFile("app.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal("cannot open dump file: ", err)
		}
	}

	// servicelUrl, _ := url.Parse(config.ServiceUrl)

	return &app{
		msgDump: msgDump,
		appDump: appDump,
		Config:  config,
	}
}

// DumpMsgLog function    outputs the message log of bubbletea package's method to file messages.log
func DumpMsgLog(model, msg tea.Msg) {
	mu.Lock()
	defer mu.Unlock()
	if d := getApp().msgDump; d != nil {
		fmt.Fprintf(d, "%s: %#v\n", model, msg)
	}
}

// DumpAppLog function    outputs the application log file app.log
func DumpAppLog(value interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if d := getApp().appDump; d != nil {
		spew.Fdump(d, value)
	}
}

// GetConfig function    returns a copy of current config
func GetConfig() configs.Config {
	mu.RLock()
	defer mu.RUnlock()
	return getApp().Config
}

func SetClient(client gchessclient.Client) {
	mu.Lock()
	defer mu.Unlock()
	if getApp().Client != nil {
		err := getApp().Client.Close()
		if err != nil {
			DumpAppLog(err)
		}
	}
	getApp().Client = client
}
