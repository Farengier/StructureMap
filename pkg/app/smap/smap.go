package smap

import (
	"github.com/Farengier/StructureMap/pkg/config"
	"github.com/Farengier/StructureMap/pkg/i18n"
	"github.com/Farengier/StructureMap/pkg/log"
	"github.com/Farengier/StructureMap/pkg/windows"
)

type Application struct {
	cfg config.Conf
	loc i18n.Localizator
	log log.Logger
}

func Init(cfg string) Application {
	conf := config.FromFile(cfg)
	app := Application{
		cfg: conf,
		loc: i18n.Init(conf),
		log: log.Init(),
	}
	return app
}

func (a Application) Run() {
	a.log.Debug(a.loc.T("Log|Application|Run|Starting application"))
	windows.Main(a)
}

func (a Application) Loc() i18n.Localizator {
	return a.loc
}

func (a Application) Log() log.Logger {
	return a.log
}
