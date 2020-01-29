package logger

import (
	"fmt"
	"io"
	"os"
)

type Logger interface {
	Log(format string, a ...interface{}) error
	SetWriter(writer io.Writer)
}

type logger struct {
	w io.Writer
}

var instance *logger

func GetLogger() Logger {
	if instance == nil {
		instance = new(logger)
		instance.SetWriter(os.Stdout)
	}
	return instance
}

func (l *logger) Log(format string, a ...interface{}) error {
	_, err := fmt.Fprintf(l.w, format, a...)
	return err
}

func (l *logger) SetWriter(writer io.Writer) {
	l.w = writer
}