package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	TaskName    string `json:"task_name"`
	IsCompleted bool   `json:"is_completed"`
}

type Memo struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
