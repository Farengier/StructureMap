package i18n

import (
	"github.com/Farengier/StructureMap/pkg/config"
	"github.com/leonelquinteros/gotext"
)

type Loc interface {
	T(message string, vars ...interface{}) string
}

type loc struct {
	gl *gotext.Locale
}

func Init(cfg config.Conf) Loc {
	l := gotext.NewLocale(cfg.LocaleDirectory(), cfg.Locale())
	for _, d := range cfg.LocaleDomains() {
		l.AddDomain(d)
	}
	return loc{gl: l}
}

func (l loc) T(message string, vars ...interface{}) string {
	return l.gl.Get(message, vars...)
}
