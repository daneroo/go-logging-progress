package coco

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

type PBHook struct {
}

func (hook *PBHook) Fire(entry *logrus.Entry) error {
	clear := "\r" + strings.Repeat(" ", 120) + "\r"
	fmt.Print(clear)
	fmt.Print("I am a hooook.\n")
	return nil
}

func (hook *PBHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
