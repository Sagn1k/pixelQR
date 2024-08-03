package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

func NewLogger() *zap.Logger {
    config := zap.NewProductionEncoderConfig()
    config.TimeKey = "timestamp"
    config.EncodeTime = zapcore.ISO8601TimeEncoder
    config.EncodeLevel = zapcore.CapitalLevelEncoder

    fileEncoder := zapcore.NewJSONEncoder(config)
    consoleEncoder := zapcore.NewConsoleEncoder(config)

    // Ensure the logs directory exists
    if err := os.MkdirAll("logs", os.ModePerm); err != nil {
        panic(err)
    }

    logFile, err := os.Create("logs/server.log")
    if err != nil {
        panic(err)
    }
    logWriter := zapcore.AddSync(logFile)

    core := zapcore.NewTee(
        zapcore.NewCore(fileEncoder, logWriter, zapcore.DebugLevel),
        zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
    )

    logger := zap.New(core, zap.AddCaller())
    return logger
}
