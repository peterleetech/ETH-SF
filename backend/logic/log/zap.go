package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"runtime"
)

var logger *zap.Logger

func init() {
	var err error

	if false {
		developmentConfig := zap.NewDevelopmentConfig()

		if _, file, _, ok := runtime.Caller(0); ok {
			basePath := filepath.Dir(filepath.Dir(filepath.Dir(file))) + "/"
			developmentConfig.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
				rel, err := filepath.Rel(basePath, caller.File)
				if err != nil {
					encoder.AppendString(caller.FullPath())
				} else {
					encoder.AppendString(fmt.Sprintf("%s:%d", rel, caller.Line))
				}
			}
		}

		logger, err = developmentConfig.Build()
		if err != nil {
			panic(err)
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}
}

func GetLogger() *zap.Logger {
	return logger
}
