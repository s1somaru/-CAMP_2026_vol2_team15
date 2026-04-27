package main

import (
	"go-server/database"
	"go-server/handlers"
	"go-server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// データベース接続の初期化
	database.InitDB()

	r := gin.Default()

	// セッションの設定
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// CORS設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5500"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// 公開API
	r.POST("/api/login", handlers.Login)
	r.POST("/api/logout", handlers.Logout)

	// 認証が必要なAPIグループ
	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthCheck())
	{
		// 企業一覧取得API
		authGroup.GET("/company", handlers.GetCompanies)
	}

	r.Run(":8080")
}