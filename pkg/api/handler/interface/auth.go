package interfaces

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	//userSide

	UserSignUp(ctx *gin.Context)
}
