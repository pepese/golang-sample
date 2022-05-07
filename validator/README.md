# validator

```bash
$ go mod init github.com/pepese/golang-sample/validator
$ go get github.com/go-playground/validator/v10
```

## 概要

- https://github.com/go-playground/validator

構造体タグベースで入力チェックを行うシンプルなライブラリです。  
対抗馬として以下があります。

- https://github.com/asaskevich/govalidator

## 実行

```bash
$ go run main.go
```

全てのバリデーションルールをチェックせず、 1 つでも違反があった時点で検査は終了する模様です。