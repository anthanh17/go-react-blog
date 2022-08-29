package storage

import (
	"github.com/anthanh17/go-react-blog/pkg/db/models"
)

func (storage *MysqlStorage) CreatePost(post *models.Post) {
	storage.Db.Create(&post)
}

func (storage *MysqlStorage) CreatePosts(posts []map[string]interface{}) {
	storage.Db.Create(&posts)
}

func (storage *MysqlStorage) FindPost(condition map[string]interface{}) (*models.Post, error) {
	var post models.Post

	err := storage.Db.Where(condition).Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (storage *MysqlStorage) UpdatePost(condition map[string]interface{}, post models.Post) bool {
	if storage.Db.Where(condition).Updates(&post).RowsAffected == 0 {
		return false
	}
	return true
}

func (storage *MysqlStorage) DeletePost(account *models.Post) bool {
	if storage.Db.Model(account).Update("mask", true).RowsAffected == 0 {
		return false
	}
	return true
}