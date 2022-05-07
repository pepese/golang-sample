# gorm

```bash
$ go mod init github.com/pepese/golang-sample/gorm
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/mysql
```

## 概要

- [GORM v1 と v2 のソースコードリーディングしてみた](https://future-architect.github.io/articles/20210729a/)
- GORM v1: https://github.com/jinzhu/gorm
- GORM v2: https://github.com/go-gorm/gorm
  - Docs: https://gorm.io/docs/
  - Docs(日本語): https://gorm.io/ja_JP/docs/index.html

ここでは GORM v2 を扱います。

## 実行方法

```bash
$ docker-compose up -d
$ go run main.go
&{0   2022-05-07 11:51:17.681 +0900 JST 2022-05-07 11:51:17.681 +0900 JST <nil>}
```
