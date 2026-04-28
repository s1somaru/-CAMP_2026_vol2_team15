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

	// --- 公開API ---
	public := r.Group("/api")
	{
	    public.POST("/signup", handlers.Signup)
	    public.POST("/login", handlers.Login)
	}
	
	// --- 保護API（AuthCheckを適用） ---
	protected := r.Group("/api")
	protected.Use(middleware.AuthCheck())
	{
	    protected.POST("/logout", handlers.Logout) // ここに入れる
	    protected.GET("/companies", handlers.GetCompanies)
	    protected.POST("/companies", handlers.CreateCompany)
	}

	r.Run(":8080")
}