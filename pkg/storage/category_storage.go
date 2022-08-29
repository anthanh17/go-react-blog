package storage

import (
	"github.com/anthanh17/go-react-blog/pkg/db/models"
)

func (storage *MysqlStorage) CreateCategory(category *models.Category) {
	storage.Db.Create(&category)
}

func (storage *MysqlStorage) FindCategory(condition map[string]interface{}) (*models.Category, error) {
	var category models.Category

	err := storage.Db.Where(condition).Find(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (storage *MysqlStorage) UpdateCategory(condition map[string]interface{}, category models.Category) bool {
	if storage.Db.Where(condition).Updates(&category).RowsAffected == 0 {
		return false
	}
	return true
}

func (storage *MysqlStorage) DeleteCategory(category *models.Category) bool {
	if storage.Db.Model(category).Update("mask", true).RowsAffected == 0 {
		return false
	}
	return true
}