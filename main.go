package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/database"
	"github.com/nazala-mafa/go-crud/models"
	"github.com/nazala-mafa/go-crud/routes"
)

func main() {
	cfg := config.Load()

	db := database.Connect(cfg)

	err := db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.RevokedToken{},
	)

	if err != nil {
		log.Fatal("Migration error: ", err)
	}

	r := gin.Default()

	routes.Routes(r, db, cfg)

	r.Run(":8080")
}
