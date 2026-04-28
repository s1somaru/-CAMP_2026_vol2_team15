package handlers

import (
	"net/http"
	"go-server/database"
	"go-server/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 1. リクエストボディのバリデーション
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力データが不正である"})
		return
	}

	// 2. ユーザーの重複チェック（任意だが推奨）
	var existingUser models.User
	if err := database.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "このユーザー名は既に使用されている"})
		return
	}

	// 3. データベースへの保存
	newUser := models.User{
		Username: input.Username,
		Password: input.Password, // 本来はハッシュ化すべきだが、まずはそのまま保存する
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーの登録に失敗した"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ユーザー登録が完了した"})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "入力が不正である"})
		return
	}

	var user models.User
	// DBからユーザーを探す
	result := database.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil || user.Password != input.Password {
		c.JSON(401, gin.H{"error": "ユーザー名またはパスワードが違う"})
		return
	}

	// ここでCookie（セッション）を保存する
	session := sessions.Default(c)
	session.Set("user_email", user.Username) // AuthCheckでこれを確認している
	session.Save()

	c.JSON(200, gin.H{"message": "ログインに成功した"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	
	// セッションからユーザー情報を削除する
	session.Delete("user_email")
	
	// セッションの中身を空にして即時保存（これによりブラウザのCookieが無効になる）
	session.Options(sessions.Options{MaxAge: -1}) 
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "ログアウトに成功した"})
}