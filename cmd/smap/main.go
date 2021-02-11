package main

import "github.com/Farengier/StructureMap/pkg/app/smap"

func main() {
	app := smap.Init()
	app.Run()
}
