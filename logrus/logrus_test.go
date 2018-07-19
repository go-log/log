package logrus

import (
	"testing"

	"github.com/go-log/log"
	"github.com/sirupsen/logrus"
)

func testLog(l log.Logger) {
	l.Log("test")
}

func testLogf(l log.Logger) {
	l.Logf("%s", "test")
}

func TestLogLogger(t *testing.T) {
	var l log.Logger

	// basic logger
	l = New()
	testLog(l)
	testLogf(l)

	// with fields
	l = WithFields(logrus.Fields{
		"test": "test",
	})
	testLog(l)
	testLogf(l)
}
