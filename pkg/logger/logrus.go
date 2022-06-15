package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var DefaultLogger = &logrusLoggerWrapper{
	&logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	},
}

// In order to make logrus.Entry to conform interfaces
type logrusEntryWrapper struct {
	*logrus.Entry
}

func (e *logrusEntryWrapper) WithField(key string, value interface{}) FieldLogger {
	e.Entry = e.Entry.WithField(key, value)
	return e
}
func (e *logrusEntryWrapper) WithFields(fields Fields) FieldLogger {
	e.Entry = e.Entry.WithFields(map[string]interface{}(fields))
	return e
}
func (e *logrusEntryWrapper) WithError(err error) FieldLogger {
	e.Entry = e.Entry.WithError(err)
	return e
}

// In order to make logrus conform interface
type logrusLoggerWrapper struct {
	*logrus.Logger
}

func (l *logrusLoggerWrapper) WithField(key string, value interface{}) FieldLogger {
	return &logrusEntryWrapper{l.Logger.WithField(key, value)}
}
func (l *logrusLoggerWrapper) WithFields(fields Fields) FieldLogger {
	return &logrusEntryWrapper{l.Logger.WithFields(map[string]interface{}(fields))}
}
func (l *logrusLoggerWrapper) WithError(err error) FieldLogger {
	return &logrusEntryWrapper{l.Logger.WithError(err)}

}
