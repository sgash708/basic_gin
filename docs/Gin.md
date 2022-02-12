# gin
## コード解説
## 0. データ作成
```
// album レコード記録用の構造体
type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
```

- en
> Beneath the package declaration, paste the following declaration of an album struct. You’ll use this to store album data in memory.
> Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.

- ja
> パッケージ宣言の下に、以下のアルバム構造体の宣言を貼り付けます。これを利用して、アルバムデータをメモリに格納します。
> json: "artist "などの構造体タグは、構造体の内容をJSONにシリアライズする際に、フィールド名をどうするかを指定するものです。このタグがない場合、JSONでは構造体の大文字のフィールド名が使用されます（JSONではあまり一般的でないスタイルです）。

## 1. 全項目取得ハンドラAPI
### 1.1.1 `gin.Context`について
```:go
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
```

- en
> gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)

- ja
> gin.ContextはGinの最も重要な部分です。リクエストの詳細、JSONの検証、シリアライズなどを行います。(似たような名前ですが、これはGoの組み込みコンテキストパッケージとは異なります)。

### 1.1.2 JSONシリアライズ
- en
> Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
>The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
>Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.

- ja
> `Context.IndentedJSON`を呼び出して、構造体をJSONにシリアライズし、レスポンスに追加します。
> この関数の最初の引数は、クライアントに送信したいHTTPステータスコードです。ここでは、200 OK を示すために、net/http パッケージの StatusOK 定数を渡しています。
>なお、`Context.IndentedJSON`を`Context.JSON`の呼び出しに置き換えると、よりコンパクトなJSONを送信することができる。実際には、インデントされた形の方がデバッグの際に作業しやすく、サイズの差も通常は小さいです。

### 1.2 APIエンドポイントの割り当て
```
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)

    // WARN: Docker
    router.Run(":3030")

    // On HostMachine
    // router.Run("localhost:3030")
}
```

```:bash
$ curl http://localhost:3030/albums
[
    {
        "id": "1",
        "title": "Blue",
        "artist": "john",
        "price": 100
    },
    {
        "id": "2",
        "title": "Red",
        "artist": "Mike",
        "price": 50
    }
]
```