package tui

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yelaco/gchess-tui/infrastructures/login"
	"github.com/yelaco/gchess-tui/usecases"
	"github.com/yelaco/gchess-tui/util"
)

var lock = &sync.Mutex{}

type app struct {
	LoginUsecase usecases.LoginUsecaseInterface
	dump         *os.File
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
	if d := GetApp().dump; d != nil {
		fmt.Fprintf(d, "%s: %#v\n", model, msg)
		// spew.Fdump(d, "%s: %#v\n", model, msg)
	}
}

func newApp() *app {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	var dump *os.File
	if config.RunMode == "debug" {
		var err error
		dump, err = os.OpenFile("messages.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal("cannot open dump file: ", err)
		}
	}

	servicelUrl, _ := url.Parse(config.ServiceUrl)

	// Dependency injection
	loginOperation := login.NewOperation(servicelUrl)

	loginUsecase := usecases.NewLoginUsecase(loginOperation)

	return &app{
		LoginUsecase: loginUsecase,
		dump:         dump,
		config:       &config,
	}
}
