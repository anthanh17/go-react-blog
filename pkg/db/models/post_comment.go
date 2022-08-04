package models

import (
	"gorm.io/gorm"
)

type PostComment struct {
	Id      			uint64	`gorm:"primaryKey; AUTO_INCREMENT"`
	ParentCommentId		uint64	`gorm:"not null"`
	AccountCommentId	uint64	`gorm:"not null"`	
	PostId				string	`gorm:"type:varchar(255); not null; unique"`
	Content 			string	`gorm:"type:text"`
	Mask				bool
	gorm.Model
}