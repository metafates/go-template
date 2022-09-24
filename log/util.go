package log

import log "github.com/sirupsen/logrus"

func Panic(args ...interface{}) {
	if writeLogs {
		log.Panic(args...)
	}
}

func Panicf(format string, args ...interface{}) {
	if writeLogs {
		log.Panicf(format, args...)
	}
}

func Fatal(args ...interface{}) {
	if writeLogs {
		log.Fatal(args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	if writeLogs {
		log.Fatalf(format, args...)
	}
}

func Error(args ...interface{}) {
	if writeLogs {
		log.Error(args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if writeLogs {
		log.Errorf(format, args...)
	}
}

func Warn(args ...interface{}) {
	if writeLogs {
		log.Warn(args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if writeLogs {
		log.Warnf(format, args...)
	}
}

func Info(args ...interface{}) {
	if writeLogs {
		log.Info(args...)
	}
}

func Infof(format string, args ...interface{}) {
	if writeLogs {
		log.Infof(format, args...)
	}
}

func Debug(args ...interface{}) {
	if writeLogs {
		log.Debug(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if writeLogs {
		log.Debugf(format, args...)
	}
}

func Trace(args ...interface{}) {
	if writeLogs {
		log.Trace(args...)
	}
}

func Tracef(format string, args ...interface{}) {
	if writeLogs {
		log.Tracef(format, args...)
	}
}
