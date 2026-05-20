package seeders

import (
	"github.com/go-faker/faker/v4"
	"github.com/nazala-mafa/go-crud/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	users := make([]models.User, 0, 22)

	for i := 0; i < 22; i++ {
		var p models.User

		if err := faker.FakeData(&p); err != err {
			panic(err)
		}

		hash, err := bcrypt.GenerateFromPassword(
			[]byte("password"),
			bcrypt.DefaultCost,
		)

		if err != nil {
			panic(err)
		}

		p.ID = 0
		p.Password = string(hash)
		p.DeletedAt = gorm.DeletedAt{}
		p.AvatarURL = "https://placehold.co/640x480/png?text=et-dolores-nulla-eum-sit-atque&w=640&q=75"

		if i == 0 {
			p.Name = "test"
			p.Email = "test@example.com"
		}

		users = append(users, p)
	}

	if err := db.Create(&users).Error; err != nil {
		panic(err)
	}
}
