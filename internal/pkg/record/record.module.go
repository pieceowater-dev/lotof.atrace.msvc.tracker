package record

import (
	"app/internal/core/cfg"
	"app/internal/pkg/record/ctrl"
	"app/internal/pkg/record/ent"
	"app/internal/pkg/record/svc"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	Controller *ctrl.RecordController
}

func New() *Module {
	entities := []any{
		&ent.Record{},
	}
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
			"someSchema",
		},
		entities,
	)
	if err != nil {
		log.Fatalf("Failed to migrate tenants: %v", err)
	}

	return &Module{
		Controller: ctrl.NewRecordController(
			svc.NewRecordService(database),
		),
	}
}
