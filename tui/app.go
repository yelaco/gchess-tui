package tui

import (
	"fmt"
	"log"
	"os"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
	"github.com/yelaco/gchess-tui/infrastructures/login"
	"github.com/yelaco/gchess-tui/usecases"
	"github.com/yelaco/gchess-tui/util"
)

var lock = &sync.Mutex{}

type app struct {
	LoginUsecase usecases.LoginUsecaseInterface
	msgDump      *os.File
	appDump      *os.File
	config       *util.Config
}

var singleApp *app

func GetApp() *app {
	if singleApp == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleApp == nil {
			singleApp = newApp()
		}
	}

	return singleApp
}

func DumpMsgLog(model, msg tea.Msg) {
	if d := GetApp().msgDump; d != nil {
		fmt.Fprintf(d, "%s: %#v\n", model, msg)
		// spew.Fdump(d, "%s: %#v\n", model, msg)
	}
}

func DumpAppLog(value interface{}) {
	if d := GetApp().appDump; d != nil {
		spew.Fdump(d, value)
	}
}

func newApp() *app {
	config, err := util.LoadConfig("./.infra")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	var msgDump *os.File
	var appDump *os.File
	if config.Debug {
		var err error
		msgDump, err = os.OpenFile("messages.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal("cannot open message dump file: ", err)
		}
		appDump, err = os.OpenFile("app.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal("cannot open app dump file: ", err)
		}
	}

	// servicelUrl, _ := url.Parse(config.ServiceUrl)

	// Dependency injection
	loginOperation := login.NewOperation(config.ServiceUrl)

	loginUsecase := usecases.NewLoginUsecase(loginOperation)

	return &app{
		LoginUsecase: loginUsecase,
		msgDump:      msgDump,
		appDump:      appDump,
		config:       &config,
	}
}
