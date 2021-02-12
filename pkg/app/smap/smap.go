package smap

import (
	"fmt"
	"github.com/Farengier/StructureMap/pkg/config"
	"github.com/Farengier/StructureMap/pkg/i18n"
)

type Application interface {
	Run()
}

type application struct {
	cfg config.Conf
	loc i18n.Loc
}

func Init(cfg config.Conf) Application {
	app := application{
		cfg: cfg,
		loc: i18n.Init(cfg),
	}
	return app
}

func (a application) Run() {
	fmt.Println(a.loc.T("Test|Hello message"))
}
