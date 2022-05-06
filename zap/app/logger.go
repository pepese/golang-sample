package app

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger     *zap.Logger
	loggerOnce sync.Once
)

func Logger() *zap.Logger {
	return getLogger()
}

func getLogger() *zap.Logger {
	loggerOnce.Do(func() {
		// 設定が面倒な場合
		// logger := zap.NewDevelopmentConfig()
		// logger := zap.NewProductionConfig()
		// などが用意されています
		tmpLogger, err := zap.Config{
			Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "message",
				LevelKey:       "level",
				TimeKey:        "timestamp",
				NameKey:        "name",
				CallerKey:      "caller",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}.Build(
			zap.AddCaller(), // zapの呼び出し元のファイル名、行番号、および関数名で各メッセージに注釈を付けるようにロガーを構成します。 WithCallerも参照してください。
			zap.Fields(
				zap.String("permanent-key", "permanent-value"), // 追加のエントリです
			),
		)

		if err != nil {
			panic(err)
		}

		logger = tmpLogger
	})

	return logger
}

func LoggerWithKeyValue(key string, value string) *zap.Logger {
	tmpLogger := getLogger()
	tmpLogger = tmpLogger.With(
		zap.String(key, value),
	)
	return tmpLogger
}
