package api

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/juanignaciorc/microbloggin-pltf/internal/api_adapters"
	"github.com/juanignaciorc/microbloggin-pltf/internal/repositories"
	"github.com/juanignaciorc/microbloggin-pltf/internal/services"
)

const basePath = "/api/v1"

func SetupEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	db := repositories.NewInMemoryDB()
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	tweetService := services.NewTweetsService(db)
	tweetHandler := handlers.NewTweetHandler(tweetService)

	router.GET("/ping", handlers.PingHandler)
	router.POST(basePath+"/users", userHandler.Create)
	router.GET(basePath+"/users/:id", userHandler.Get)
	router.POST(basePath+"/users/:id/tweet", tweetHandler.CreateTweet)
	router.POST(basePath+"/users/:id/follow/:following_user_id", userHandler.FollowUser)
	router.GET(basePath+"/users/:id/timeline", userHandler.GetUserTimeline)

	return router
}
