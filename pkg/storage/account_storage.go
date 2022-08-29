package storage

import (
	"github.com/anthanh17/go-react-blog/pkg/db/models"
)

func (storage *MysqlStorage) CreateAccount(account *models.Account) {
	storage.Db.Create(&account)
}

func (storage *MysqlStorage) FindAccount(condition map[string]interface{}) (*models.Account, error) {
	var account models.Account

	err := storage.Db.Where(condition).Find(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (storage *MysqlStorage) UpdateAccount(condition map[string]interface{}, account models.Account) bool {
	if storage.Db.Where(condition).Updates(&account).RowsAffected == 0 {
		return false
	}
	return true
}

func (storage *MysqlStorage) DeleteAccount(account *models.Account) bool {
	// if storage.Db.Where(condition).Updates(&account).RowsAffected == 0 {
	// 	return false
	// }

	// if storage.Db.Model(&account).Where(condition).Updates(models.Acco)).RowsAffected == 0 {
	// 	return false
	// }

	if storage.Db.Model(account).Update("mask", true).RowsAffected == 0 {
		return false
	}
	return true
}