package log

import (
	"file_manager/configs"
	"file_manager/internal/common/log/hooks"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

const (
	OutputModeConsole = "console"
	OutputModeJson    = "json"
)

const callerSkip = 2

type logger struct {
	hookProcessor *hooks.HookProcessor
	zap           *zap.SugaredLogger
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func NewLogger() (Logging, error) {
	var level zapcore.Level
	var outputMode string
	cf := configs.Get()
	if cf.AppEnv == "dev" {
		outputMode = OutputModeConsole
		level = zap.DebugLevel
	} else if cf.AppEnv == "prod" {
		outputMode = OutputModeJson
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeTime:   SyslogTimeEncoder,
		EncodeLevel:  CustomLevelEncoder,
	}

	zapLogger, err := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         outputMode,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    encoderConfig,
	}.Build()
	if err != nil {
		panic(err)
	}

	options := make([]zap.Option, 0)
	options = append(options, zap.AddCallerSkip(callerSkip))

	hookProcessor := hooks.NewHookProcessor(configs.Get().AppEnv)
	hook := zap.Hooks(func(entry zapcore.Entry) error {
		/*
			if entry.Level.String() == "error" {
				hookProcessor.ProcessEvent(entry)
			}
		*/
		hookProcessor.ProcessEvent(entry)
		return nil
	})
	options = append(options, hook)

	return &logger{
		hookProcessor: hookProcessor,
		zap:           zapLogger.WithOptions(options...).Sugar(),
	}, nil
}

func (l *logger) Info(args ...interface{}) {
	l.zap.Info(args...)
}

func (l *logger) Infof(msg string, args ...interface{}) {
	l.zap.Infof(msg, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.zap.Error(args...)
}

func (l *logger) Errorf(msg string, args ...interface{}) {
	l.zap.Errorf(msg, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.zap.Fatal(args...)
}

func (l *logger) Fatalf(msgFormat string, args ...interface{}) {
	l.zap.Fatalf(msgFormat, args...)
}

func (l *logger) GetZap() *zap.SugaredLogger {
	return l.zap
}
