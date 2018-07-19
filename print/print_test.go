package print

import (
	"fmt"
	"testing"

	"github.com/go-log/log"
)

type printer struct {}

func (*printer) Print(v ...interface{}) {
	fmt.Print(v...)
}

func (*printer) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func testLog(l log.Logger) {
	l.Log("test\n")
}

func testLogf(l log.Logger) {
	l.Logf("%s", "test\n")
}

func TestNew(t *testing.T) {
	l := New(&printer{})
	testLog(l)
	testLogf(l)
}
