package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Cors() gin.HandlerFunc
}

type middleware struct{}

func NewMiddleware() Middleware { // ✅ Return the interface, not *Middleware
	return &middleware{}
}
