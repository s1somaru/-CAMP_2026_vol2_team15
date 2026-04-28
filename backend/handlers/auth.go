package handlers

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 簡易的なユーザー情報
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// テスト用データ
var testUser = User{Email: "test@example.com", Password: "password123"}

func Login(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストが不正である"})
		return
	}

	// メアドとパスワードの照合（簡易版）
	if req.Email != testUser.Email || req.Password != testUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証に失敗した"})
		return
	}

	// セッションに保存
	session := sessions.Default(c)
	session.Set("user_email", req.Email)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "ログイン成功"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "ログアウト成功"})
}