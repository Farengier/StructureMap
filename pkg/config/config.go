package config

type Conf interface {
	Locale() string
	LocaleDomains() []string
	LocaleDirectory() string
}

type conf struct {
}

func (c conf) Locale() string {
	return "ru"
}

func (c conf) LocaleDirectory() string {
	return "./localization/"
}

func (c conf) LocaleDomains() []string {
	return []string{"messages"}
}

func FromFile(fileName string) Conf {
	c := conf{}
	return c
}
