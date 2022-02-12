package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album レコード記録用の構造体
type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums 今回使用するテストデータ構造体
var albums = []album{
	{ID: "1", Title: "Blue", Artist: "john", Price: 100.0},
	{ID: "2", Title: "Red", Artist: "Mike", Price: 50},
}

// main メインエントリ
func main() {
	r := gin.Default()
	r.GET("/albums", getAlbums)

	r.Run(":3030")
}

// getAlbums アルバム取得
func getAlbums(gc *gin.Context) {
	gc.IndentedJSON(http.StatusOK, albums)
}
