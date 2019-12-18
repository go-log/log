package log

import (
	"testing"

	"github.com/go-log/log"
)

func testLog(l log.Logger) {
	l.Log("test")
}

func testLogf(l log.Logger) {
	l.Logf("%s", "test")
}

func TestLogLogger(t *testing.T) {
	l := New()
	testLog(l)
	testLogf(l)
}
