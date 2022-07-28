package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/anthanh17/go-react-blog/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db          *gorm.DB
	MySQLConfig *config.MySQLConfig
)

func SetupConfig(cfg config.Config) {
	if cfg.MySQLConfig == nil {
		log.Fatal(errors.New(fmt.Sprintf("Can not load MySQL config!")))
		return
	}
	MySQLConfig = cfg.MySQLConfig
}

func GetMySqlDatabase() (*gorm.DB, error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		MySQLConfig.User,
		MySQLConfig.Password,
		MySQLConfig.Host,
		MySQLConfig.Port,
		MySQLConfig.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err == nil {
	// 	db.SetMaxOpenConns(10)
	// 	db.SetMaxIdleConns(10)
	// 	err = db.Ping()
	// } else
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil, nil
	}
	return db, err
}
