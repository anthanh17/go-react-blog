package storage

import (
	"github.com/anthanh17/go-react-blog/pkg/db/models"
)

func (storage *MysqlStorage) CreateComment(comment *models.Comment) {
	storage.Db.Create(&comment)
}

func (storage *MysqlStorage) FindComment(condition map[string]interface{}) (*models.Comment, error) {
	var comment models.Comment

	err := storage.Db.Where(condition).Find(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (storage *MysqlStorage) UpdateComment(condition map[string]interface{}, comment models.Comment) bool {
	if storage.Db.Where(condition).Updates(&comment).RowsAffected == 0 {
		return false
	}
	return true
}

func (storage *MysqlStorage) DeleteComment(comment models.Comment) bool {
	if storage.Db.Model(comment).Update("mask", true).RowsAffected == 0 {
		return false
	}
	return true
}