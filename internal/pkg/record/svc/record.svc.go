package svc

import (
	"app/internal/pkg/record/ent"
	"context"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type RecordService struct {
	db gossiper.Database
}

func NewRecordService(db gossiper.Database) *RecordService {
	return &RecordService{db: db}
}

func (s RecordService) GetRecords(ctx context.Context, filter gossiper.Filter[string]) (gossiper.PaginatedResult[ent.Record], error) {
	//TODO implement me
	panic("implement me")
}

func (s RecordService) CreateRecord(ctx context.Context, request *ent.Record) (*ent.Record, error) {
	//TODO implement me
	panic("implement me")
}
