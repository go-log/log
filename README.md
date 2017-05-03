# Log [![GoDoc](https://godoc.org/github.com/go-log/log?status.svg)](https://godoc.org/github.com/go-log/log)

Log is a logging interface for Go. That's it. Pass around the interface.

## Rationale

Users want to standardise logging. Sometimes libraries log. We leave the underlying logging implementation to the user 
while allowing libraries to log by simply expecting something that satisfies the Logger interface. This leaves 
the user free to pre-configure structure, output, etc.
