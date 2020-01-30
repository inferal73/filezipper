package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger interface {
	Log(format string, a ...interface{}) error
	SetWriter(writer io.Writer)
}

type logger struct {
	w io.Writer
	sync.RWMutex
}

var instance *logger
var once sync.Once

func GetLogger() Logger {
	once.Do(func() {
		instance = new(logger)
		instance.SetWriter(os.Stdout)
	})
	return instance
}

func (l *logger) Log(format string, a ...interface{}) error {
	_, err := fmt.Fprintf(l.w, format, a...)
	return err
}

func (l *logger) SetWriter(writer io.Writer) {
	l.Lock()
	defer l.Unlock()
	l.w = writer
}