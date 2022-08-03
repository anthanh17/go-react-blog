package models

import (
	"gorm.io/gorm"
)

type Post struct {
	Id      	uint64       	`gorm:"primaryKey; AUTO_INCREMENT"`
	Title		string       	`gorm:"type:varchar(255); not null; unique"`
	Content 	string       	`gorm:"type:text"`
	AuthorId	uint8
	Media		string       	`gorm:"type:varchar(255); not null"`
	Views		uint32
	Likes		uint32
	Mask		bool
	Slug		string       	`gorm:"type:varchar(255); not null"`
	gorm.Model
	Categories	[]Categories	`gorm:"many2many:post_categories;"`
	PostComment	[]PostComment	`gorm:"foreignKey:PostId"`
}