// di/wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	config "github.com/joejosephvarghese/message/server/pkg" // <- Renamed to "myhttp"
	api "github.com/joejosephvarghese/message/server/pkg/api"
	"github.com/joejosephvarghese/message/server/pkg/api/middleware"
	"github.com/joejosephvarghese/message/server/pkg/db"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {
	wire.Build(
		db.ConnectDatabase, // Ensure database connection is included
		middleware.NewMiddleware,
		api.NewServerHTTP,
	)
	return &api.Server{}, nil
}
