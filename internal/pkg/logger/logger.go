package logger

type Logger interface {
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
}
