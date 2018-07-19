// Package logrus allows the creation of Loggers based on
// logrus.StandardLogger().
//
// Deprecated: Use github.com/go-log/log/print instead.
package logrus

import (
	"github.com/go-log/log"
	"github.com/go-log/log/print"
	"github.com/sirupsen/logrus"
)

// WithFields creates a new logger based on logrus.WithFields().
//
// Deprecated: Replace with:
//
//   print.New(logrus.WithFields(f))
func WithFields(f logrus.Fields) log.Logger {
	return print.New(logrus.WithFields(f))
}

// WithFields creates a new logger based on logrus.StandardLogger().
//
// Deprecated: Replace with:
//
//   print.New(logrus.StandardLogger())
func New() log.Logger {
	return print.New(logrus.StandardLogger())
}
