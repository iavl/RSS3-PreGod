package logger

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger/engine"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger
var ShortcutLogger *zap.SugaredLogger

// Some simple encapsulations are made for the upper layer.
// The Sugared mode of the zap library is used by default.
// You can customize the encapsulation here to use other log libraries.
func Setup() error {
	var err error
	Logger, err = engine.InitZapLogger(config.Config.Logger)

	ShortcutLogger = Logger.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar() // to skip the caller of this function.

	if err != nil {
		return err
	}

	return nil
}

// Some shortcuts:

func Debug(args ...interface{}) {
	ShortcutLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	ShortcutLogger.Debugf(format, args...)
}

func Error(args ...interface{}) {
	ShortcutLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	ShortcutLogger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	ShortcutLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	ShortcutLogger.Fatalf(format, args...)
}

func Info(args ...interface{}) {
	ShortcutLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	ShortcutLogger.Infof(format, args...)
}

func Panic(args ...interface{}) {
	ShortcutLogger.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	ShortcutLogger.Panicf(format, args...)
}

func Warn(args ...interface{}) {
	ShortcutLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	ShortcutLogger.Warnf(format, args...)
}
