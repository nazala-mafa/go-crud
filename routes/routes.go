package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nazala-mafa/go-crud/config"
	"github.com/nazala-mafa/go-crud/handlers"
	"github.com/nazala-mafa/go-crud/middlewares"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	authHandler := handlers.NewAuthHandler(db, cfg)
	homeHandler := handlers.NewHomeHandler()
	fileHandler := handlers.NewFileHandler(cfg)
	productHandler := handlers.NewProductHandler(db, cfg)
	userHandler := handlers.NewUserHandler(db, cfg)

	r.GET("/", homeHandler.Index)
	api := r.Group("/api")
	{
		// Public routes
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.GET("/", homeHandler.Index)

		// Protected routes
		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware(db, cfg))
		{
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/me", authHandler.Me)
			auth.POST("/upload", fileHandler.Upload)

			products := auth.Group("/product")
			{
				products.GET("/", productHandler.Index)
				products.POST("/", productHandler.Create)
				products.GET("/:id", productHandler.Show)
				products.PUT("/:id", productHandler.Update)
				products.DELETE("/:id", productHandler.Destroy)
			}

			users := auth.Group("/user")
			{
				users.GET("/", userHandler.Index)
				users.POST("/", userHandler.Create)
				users.GET("/:id", userHandler.Show)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", userHandler.Destroy)
			}
		}
	}

	r.Static("/files", "./uploads")
}
