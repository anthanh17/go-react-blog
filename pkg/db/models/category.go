package models

type Category struct {
	Id  	uint8	`gorm:"primaryKey; AUTO_INCREMENT"`
	Name	string	`gorm:"type:varchar(255); not null; unique"`
	Mask	bool
	Slug	string	`gorm:"type:varchar(255); not null"`
	Post	[]Post	`gorm:"many2many:post_categories"`
	// PostCategory	[]PostCategory	`gorm:"many2one:CategoryId"`
}