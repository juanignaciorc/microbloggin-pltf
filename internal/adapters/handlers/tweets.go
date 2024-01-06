package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juanignaciorc/microbloggin-pltf/internal/services"
	"net/http"
)

type TweetHandler struct {
	service services.TweetsService
}

func NewTweetHandler(service services.TweetsService) *TweetHandler {
	return &TweetHandler{
		service: service,
	}
}

func (h *TweetHandler) CreateTweet(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uuid.Validate(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var body CreateTweetBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet, err := h.service.CreateTweet(id, body.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tweet created successfully", "tweet": tweet})
}
