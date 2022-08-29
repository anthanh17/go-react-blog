package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	Id              	uint64	`gorm:"primaryKey; AUTO_INCREMENT"`
	ParentCommentId		uint64	`gorm:"not null"`
	AccountCommentId	uint64	`gorm:"not null"`	
	PostId          	uint64	`gorm:"not null; unique"`
	Content         	string	`gorm:"type:text"`
	Mask            	bool
	gorm.Model
}