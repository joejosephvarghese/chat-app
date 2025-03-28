// di/wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	config "github.com/joejosephvarghese/message/server/pkg" // <- Renamed to "myhttp"
	api "github.com/joejosephvarghese/message/server/pkg/api"
	"github.com/joejosephvarghese/message/server/pkg/api/middleware"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {
	wire.Build(
		middleware.NewMiddleware,
		api.NewServerHTTP,
	)
	return &api.Server{}, nil
}
