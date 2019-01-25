package models

// Todo ...
type Todo struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	IsFinished bool   `json:"is_finished"`
}
