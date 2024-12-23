package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/record/svc"
)

type RecordController struct {
	recordService *svc.RecordService
	pb.UnimplementedRecordServiceServer
}

func NewRecordController(service *svc.RecordService) *RecordController {
	return &RecordController{recordService: service}
}
