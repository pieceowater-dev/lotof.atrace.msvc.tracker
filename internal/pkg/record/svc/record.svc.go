package svc

import gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"

type RecordService struct {
	db gossiper.Database
}

func NewRecordService(db gossiper.Database) *RecordService {
	return &RecordService{db: db}
}
