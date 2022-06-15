package logger

// Allias for fields
type Fields map[string]interface{}

// Extended logger interface supporting fields
type FieldLogger interface {
	WithField(key string, value interface{}) FieldLogger
	WithFields(fields Fields) FieldLogger
	WithError(err error) FieldLogger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}

type Logger struct {
	FieldLogger
}

func NewLogger(fl FieldLogger) *Logger {

	var l = &Logger{}
	if fl != nil {
		l.FieldLogger = fl
	} else {
		l.FieldLogger = DefaultLogger
	}
	return l
}
