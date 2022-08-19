package storage

import (
	"github.com/anthanh17/go-react-blog/pkg/db/models"
)

func (storage *MysqlStorage) CreatePostCategory(post *models.Post, category *models.Category) {
	var PostCategory models.PostCategory
	PostCategory.PostId = post.Id
	PostCategory.CategoryId = category.Id

	storage.Db.Create(&PostCategory)
}