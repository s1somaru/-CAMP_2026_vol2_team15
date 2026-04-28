package main

import (
	"go-server/database"
	"go-server/handlers"
	"go-server/middleware"
	"go-server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// データベース接続の初期化
	database.InitDB()

	// ここでテーブルを自動作成・更新する指示を出す
	err := database.DB.AutoMigrate(&models.Company{})
	if err != nil {
		panic("マイグレーションに失敗した: " + err.Error())
	}

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
		//authGroup.GET("/company", handlers.GetCompanies)
		// 企業情報登録API
		//authGroup.POST("/company", handlers.CreateCompany)
	}

	r.Run(":8080")
}
