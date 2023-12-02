package api

import "github.com/gin-gonic/gin"

func SetupEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})
	})
	
	return router
}