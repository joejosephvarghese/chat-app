package handler

import (
	"github.com/gin-gonic/gin"
	interfaces "github.com/joejosephvarghese/message/server/pkg/api/handler/interface"
	"github.com/joejosephvarghese/message/server/pkg/api/handler/request"
	useCaseInterface "github.com/joejosephvarghese/message/server/pkg/usecase/interface"
)

type AuthHandler struct {
	authUseCase useCaseInterface.AuthUseCase
}

func NewAuthHandler(authUseCase useCaseInterface.AuthUseCase) interfaces.AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}
func (c *AuthHandler) UserSignUp(ctx *gin.Context) {

	var body request.UserSignUp

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "User sign-up successful", "data": body})
}
