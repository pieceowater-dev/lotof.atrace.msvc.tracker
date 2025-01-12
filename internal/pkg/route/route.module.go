package route

import (
	"app/internal/core/cfg"
	"app/internal/pkg/route/ctrl"
	"app/internal/pkg/route/ent"
	"app/internal/pkg/route/svc"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	Controller *ctrl.RouteController
}

func New() *Module {
	entities := []any{
		&ent.Route{},
		&ent.RouteMilestone{},
	}
	// Create database instance
	database, err := gossiper.NewDB(
		gossiper.PostgresDB,
		cfg.Inst().PostgresDatabaseDSN,
		false,
		entities,
	)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	err = database.MigrateTenants(
		[]string{
			"excepteur_ipsum",
		},
		entities,
	)
	if err != nil {
		log.Fatalf("Failed to migrate tenants: %v", err)
	}

	// Initialize and return the module
	return &Module{
		Controller: ctrl.NewRouteController(
			svc.NewRouteService(database),
		),
	}
}
