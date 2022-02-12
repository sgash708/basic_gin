# gin
## コード解説
### 全項目取得ハンドラ
```:go
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
```

#### `gin.Context`について
- en
> gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)

- ja
> gin.ContextはGinの最も重要な部分です。リクエストの詳細、JSONの検証、シリアライズなどを行います。(似たような名前ですが、これはGoの組み込みコンテキストパッケージとは異なります)。

#### JSONシリアライズ
- en
> Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
>The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
>Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.

- ja
> `Context.IndentedJSON`を呼び出して、構造体をJSONにシリアライズし、レスポンスに追加します。
> この関数の最初の引数は、クライアントに送信したいHTTPステータスコードです。ここでは、200 OK を示すために、net/http パッケージの StatusOK 定数を渡しています。
>なお、`Context.IndentedJSON`を`Context.JSON`の呼び出しに置き換えると、よりコンパクトなJSONを送信することができる。実際には、インデントされた形の方がデバッグの際に作業しやすく、サイズの差も通常は小さいです。
