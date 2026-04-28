package handlers

import (
	"go-server/database"
	"go-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateTodo(c *gin.Context) {

	var req models.Todo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ここで実際の保存処理を行う（例: データベースに保存）
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
		"data":    req,
	})
}

func GetTodo(c *gin.Context) {

	var todos []models.Todo
	err := database.DB.Find(&todos).Error
	if HandleDBError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo一覧の取得成功",
		"data":    todos,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var req models.Todo
	err := c.ShouldBindJSON(&req)
	if HandleDBError(c, err) {
		return
	}

	// ここで実際の更新処理を行う（例: データベースを更新）
	if err := database.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(models.Todo{TaskName: req.TaskName, IsCompleted: req.IsCompleted}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo更新成功",
		"id":      req,
	})

}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	// ここで実際の削除処理を行う（例: データベースから削除）
	if err := database.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo削除成功",
		"id":      id,
	})
}

func CreateMemo(c *gin.Context) {
	var req models.Memo

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ここで実際の保存処理を行う（例: データベースに保存）
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "追加しました",
		"data":    req,
	})
}

func GetMemo(c *gin.Context) {
	var memos []models.Memo
	if err := database.DB.Find(&memos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Memo一覧の取得成功",
		"data":    "memos",
	})
}

func UpdateMemo(c *gin.Context) {
	id := c.Param("id")

	var req models.Memo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ここで実際の更新処理を行う（例: データベースを更新）
	if err := database.DB.Model(&models.Memo{}).Where("id = ?", id).Updates(models.Memo{Title: req.Title, Content: req.Content}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Memoの更新成功",
		"id":      req,
	})
}

func DeleteMemo(c *gin.Context) {
	id := c.Param("id")
	// ここで実際の削除処理を行う（例: データベースから削除）
	if err := database.DB.Delete(&models.Memo{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Memoの削除成功",
		"id":      id,
	})
}
