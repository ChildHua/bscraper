package app

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("scrap.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	return zapcore.AddSync(file)
}
