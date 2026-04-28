package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DBエラーを共通で処理する関数
func HandleDBError(c *gin.Context, err error) bool {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "データベースエラー"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラー: " + err.Error()})
		}
		return true // エラー時はtrue
	}
	return false // エラーなし
}
