package repository

import (
	"time"
	"gorm.io/gorm"

	"github.com/anthanh17/go-react-blog/pkg/db"
	"github.com/anthanh17/go-react-blog/pkg/db/models"
)
// func GetMySqlDatabase() (*gorm.DB, error) {
func SetupDb() (*gorm.DB, error) {
	db, err := db.GetMySqlDatabase()

	// fmt.Println(&db)

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
				
	db.AutoMigrate(&models.Account{}, &models.Post{}, &models.Category{}, &models.Comment{})

	return db, err
}
