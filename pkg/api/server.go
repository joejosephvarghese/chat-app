package api

import (
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/joejosephvarghese/message/server/pkg"
	interfaces "github.com/joejosephvarghese/message/server/pkg/api/handler/interface"
	"github.com/joejosephvarghese/message/server/pkg/api/middleware"
	"github.com/joejosephvarghese/message/server/pkg/api/routes"

	"gorm.io/gorm"
)

type Server struct {
	engine *gin.Engine
	port   string
	db     *gorm.DB // ✅ Add Database field
}

func NewServerHTTP(cfg config.Config, authHandler interfaces.AuthHandler, db *gorm.DB, middleware middleware.Middleware) *Server {
	engine := gin.New()

	engine.Use(middleware.Cors()) // Now, mw.Cors() works correctly
	engine.Use(gin.Logger())
	routes.SetUpRoutes(engine, authHandler)
	return &Server{engine: engine, port: cfg.Port}
}
func (c *Server) Start() error {

	// Ensure a valid port
	port := c.port
	if port == "" || port == "5432" { // Prevent using the DB port
		port = "8080"
		log.Println("Defaulting to port 8080")
	}
	log.Println("Starting server on port:", port)
	return c.engine.Run(":" + port)
}
