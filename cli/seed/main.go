package main

import (
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/database"
	"github.com/nazala-mafa/go-crud/seeders"
)

func main() {
	config.Load()
	db := database.Connect()

	seeders.UserSeeder(db)
	seeders.ProductSeeder(db)
}
