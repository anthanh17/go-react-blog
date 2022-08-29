package models

type PostCategory struct {
	PostId   	uint64	`gorm:"primaryKey" column:"post_id"`
	CategoryId	uint8	`gorm:"primaryKey" column:"category_id"`
}

type Tabler interface {
	TableName() string
}

func (PostCategory) TableName() string {
	return "post_categories"
}