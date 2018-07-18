package info

import (
	"fmt"
	"testing"

	"github.com/go-log/log"
)

type info struct {}

func (*info) Info(v ...interface{}) {
	fmt.Print(v...)
}

func (*info) Infof(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func testLog(l log.Logger) {
	l.Log("test\n")
}

func testLogf(l log.Logger) {
	l.Logf("%s", "test\n")
}

func TestNew(t *testing.T) {
	l := New(&info{})
	testLog(l)
	testLogf(l)
}
