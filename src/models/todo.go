package models

import "time"

// Todo ...
type Todo struct {
	ID         uint      `json:"id" gorm:"AUTO_INCREMENT"`
	Content    string    `json:"content"`
	IsFinished bool      `json:"is_finished"`
	CreatedAt  time.Time `json:"created_at"`
}
