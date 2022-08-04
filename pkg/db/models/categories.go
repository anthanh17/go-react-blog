package models

type Categories struct {
	Id  	uint8	`gorm:"primaryKey; AUTO_INCREMENT"`
	Name	string	`gorm:"type:varchar(255); not null; unique"`
	Slug	string	`gorm:"type:varchar(255); not null"`
	// PostCategory	[]PostCategory	`gorm:"many2one:CategoryId"`
}