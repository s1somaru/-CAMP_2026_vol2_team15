package handlers

import (
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
	// req,err := db.QueryRow("insert into todos (TaskName, IsCompleted) values ($1, $2) returning id ,TaskName,IsCompleted", req.TaskName, req.IsCompleted).Scan(&req.Id, &req.TaskName, &req.IsCompleted)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
		"data":    req,
	})
}

func GetTodo(c *gin.Context) {

	// ここで実際のデータ取得処理を行う（例: データベースから取得）
	//rows ,err := db.Query("select id, TaskName, IsCompleted from todos ")
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// }
	//
	//
	// var todos []models.Todo
	// for rows.Next() {
	// 	var todo models.Todo
	// 	if err := rows.Scan(&todo.Id, &todo.TaskName, &todo.IsCompleted); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	}
	// 	todos = append(todos, todo)
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo一覧の取得成功",
		"data":    "todos",
	})
}

func UpdateTodo(c *gin.Context) {
	//id := c.Param("id")

	var req models.Todo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ここで実際の更新処理を行う（例: データベースを更新）
	// req ,err := db.Exec("update todos set TaskName = $1, IsCompleted = $2 where id = $3", req.TaskName, req.IsCompleted, id returning id, TaskName, IsCompleted).Scan(&req.Id, &req.TaskName, &req.IsCompleted)
	// if err != nil{
	// 	c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo更新成功",
		"id":      req,
	})

}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	// ここで実際の削除処理を行う（例: データベースから削除）
	// _ ,err := db.Exec("delete from todos where id = $1", id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
	// }

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
	//req,err := db.Exec("insert into memos (Title, Content) values ($1,$2) returning Title,Content" ,req.Title,req.Content).Scan(&req.Title,&req.Content)

	c.JSON(http.StatusCreated, gin.H{
		"message": "追加しました",
		"data":    req,
	})
}

func GetMemo(c *gin.Context) {
	// ここで実際のデータ取得処理を行う（例: データベースから取得）
	// row ,err := db.Query(select Title, Content from memos)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// }
	// var memos []models.Memo
	// for row.Next() {
	// 	var memo models.Memo
	// 	if err := row.Scan(&memo.Title, &memo.Content); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	}
	// 	memos = append(memos, memo)
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "Memo一覧の取得成功",
		"data":    "memos",
	})
}

func UpdateMemo(c *gin.Context) {
	// id := c.Param("id")

	var req models.Memo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ここで実際の更新処理を行う（例: データベースを更新）
	// req,err db.Exec("update memos set Title = $1, Content = $2 where id = $3", req.Title, req.Content, id returning id, Title, Content).Scan(&req.Id, &req.Title, &req.Content)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "Memoの更新成功",
		"id":      req,
	})
}

func DeleteMemo(c *gin.Context) {
	id := c.Param("id")

	// ここで実際の削除処理を行う（例: データベースから削除）
	// _,err := db.Exec("delete from memos where id = $1", id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "Memoの削除成功",
		"id":      id,
	})
}
