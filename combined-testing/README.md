# combined-testing

```bash
$ brew install postman --cask
$ brew install newman
```

## 概要

postman で作成したテストシナリオを JSON 形式で出力し、 newman でコマンドライン実行します。  
ここでは、本リポジトリの `gin` に対してリクエストを発行するサンプルとなっています。

```bash
# newman run -e <環境変数ファイル>.json <コレクションファイル>.json
$ newman run -e SampleEnvironment.postman_environment.json SampleCollection.postman_collection.json 
newman

SampleCollection

→ List Users
  GET http://localhost:8080/api/v1/users [200 OK, 124B, 81ms]

→ Get User
  GET http://localhost:8080/api/v1/users/1 [200 OK, 124B, 6ms]

→ Create User
  POST http://localhost:8080/api/v1/users [200 OK, 124B, 6ms]

→ Update User
  PUT http://localhost:8080/api/v1/users/1 [200 OK, 124B, 5ms]

→ Delete User
  DELETE http://localhost:8080/api/v1/users/1 [200 OK, 126B, 6ms]

┌─────────────────────────┬──────────────────┬──────────────────┐
│                         │         executed │           failed │
├─────────────────────────┼──────────────────┼──────────────────┤
│              iterations │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│                requests │                5 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│            test-scripts │                0 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│      prerequest-scripts │                0 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│              assertions │                0 │                0 │
├─────────────────────────┴──────────────────┴──────────────────┤
│ total run duration: 178ms                                     │
├───────────────────────────────────────────────────────────────┤
│ total data received: 12B (approx)                             │
├───────────────────────────────────────────────────────────────┤
│ average response time: 20ms [min: 5ms, max: 81ms, s.d.: 30ms] │
└───────────────────────────────────────────────────────────────┘
```

## コレクションの作成

Postman の GUI でテストシナリオを作成します。

- GUI を起動して「 skip and go to the app 」（アカウント作成してもいいです）
- 左メニューから「 Collections 」を選択してコレクションを作成
- 初回は作成したコレクション下の「 Add a request 」でリクエストを作成
  - その後はコレクション右の「...」から「 Add request 」でリクエストを作成
- リクエストを作成したらコレクション右の「...」から「 Export 」を選択してコレクションを JSON 形式で出力

リクエストの設定項目には以下のようなものがあります。

|項目|内容|
|:---|:---|
|Authorization|Basic AuthやOAuth2など、リクエストの認証方法の指定|
|Headers|content-typeやAuthorizationなど、リクエストヘッダの指定|
|Body|PostやPUTなどの場合に指定するリクエストボディ|
|Pre-request Script|リクエスト前に実行するJSコード|
|Tests|レスポンスペイロードの妥当性を検証するJSコード|

## 環境変数の設定・作成

Postman の GUI で設定可能です。  

- 右上の目のようなマークから「 ENVIRONMENS 」 -> 「 Add 」を選択
- 右上の「...」 -> 「 Export 」を選択して環境変数を JSON 形式で出力

なお、環境変数は「 `{{環境変数}}` 」の形式で埋め込むことで利用できます。