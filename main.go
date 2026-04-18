package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトミドルウェア（loggerとrecovery）を含むGinルーターを作成
	r := gin.Default()

	// シンプルなGETエンドポイントを定義
	r.GET("/ping", func(c *gin.Context) {
		// JSONレスポンスを返す
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// ポート8080でサーバーを起動（デフォルト）
	// サーバーは0.0.0.0:8080でリッスンします（Windowsではlocalhost:8080）
	r.Run()
}
