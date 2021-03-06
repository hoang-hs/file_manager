package log

var globalLogger Logging

func RegisterGlobal(logger Logging) {
	globalLogger = logger
}

func GetGlobalLog() Logging {
	return globalLogger
}

func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

func Infof(msg string, args ...interface{}) {
	globalLogger.Infof(msg, args...)
}

func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

func Errorf(msg string, args ...interface{}) {
	globalLogger.Errorf(msg, args...)
}

func Fatal(args ...interface{}) {
	globalLogger.Fatal(args...)
}

func Fatalf(msgFormat string, args ...interface{}) {
	globalLogger.Fatalf(msgFormat, args...)
}
