package log

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var logger, _ = zap.NewProduction()

// var (
// 	Debug = logger.Debug
// 	Info  = logger.Info
// 	Warn  = logger.Warn
// 	Error = logger.Error
// 	Fatal = logger.Fatal()
// )

var (
	Debug func(string, ...zapcore.Field)
	Info  func(string, ...zapcore.Field)
	Warn  func(string, ...zapcore.Field)
	Error func(string, ...zapcore.Field)
	Fatal func(string, ...zapcore.Field)
)

func init() {
	initLogger(false)
}

func initLogger(debug bool) {
	var (
		logger *zap.Logger
		err    error
	)
	if debug {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalln("init logger failed, exiting")
	}
	zap.ReplaceGlobals(logger)

	Debug = zap.L().Debug
	Info = zap.L().Info
	Warn = zap.L().Warn
	Error = zap.L().Error
	Fatal = zap.L().Fatal
}

func SetDebugMode() {
	initLogger(true)
}
