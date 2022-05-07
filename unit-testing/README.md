# unit-testing

```bash
$ go mod init github.com/pepese/golang-sample/unit-testing
$ go get -u github.com/golang/mock/gomock
$ go install github.com/golang/mock/mockgen@v1.6.0
$ mockgen -version
v1.6.0
```

## 概要

[gomock](https://github.com/golang/mock) を利用します。  
以下でモックのコードを生成します。  
なお、 `hoge` というディレクトリ・パッケージがあった場合、そのディレクトリ内に `mock_hoge` というディレクトリを作成してモックコードを出力のが慣例のようです。

```bash
$ mockgen -source app/sample.go -destination app/mock_app/mock_sample.go
```

## 実行

モックの生成は以下の通りです。

```bash
$ mockgen -source app/sample.go -destination app/mock_app/mock_sample.go
```

以下で単体テストを実行します。

```bash
# 一部のテスト
# go test -run <テスト関数名>
# go test <パッケージ名>
# go test <パッケージ名>/... # パッケージ以下全て
# go test <パッケージ名> -run <テスト関数名>
$ go test github.com/pepese/golang-sample/unit-testing/app -run TestSample
# 「-v」をつけると標準出力も表示される
$ go test github.com/pepese/golang-sample/unit-testing/app -v

# 全テスト
$ go test ./...

# カバレッジの標準出力
$ go test -v -cover ./...

# カバレッジのファイル出力
$ mkdir test; mkdir test/coverage
$ go test -coverprofile=./test/coverage/cover.out ./...
$ go tool cover -html=./test/coverage/cover.out -o ./test/coverage/cover.html
$ open ./test/coverage/cover.html
```