package route

import (
	"app/internal/core/cfg"
	"app/internal/pkg/route/ctrl"
	"app/internal/pkg/route/svc"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	Controller *ctrl.RouteController
}

func New() *Module {
	// Create database instance
	database, err := gossiper.NewDB(
		gossiper.PostgresDB,
		cfg.Inst().PostgresDatabaseDSN,
		false,
		[]any{},
	)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	// Initialize and return the module
	return &Module{
		Controller: ctrl.NewRouteController(
			svc.NewRouteService(database),
		),
	}
}
