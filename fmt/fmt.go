package fmt

import (
	"fmt"
)

type fmtLogger struct{}

func (t *fmtLogger) Log(v ...interface{}) {
	fmt.Print(v...)
}

func (t *fmtLogger) Logf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func New() *fmtLogger {
	return &fmtLogger{}
}
