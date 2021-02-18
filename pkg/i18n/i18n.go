package i18n

import (
	"github.com/leonelquinteros/gotext"
)

type cfg interface {
	Locale() string
	LocaleDirectory() string
	LocaleDomains() []string
}

type Localizator interface {
	T(message string, vars ...interface{}) string
}

type Loc struct {
	gl *gotext.Locale
}

func Init(cfg cfg) Loc {
	l := gotext.NewLocale(cfg.LocaleDirectory(), cfg.Locale())
	for _, d := range cfg.LocaleDomains() {
		l.AddDomain(d)
	}
	return Loc{gl: l}
}

func (l Loc) T(message string, vars ...interface{}) string {
	return l.gl.Get(message, vars...)
}
