package log

import "go.uber.org/zap"

type Logging interface {
	// Info uses fmt.Sprint to construct and log a message.
	Info(args ...interface{})

	// Infof uses fmt.Sprintf to log a template message.
	Infof(msgFormat string, args ...interface{})

	// Error uses fmt.Sprint to construct and log a message.
	Error(args ...interface{})

	// Errorf uses fmt.Sprintf to log a template message.
	Errorf(msgFormat string, args ...interface{})

	// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
	Fatal(args ...interface{})

	// Fatalf uses fmt.Sprintf to log a template message, then calls os.Exit.
	Fatalf(msgFormat string, args ...interface{})

	GetZap() *zap.SugaredLogger
}
