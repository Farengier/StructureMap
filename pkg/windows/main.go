package windows

import (
	"github.com/Farengier/StructureMap/pkg/i18n"
	"github.com/Farengier/StructureMap/pkg/log"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

type app interface {
	Log() log.Logger
	Loc() i18n.Localizator
}

func Main(app app) {
	lg := app.Log()
	lc := app.Loc()
	setup := func() {
		mainwin := ui.NewWindow(lc.T("Window|Main|Title|Smap"), 640, 480, true)
		mainwin.OnClosing(func(*ui.Window) bool {
			lg.Debug(lc.T("Log|Window|Main|Closing"))
			ui.Quit()
			return true
		})
		ui.OnShouldQuit(func() bool {
			lg.Debug(lc.T("Log|Window|Main|Should quit"))
			mainwin.Destroy()
			return true
		})
		mainwin.Show()
	}
	err := ui.Main(setup)
	if err != nil {
		lg.Debug(lc.T("Log|Window|Main|Initialising error"), err)
	}
}
