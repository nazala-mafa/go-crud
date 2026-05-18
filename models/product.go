package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string  `gorm:"type:varchar(256)" json:"name" faker:"sentence"`
	Description string  `gorm:"type:longtext" json:"description" faker:"sentence"`
	Price       float64 `gorm:"type:decimal(20,2)" json:"price" faker:"boundary_start=1000, boundary_end=5000"`
	ImageURL    string  `gorm:"type:varchar(256);null" json:"image_url" faker:"url"`
}
