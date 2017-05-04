package log

import (
	"fmt"
	"testing"
)

type testLogger struct{}

func (t *testLogger) Log(v ...interface{}) {
	fmt.Print(v...)
}

func (t *testLogger) Logf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func testLog(l Logger) {
	l.Log("test\n")
}

func testLogf(l Logger) {
	l.Logf("%s\n", "test")
}

func TestLogger(t *testing.T) {
	l := new(testLogger)
	testLog(l)
	testLogf(l)
}
