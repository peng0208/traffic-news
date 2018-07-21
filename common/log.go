package common

import (
	"os"
	"github.com/sirupsen/logrus"
	"sync"
)

var logger *logrus.Logger
var once sync.Once

func Logger() *logrus.Logger {
	f := Cfg.Section("log").Key("log").String()
	once.Do(func() {
		logger = newLogger(f)
	})
	return logger
}

func newLogger(f string) *logrus.Logger {
	l := logrus.New()
	logf, err := os.OpenFile(f, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	l.SetOutput(logf)
	l.SetLevel(logrus.DebugLevel)

	return l
}
