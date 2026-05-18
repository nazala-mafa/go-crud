package main

import (
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/database"
	"github.com/nazala-mafa/go-crud/seeders"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	seeders.UserSeeder(db)
	seeders.ProductSeeder(db)
}
