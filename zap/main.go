package main

import (
	"github.com/pepese/golang-sample/zap/app"
	"go.uber.org/zap"
)

func main() {

	logger := app.Logger()
	defer logger.Sync()

	logger.Info(
		"こんにちは 世界",
		zap.String("type", "greeting"),
		zap.String("target", "world"),
	)

	// テンポラリでログエントリを追加：zap.String() を隠蔽したい場合
	tmpLogger := app.LoggerWithKeyValue("tmp-key", "tmp-value")
	tmpLogger.Info(
		"こんにちは 世界2",
	)

	// sugaredLoggerは少しだけ遅いが記述量は減る：zap.String() の省略
	sugaredLogger := logger.Sugar()
	sugaredLogger.Errorw(
		"こんにちは 世界3",
		"type", "greeting",
		"target", "world",
	)

	// Fatalを呼び出すとスタックトレースを出力してプロセスは終了する
	logger.Fatal(
		"さようなら 世界",
		zap.String("type", "greeting"),
		zap.String("target", "world"),
	)
}

// 出力結果：
// {"level":"info","timestamp":"2022-05-06T17:21:02.391+0900","caller":"zap/main.go:13","message":"こんにちは 世界","permanent-key":"permanent-value","type":"greeting","target":"world"}
// {"level":"info","timestamp":"2022-05-06T17:21:02.391+0900","caller":"zap/main.go:21","message":"こんにちは 世界2","permanent-key":"permanent-value","tmp-key":"tmp-value"}
// {"level":"error","timestamp":"2022-05-06T17:21:02.391+0900","caller":"zap/main.go:27","message":"こんにちは 世界3","permanent-key":"permanent-value","type":"greeting","target":"world","stacktrace":"main.main\n\t/Users/pepese/workspace/github/pepese/golang-sample/zap/main.go:27\nruntime.main\n\t/Users/pepese/.asdf/installs/golang/1.18.1/go/src/runtime/proc.go:250"}
// {"level":"fatal","timestamp":"2022-05-06T17:21:02.391+0900","caller":"zap/main.go:34","message":"さようなら 世界","permanent-key":"permanent-value","type":"greeting","target":"world","stacktrace":"main.main\n\t/Users/pepese/workspace/github/pepese/golang-sample/zap/main.go:34\nruntime.main\n\t/Users/pepese/.asdf/installs/golang/1.18.1/go/src/runtime/proc.go:250"}
// exit status 1
