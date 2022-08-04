package models

import (
	"gorm.io/gorm"
)

type PostComment struct {
	Id      	uint64		`gorm:"primaryKey; AUTO_INCREMENT"`
	PostId		string		`gorm:"type:varchar(255); not null; unique"`
	Content 	string		`gorm:"type:text"`
	Mask		bool
	gorm.Model
}