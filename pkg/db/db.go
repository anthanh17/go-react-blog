package db

import (
	"fmt"

	"github.com/anthanh17/go-react-blog/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMySqlDatabase(config config.Settings) (*gorm.DB, error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySqlCfg.User,
		config.MySqlCfg.Password,
		config.MySqlCfg.Host,
		config.MySqlCfg.Port,
		config.MySqlCfg.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		//fmt.Println("Error connecting to database : error=%v", err)
		return nil, nil
	}

	return db, err
}
