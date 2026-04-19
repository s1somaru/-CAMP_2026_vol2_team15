package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-server/handlers"
	"go-server/middleware"
)

func main() {
	r := gin.Default()

	// セッションの設定（"secret"はハッカソン用の適当な文字列）
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// CORS設定（フロントのポートに合わせて調整）
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5500"} // フロントのURL
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// 誰でも叩けるAPI
	r.POST("/api/login", handlers.Login)
	r.POST("/api/logout", handlers.Logout)

	// ログインが必要なAPIグループ
	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthCheck())
	{
		authGroup.GET("/company", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "これはログイン中のみ見える企業情報である"})
		})
	}

	r.Run(":8080")
}