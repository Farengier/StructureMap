package main

import (
	"github.com/Farengier/StructureMap/pkg/app/smap"
	"github.com/Farengier/StructureMap/pkg/config"
)

func main() {
	cfg := config.FromFile("config.cfg")
	app := smap.Init(cfg)
	app.Run()
}
