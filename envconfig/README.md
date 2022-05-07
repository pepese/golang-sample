# envconfig

```bash
$ go mod init github.com/pepese/golang-sample/envconfig
$ go get -u github.com/kelseyhightower/envconfig
```

## 概要

- https://github.com/kelseyhightower/envconfig

以下ではないので注意してください

- https://github.com/vrischmann/envconfig

## 実行

```bash
$ go run main.go
golang-sample

$ APP_NAME=hoge go run main.go
hoge
```