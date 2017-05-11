// Package log provides a log interface
package log

// Logger is a generic logging interface
type Logger interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
}

var (
	// The global default logger
	DefaultLogger = &noOpLogger{}
)

// noOpLogger is used as a placeholder for the default logger
type noOpLogger struct{}

func (n *noOpLogger) Log(v ...interface{}) {}

func (n *noOpLogger) Logf(format string, v ...interface{}) {}
