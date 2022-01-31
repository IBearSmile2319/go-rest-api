package database

import (
	"github.com/IBearSmile2319/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

// Migration - migrate our database and create our comment tables
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil

}
