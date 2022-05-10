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

## context.Context 利用方針

- 主に goroutine や API コールの際に、タイムアウト制御を適切に行うために使用
- UserID 、 JWT など、リクエストスコープに限定される値を Get/Set するときのみ Value を使う
- Go 1.7 以降ではリクエストから context.Context を取得可能（ `ctx := r.Context()` ）
- その他の用途、受け渡しでは絶対使わない

↓

- 無駄に構造体に context を入れない
- Context を適切にハンドラに渡す
  - タイムアウトは ServeHTTP で行う（ Middleware でも可）
  - `func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request)`
- 各種クライアントや Logger は Handler のレシーバを定義してその中に入れる
  - `function (app *App) GetItemHandler(ctx context.Context, ...)`

```go
type AppHandler struct {
    Impl func(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

func (a AppHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    ctx, cancel := context.WithTimeout(ctx, TimeOutLimit)
    defer cancel()
    a.Impl(ctx, w, r)
}
```

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