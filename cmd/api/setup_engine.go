package api

import (
	"github.com/gin-gonic/gin"
	handlers2 "github.com/juanignaciorc/microbloggin-pltf/internal/adapters/handlers"
	"github.com/juanignaciorc/microbloggin-pltf/internal/adapters/repositories/in_memory_db"
	"github.com/juanignaciorc/microbloggin-pltf/internal/services"
)

const basePath = "/handlers/v1"

func SetupEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	db := in_memory_db.NewInMemoryDB()
	userService := services.NewUserService(db)
	userHandler := handlers2.NewUserHandler(userService)

	tweetService := services.NewTweetsService(db)
	tweetHandler := handlers2.NewTweetHandler(tweetService)

	router.GET("/ping", handlers2.PingHandler)
	router.POST(basePath+"/users", userHandler.Create)
	router.GET(basePath+"/users/:id", userHandler.Get)
	router.POST(basePath+"/users/:id/tweet", tweetHandler.CreateTweet)
	router.POST(basePath+"/users/:id/follow/:following_user_id", userHandler.FollowUser)
	router.GET(basePath+"/users/:id/timeline", userHandler.GetUserTimeline)

	return router
}
