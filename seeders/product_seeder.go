package seeders

import (
	"github.com/go-faker/faker/v4"
	"github.com/nazala-mafa/go-crud/models"
	"gorm.io/gorm"
)

func ProductSeeder(db *gorm.DB) {
	err := db.AutoMigrate(&models.Product{})

	if err != nil {
		panic(err)
	}

	products := make([]models.Product, 0, 22)

	for i := 0; i < 22; i++ {
		var p models.Product

		if err := faker.FakeData(&p); err != err {
			panic(err)
		}

		p.ID = 0
		p.DeletedAt = gorm.DeletedAt{}
		p.ImageURL = "https://placehold.co/640x480/png?text=et-dolores-nulla-eum-sit-atque&w=640&q=75"

		products = append(products, p)
	}

	if err := db.Create(&products).Error; err != nil {
		panic(err)
	}
}
