package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/record/svc"
	"context"
)

type RecordController struct {
	recordService *svc.RecordService
	pb.UnimplementedRecordServiceServer
}

func NewRecordController(service *svc.RecordService) *RecordController {
	return &RecordController{recordService: service}
}

func (r RecordController) GetRecords(ctx context.Context, request *pb.GetRecordsRequest) (*pb.GetRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RecordController) CreateRecord(ctx context.Context, request *pb.CreateRecordRequest) (*pb.Record, error) {
	//TODO implement me
	panic("implement me")
}
