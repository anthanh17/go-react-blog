package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	Id      	uint64			`gorm:"primaryKey; AUTO_INCREMENT"`
	Username	string			`gorm:"type:varchar(255); not null; unique"`
	Password 	string			`gorm:"type:varchar(255); not null"`
	FullName	string			`gorm:"type:varchar(255); not null"`
	Email		string			`gorm:"type:varchar(255)"`
	Telephone	string			`gorm:"type:varchar(255)"`
	LastLogin	time.Time		`json:"created_at"`
	Permission	bool
	Profile		string			`gorm:"type:text"`
	Subscribe	bool			`gorm:"not null"`
	gorm.Model
	Post		[]Post			`gorm:"foreignKey:AuthorId"`
	PostComment	[]PostComment	`gorm:"foreignKey:AccountCommentId"`
}