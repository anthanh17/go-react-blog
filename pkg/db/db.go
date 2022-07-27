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
	dns := fmt.Sprintf(
		"%s:%s@(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		MySQLConfig.User,
		MySQLConfig.Password,
		MySQLConfig.Host,
		MySQLConfig.DbName)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
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
