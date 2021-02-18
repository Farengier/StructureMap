package log

import (
	"fmt"
)

type Logger interface {
	Debug(vars ...interface{})
}

type Log struct {
}

func Init() Log {
	return Log{}
}

func (l Log) Debug(vars ...interface{}) {
	fmt.Println(vars...)
}
