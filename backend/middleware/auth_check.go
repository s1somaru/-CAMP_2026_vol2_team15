package middleware

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userEmail := session.Get("user_email")

		if userEmail == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ログインが必要である"})
			c.Abort() // 以降のハンドラーを実行させない
			return
		}
		c.Next() // 次の処理へ
	}
}