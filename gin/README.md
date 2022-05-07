# gin

```bash
$ go mod init github.com/pepese/golang-sample/gin
$ go get -u github.com/gin-gonic/gin
$ go get github.com/rs/xid
```

## 設計

- `infrastructure/server`
  - gin.Engine の作成
  - middleware の設定
    - `requestScope` 関数にてリクエストスコープデータとして traceID を `context.Context` へ追加します
    - `accessLog` 関数にてアクセスログを出力します
- `adapter/controller` ： 以下のみを実施
  - リクエストのパースとバリデーション
    - JSON パースやバリデーションなどの構造体タグは usecase の *Request/*Response にのみ付与する（domain entity には決して付与しない）
  - usecase の呼び出し
  - レスポンスの view を生成
- `usecase`
  - ビジネスロジックの実行

## 実行

```bash
# 起動
$ go run main.go

# リクエスト例
$ curl localhost:8080/api/v1/users
{}
$ curl localhost:8080/api/v1/users/1
{}
$ curl -X POST localhost:8080/api/v1/users
{}
$ curl -X PUT localhost:8080/api/v1/users/1
{}
$ curl -X DELETE localhost:8080/api/v1/users/1
null
```