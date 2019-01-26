package models

import "time"

// Todo ...
type Todo struct {
	ID         uint   `gorm:"AUTO_INCREMENT"`
	Content    string `json:"content"`
	IsFinished bool   `json:"is_finished"`
	CreatedAt  time.Time
}
