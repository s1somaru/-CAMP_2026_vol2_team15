package models

type Todo struct {
	Id          int    `json:"id"`
	TaskName    string `json:"task_name"`
	IsCompleted bool   `json:"is_completed"`
}

type Memo struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
