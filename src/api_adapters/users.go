package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanignaciorc/microbloggin-pltf/src/services"
)

type UserHandler struct {
	service services.UserService
}

type CreateUserBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var body CreateUserBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateUser(body.Name, body.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
