# zap

```bash
$ go mod init github.com/pepese/golang-sample/zap
$ go get -u go.uber.org/zap
```

## 概要

zap は golang のロギングライブラリとして最も多く利用されているものの 1 つです。  
ここでは簡単な使い方について記載します。

- Githug: https://github.com/uber-go/zap
- Docs: 
  - https://pkg.go.dev/go.uber.org/zap
  - https://pkg.go.dev/go.uber.org/zap/zapcore

## 用語

- `Logger` ： 高速なロガーです
- `SugaredLogger` ： `Logger` に比べて低速ですが、設定の記述量を削減したロガーです

## 設定

### `zap.Config`

```go
type Config struct {
    // 出力される最小ログレベル。
    // Config.Level.SetLevelを呼び出すと、派生したすべてのロガーのログレベルが変更されます。
	Level AtomicLevel `json:"level" yaml:"level"`

    // ロガーが開発モードになります。
    // DPanicLevel の挙動が変更され、より自由にスタックトレースを取得できます。
	Development bool `json:"development" yaml:"development"`

    // 呼び出し元関数のファイル名・行番号のログ出力を停止します。
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`

    // 自動スタックトレースキャプチャを無効にします。
    // デフォルトでは、開発モードのWarnLevel以上のログと、本番モードのErrorLevel以上のログに対してキャプチャされます。
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`

    // サンプリングポリシーを設定します。
    // nil の場合、サンプリングは無効になります。
	Sampling *SamplingConfig `json:"sampling" yaml:"sampling"`

    // ロガーのエンコーディング（出力フォーマット）を設定します。
    // 有効な値は、「json」と「console」、および RegisterEncoder を介して登録されたサードパーティのエンコーディングです。
	Encoding string `json:"encoding" yaml:"encoding"`

    // エンコーダーのオプションを設定します。
    // 詳細は zapcore.EncoderConfig を参照してください。
	EncoderConfig zapcore.EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`

    // ログ出力先 URL またはファイルパスのリストです。
    // なお、「stdout」および「stderr」は os.Stdout および os.Stderr として解釈されます。
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`

    // 内部エラーログ出力先 URL またはファイルパスのリストです。
    // この設定は内部エラーにのみ影響することに注意してください。
	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`

    // ルートロガーに追加するフィールドのコレクションです。
	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`
}
```

### `zapcore.EncoderConfig`

```go
type EncoderConfig struct {
    // 各ログエントリに使用されるキーを設定します。
    // キーの設定が空の場合、そのエントリは省略されます。
	MessageKey     string `json:"messageKey" yaml:"messageKey"`
	LevelKey       string `json:"levelKey" yaml:"levelKey"`
	TimeKey        string `json:"timeKey" yaml:"timeKey"`
	NameKey        string `json:"nameKey" yaml:"nameKey"`
	CallerKey      string `json:"callerKey" yaml:"callerKey"`
	FunctionKey    string `json:"functionKey" yaml:"functionKey"`
	StacktraceKey  string `json:"stacktraceKey" yaml:"stacktraceKey"`
	SkipLineEnding bool   `json:"skipLineEnding" yaml:"skipLineEnding"`
	LineEnding     string `json:"lineEnding" yaml:"lineEnding"`
	// Configure the primitive representations of common complex types. For example, some users may want all time.Times serialized as floating-point seconds since epoch, while others may prefer ISO8601 strings.
	EncodeLevel    LevelEncoder    `json:"levelEncoder" yaml:"levelEncoder"`
	EncodeTime     TimeEncoder     `json:"timeEncoder" yaml:"timeEncoder"`
	EncodeDuration DurationEncoder `json:"durationEncoder" yaml:"durationEncoder"`
	EncodeCaller   CallerEncoder   `json:"callerEncoder" yaml:"callerEncoder"`
	// Unlike the other primitive type encoders, EncodeName is optional. The zero value falls back to FullNameEncoder.
	EncodeName NameEncoder `json:"nameEncoder" yaml:"nameEncoder"`

    // interface{} タイプのオブジェクト用にエンコーダーを構成します。
    // 指定しない場合、オブジェクトはjson.Encoderを使用してエンコードされます
	NewReflectedEncoder func(io.Writer) ReflectedEncoder `json:"-" yaml:"-"`

    // コンソールエンコーダーが使用するフィールドセパレーターを構成します。デフォルトはタブです。
	ConsoleSeparator string `json:"consoleSeparator" yaml:"consoleSeparator"`
}
```