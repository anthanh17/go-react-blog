package storage

import (
	"gorm.io/gorm"
)

type MysqlStorage struct {
	Db *gorm.DB
}

/*Constructor*/
func NewMySqlStorage(db *gorm.DB) *MysqlStorage {
	return &MysqlStorage{Db: db}
}

