package log

import log "github.com/sirupsen/logrus"

func do(f func(...any), args ...any) {
	if writeLogs {
		f(args...)
	}
}

func dof(f func(string, ...any), format string, args ...any) {
	if writeLogs {
		f(format, args...)
	}
}

func Panic(args ...interface{}) {
	do(log.Panic, args...)
}

func Panicf(format string, args ...interface{}) {
	dof(log.Panicf, format, args...)
}

func Fatal(args ...interface{}) {
	do(log.Fatal, args...)
}

func Fatalf(format string, args ...interface{}) {
	dof(log.Fatalf, format, args...)
}

func Error(args ...interface{}) {
	do(log.Error, args...)
}

func Errorf(format string, args ...interface{}) {
	dof(log.Errorf, format, args...)
}

func Warn(args ...interface{}) {
	do(log.Warn, args...)
}

func Warnf(format string, args ...interface{}) {
	dof(log.Warnf, format, args...)
}

func Info(args ...interface{}) {
	do(log.Info, args...)
}

func Infof(format string, args ...interface{}) {
	dof(log.Infof, format, args...)
}

func Debug(args ...interface{}) {
	do(log.Debug, args...)
}

func Debugf(format string, args ...interface{}) {
	dof(log.Debugf, format, args...)
}

func Trace(args ...interface{}) {
	do(log.Trace, args...)
}

func Tracef(format string, args ...interface{}) {
	dof(log.Tracef, format, args...)
}
