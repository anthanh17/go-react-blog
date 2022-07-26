package config

import (
	"fmt"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	config "blog/pkg/config"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	config, err := config.LoadDatabaseConfig("configs");
	if err != nil {
		log.Fatal("cann't load config", err);
	}

	Db = connectDB(config)
	return Db
}

func connectDB(config config.Config) (*gorm.DB) {
	db, err := gorm.Open(mysql.Open(config.Database.DbSource), &gorm.Config{})
	
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	fmt.Println("\nCome here")
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(5)
	// db.SetMaxIdleConns(5)

	// db.AutoMigrate(&models.User{})

	return db
} 