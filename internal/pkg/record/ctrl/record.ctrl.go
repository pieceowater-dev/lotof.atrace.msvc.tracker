package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/record/ent"
	"app/internal/pkg/record/svc"
	"context"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type RecordController struct {
	recordService *svc.RecordService
	pb.UnimplementedRecordServiceServer
}

func NewRecordController(service *svc.RecordService) *RecordController {
	return &RecordController{recordService: service}
}

// GetRecords retrieves paginated records.
func (r RecordController) GetRecords(ctx context.Context, request *pb.GetRecordsRequest) (*pb.GetRecordsResponse, error) {
	filter := gossiper.NewFilter[string](
		request.GetSearch(),
		gossiper.NewSort[string](
			request.GetSort().GetField(),
			gossiper.SortDirection(request.GetSort().GetDirection()),
		),
		gossiper.NewPagination(
			int(request.GetPagination().GetPage()),
			int(request.GetPagination().GetLength()),
		),
	)

	paginatedResult, err := r.recordService.GetRecords(ctx, filter)
	if err != nil {
		return nil, err
	}

	var records []*pb.Record
	for _, record := range paginatedResult.Rows {
		records = append(records, &pb.Record{
			Id:        uint64(record.ID),
			PostId:    uint64(record.PostID),
			UserId:    record.UserID,
			Timestamp: uint64(record.Timestamp.Unix()),
			Method:    pb.RecordMethod(record.Method),
		})
	}

	return &pb.GetRecordsResponse{
		Records: records,
		PaginationInfo: &pb.PaginationInfo{
			Count: int32(paginatedResult.Info.Count),
		},
	}, nil
}

func (r RecordController) GetRecordsByPostId(ctx context.Context, request *pb.GetRecordsByPostIdRequest) (*pb.GetRecordsResponse, error) {
	filter := gossiper.NewFilter[string](
		"",
		gossiper.NewSort[string](
			request.GetSort().GetField(),
			gossiper.SortDirection(request.GetSort().GetDirection()),
		),
		gossiper.NewPagination(
			int(request.GetPagination().GetPage()),
			int(request.GetPagination().GetLength()),
		),
	)

	paginatedResult, err := r.recordService.GetRecordsByPostId(ctx, int(request.GetPostId()), filter)
	if err != nil {
		return nil, err
	}

	var records []*pb.Record
	for _, record := range paginatedResult.Rows {
		records = append(records, &pb.Record{
			Id:        uint64(record.ID),
			PostId:    uint64(record.PostID),
			UserId:    record.UserID,
			Timestamp: uint64(record.Timestamp.Unix()),
			Method:    pb.RecordMethod(record.Method),
		})
	}

	return &pb.GetRecordsResponse{
		Records: records,
		PaginationInfo: &pb.PaginationInfo{
			Count: int32(paginatedResult.Info.Count),
		},
	}, nil
}

// CreateRecord creates a new record.
func (r RecordController) CreateRecord(ctx context.Context, request *pb.CreateRecordRequest) (*pb.Record, error) {
	record := &ent.Record{
		PostID: uint(request.PostId),
		UserID: request.UserId, // todo: get from token, etc
		Method: ent.RecordMethod(request.Method),
	}

	createdRecord, err := r.recordService.CreateRecord(ctx, record)
	if err != nil {
		return nil, fmt.Errorf("failed to create record: %w", err)
	}

	return &pb.Record{
		Id:        uint64(createdRecord.ID),
		PostId:    uint64(createdRecord.PostID),
		UserId:    createdRecord.UserID,
		Timestamp: uint64(createdRecord.Timestamp.Unix()), // todo: make data i guess (?)
		Method:    pb.RecordMethod(createdRecord.Method),
	}, nil
}
