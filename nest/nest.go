// Package nest allows users to use a Logger interface to
// create another Logger interface.
package nest

import (
	"fmt"

	"github.com/go-log/log"
)

type logger struct {
	logger log.Logger
	values []interface{}
}

func (logger *logger) Log(v ...interface{}) {
	logger.logger.Log(append(logger.values, v...)...)
}

func (logger *logger) Logf(format string, v ...interface{}) {
	logger.logger.Log(append(logger.values, fmt.Sprintf(format, v...))...)
}

func New(log log.Logger, v ...interface{}) *logger {
	return &logger{
		logger: log,
		values: v,
	}
}

func Newf(log log.Logger, format string, v ...interface{}) *logger {
	return &logger{
		logger: log,
		values: []interface{}{fmt.Sprintf(format, v...)},
	}
}
