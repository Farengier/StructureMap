package smap

import "fmt"

type Application interface {
	Run()
}

type application struct {
}

func Init() Application {
	app := application{}
	return app
}

func (a application) Run() {
	fmt.Print("Hello, i`m a smap program\n")
}
