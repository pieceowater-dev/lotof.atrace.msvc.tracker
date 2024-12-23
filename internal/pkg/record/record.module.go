package record

import (
	"app/internal/core/cfg"
	"app/internal/pkg/record/ctrl"
	"app/internal/pkg/record/svc"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	Controller *ctrl.RecordController
}

func New() *Module {
	database, err := gossiper.NewDB(
		gossiper.PostgresDB,
		cfg.Inst().PostgresDatabaseDSN,
		false,
		[]any{},
	)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	return &Module{
		Controller: ctrl.NewRecordController(
			svc.NewRecordService(database),
		),
	}
}
