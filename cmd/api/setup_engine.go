package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanignaciorc/microbloggin-pltf/internal/adapters/handlers"
	"github.com/juanignaciorc/microbloggin-pltf/internal/adapters/repositories/postgre_db"
	"github.com/juanignaciorc/microbloggin-pltf/internal/services"
)

const basePath = "/api/v1"

func SetupEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	ctx := context.Background()
	db, err := postgre_db.NewDB(ctx, "user=postgres password=postgres host=127.0.0.1 port=5432 dbname=Microblogging sslmode=prefer connect_timeout=5")
	if err != nil {
		fmt.Print("Error connecting to the database")
		panic(err)
	}

	//Dependency Injection
	userRepo := postgre_db.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	tweetRepo := postgre_db.NewTweetRepository(db)
	tweetService := services.NewTweetsService(tweetRepo)
	tweetHandler := handlers.NewTweetHandler(tweetService)

	router.GET("/ping", handlers.PingHandler)
	router.POST(basePath+"/users", userHandler.Create)
	router.GET(basePath+"/users/:id", userHandler.Get)
	router.POST(basePath+"/users/:id/tweet", tweetHandler.CreateTweet)
	router.POST(basePath+"/users/:id/follow/:following_user_id", userHandler.FollowUser)
	router.GET(basePath+"/users/:id/timeline", userHandler.GetUserTimeline)

	return router
}
