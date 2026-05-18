package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/database"
	"github.com/nazala-mafa/go-crud/handlers"
	"github.com/nazala-mafa/go-crud/middlewares"
	"github.com/nazala-mafa/go-crud/models"
)

func main() {
	config.Load()

	db := database.Connect()

	err := db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.RevokedToken{},
	)

	if err != nil {
		log.Fatal("Migration error: ", err)
	}

	r := gin.Default()

	authHandler := handlers.NewAuthHandler(db)
	r.POST("/api/login", authHandler.Login)

	r.Use(middlewares.AuthMiddleware(db))
	r.POST("/api/logout", authHandler.Logout)
	r.GET("/api/me", authHandler.Me)

	homeHandler := handlers.NewHomeHandler()
	fileHandler := handlers.NewFileHandler()
	r.GET("/api", homeHandler.Index)
	r.POST("/api/upload", fileHandler.Upload)

	productHandler := handlers.NewProductHandler(db)
	products := r.Group("/api/product")
	{
		products.GET("/", productHandler.Index)
		products.POST("/", productHandler.Create)
		products.GET("/:id", productHandler.Show)
		products.PUT("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Destroy)
	}

	userHandler := handlers.NewUserHandler(db)
	users := r.Group("/api/user")
	{
		users.GET("/", userHandler.Index)
		users.POST("/", userHandler.Create)
		users.GET("/:id", userHandler.Show)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Destroy)
	}

	r.Static("/files", "./uploads")

	r.Run(":8080")
}
