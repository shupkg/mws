package mws

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Printf(format string, args ...interface{})
}

func SetLogger(logger Logger) {
	log = logger
}

var log Logger = emptyLogger{}

type emptyLogger struct{}

func (l emptyLogger) Debugf(string, ...interface{}) {}
func (l emptyLogger) Infof(string, ...interface{})  {}
func (l emptyLogger) Warnf(string, ...interface{})  {}
func (l emptyLogger) Errorf(string, ...interface{}) {}
func (l emptyLogger) Printf(string, ...interface{}) {}
