package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ldmtam/serverless-todo/src/models"
)

// MySQL ...
type MySQL struct {
	db *gorm.DB
}

// NewMySQL ...
func NewMySQL() (*MySQL, error) {
	db, err := gorm.Open("mysql", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Todo{})

	return &MySQL{db}, nil
}

// Close ...
func (mysql *MySQL) Close() {
	mysql.db.Close()
}
