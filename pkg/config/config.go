package config

type Conf struct {
}

func (c Conf) Locale() string {
	return "ru"
}

func (c Conf) LocaleDirectory() string {
	return "./localization/"
}

func (c Conf) LocaleDomains() []string {
	return []string{"messages"}
}

func FromFile(fileName string) Conf {
	c := Conf{}
	return c
}
