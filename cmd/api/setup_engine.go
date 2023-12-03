package api

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/juanignaciorc/microbloggin-pltf/src/api_adapters"
	"github.com/juanignaciorc/microbloggin-pltf/src/repositories"
	"github.com/juanignaciorc/microbloggin-pltf/src/services"
)

const basePath = "/api/v1"

func SetupEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	db := repositories.NewInMemoryDB()
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})
	})
	router.POST(basePath+"/users", userHandler.Create)
	router.GET(basePath+"/users/:id", userHandler.Get)
	router.POST(basePath+"/users/:id/tweet", userHandler.CreateTweet)

	return router
}
