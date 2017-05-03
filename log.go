// Package log provides a log interface
package log

// Logger is a generic logging interface
type Logger interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
}
