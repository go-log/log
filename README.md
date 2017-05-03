# Log [![GoDoc](https://godoc.org/github.com/go-log/log?status.svg)](https://godoc.org/github.com/go-log/log)

Log is a logging interface for Go. That's it. Pass around the interface.

## Rationale

Users want to standardise logging. Sometimes libraries log. We leave the underlying logging implementation to the user 
while allowing libraries to log by simply expecting something that satisfies the Logger interface. This leaves 
the user free to pre-configure structure, output, etc.

## Example

Here's a logger that uses walrus and logs with predefined fields.

```go
import (
	"github.com/go-log/log"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	*logrus.Entry
}

func (l *logrusLogger) Log(v ...interface{}) {
	l.Entry.Print(v...)
}

func (l *logrusLogger) Logf(format string, v ...interface{}) {
	l.Entry.Printf(format, v...)
}

func WithFields(f logrus.Fields) log.Logger {
	return &logrusLogger{logrus.WithFields(f)}	
}
```

The `WithFields` func returns a struct that satisfies the Logger interface.
