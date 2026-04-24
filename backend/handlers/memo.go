package handlers

import (
	"go-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {

	var req models.Todo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ここで実際の保存処理を行う（例: データベースに保存）

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
		"data":    req,
	})
}

func GetTodo(c *gin.Context) {

	// ここで実際のデータ取得処理を行う（例: データベースから取得）

	//仮データ
	todos := []models.Todo{
		{
			TaskName:    "株式会社A のエントリーシートを提出する",
			IsCompleted: false,
		},
		{
			TaskName:    "株式会社B の面接準備",
			IsCompleted: true,
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo一覧の取得成功",
		"data":    todos,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var req models.Todo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ここで実際の更新処理を行う（例: データベースを更新）

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo更新成功",
		"id":      id,
	})

}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	// ここで実際の削除処理を行う（例: データベースから削除）

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

	c.JSON(http.StatusCreated, gin.H{
		"message": "追加しました",
		"data":    req,
	})
}

func GetMemo(c *gin.Context) {
	// ここで実際のデータ取得処理を行う（例: データベースから取得）

	//仮データ
	memos := []models.Memo{
		{
			Title:   "面接の準備",
			Content: "面接の前に企業研究をしっかり行うこと",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Memo一覧の取得成功",
		"data":    memos,
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Memoの更新成功",
		"id":      id,
	})
}

func DeleteMemo(c *gin.Context) {
	id := c.Param("id")

	// ここで実際の削除処理を行う（例: データベースから削除）

	c.JSON(http.StatusOK, gin.H{
		"message": "Memoの削除成功",
		"id":      id,
	})
}
