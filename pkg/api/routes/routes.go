package routes

import (
	"github.com/gin-gonic/gin"
	interfaces "github.com/joejosephvarghese/message/server/pkg/api/handler/interface"
)

func SetUpRoutes(api *gin.Engine, authHandler interfaces.AuthHandler) {
	auth := api.Group("")
	{ // auth
		auth.POST(SignUpURL, authHandler.UserSignUp)

	}

}
