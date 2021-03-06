package log

import (
	"file_manager/src/common/log/hooks"
	"file_manager/src/configs"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	logFile    = "./test.log"
	MaxSize    = 500 // megabytes
	MaxBackups = 3
	MaxAge     = 30   //days
	Compress   = true // disabled by default
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
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeTime:   SyslogTimeEncoder,
		EncodeLevel:  CustomLevelEncoder,
	}

	var encoder zapcore.Encoder
	var writeSyncer zapcore.WriteSyncer
	var level zapcore.Level
	cf := configs.Get()

	if cf.AppEnv == "dev" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
		writeSyncer = zapcore.AddSync(os.Stderr)
		level = zap.DebugLevel
	} else if cf.AppEnv == "prod" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
		writeSyncer = getWriteSyncer()
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	options := make([]zap.Option, 0)
	hookProcessor := hooks.NewHookProcessor(configs.Get().AppEnv)
	hook := zap.Hooks(func(entry zapcore.Entry) error {
		if entry.Level.String() == "error" {
			hookProcessor.ProcessEvent(entry)
		}
		return nil
	})
	options = append(options, hook)

	return &logger{
		hookProcessor: hookProcessor,
		zap:           zap.New(core, zap.AddCaller(), zap.AddCallerSkip(callerSkip)).WithOptions(options...).Sugar(),
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

func getWriteSyncer() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    MaxSize,
		MaxAge:     MaxAge,
		MaxBackups: MaxBackups,
		Compress:   Compress,
	}
	return zapcore.AddSync(lumberJackLogger)

}
