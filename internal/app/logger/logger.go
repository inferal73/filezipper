package logger

import (
	"fmt"
	"io"
	"os"
)

func init()  {
	instance = &logger{
		w: os.Stdout,
	}
}

type logger struct {
	w io.Writer
}

var instance *logger

func GetLogger() *logger {
	return instance
}

func SetWriter(writer io.Writer) {
	instance.w = writer
}

func Log(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(instance.w, format, a...)
}