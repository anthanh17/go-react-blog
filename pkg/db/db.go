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
	// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;chartset=utf8;encrypt=disable",
	// 	MySQLConfig.Host, MySQLConfig.User, MySQLConfig.Password, MySQLConfig.Port, MySQLConfig.DbName)

	db, err := gorm.Open(mysql.Open(MySQLConfig.DbSource), &gorm.Config{})
	if err == nil {
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		err = db.Ping()
	} else if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	return db, err
}
