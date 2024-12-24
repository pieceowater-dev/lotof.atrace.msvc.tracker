package svc

import (
	"app/internal/pkg/record/ent"
	"context"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type RecordService struct {
	db gossiper.Database
}

func NewRecordService(db gossiper.Database) *RecordService {
	return &RecordService{db: db}
}

// GetRecords retrieves paginated records from the database.
func (s RecordService) GetRecords(ctx context.Context, filter gossiper.Filter[string]) (gossiper.PaginatedResult[ent.Record], error) {
	var records []ent.Record
	var count int64

	query := s.db.GetDB().Model(&ent.Record{})

	// Apply search filters
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("user_id LIKE ?", search)
	}

	// Count total records
	if err := query.Count(&count).Error; err != nil {
		return gossiper.PaginatedResult[ent.Record]{}, fmt.Errorf("failed to count records: %w", err)
	}

	// Apply pagination
	query = query.Offset((filter.Pagination.Page - 1) * filter.Pagination.Length).Limit(filter.Pagination.Length)

	// Apply sorting dynamically
	if field := filter.Sort.Field; field != "" && gossiper.IsFieldValid(&ent.Record{}, field) {
		query = query.Order(fmt.Sprintf("%s %s", gossiper.ToSnakeCase(field), filter.Sort.Direction))
	}

	// Fetch data
	if err := query.Find(&records).Error; err != nil {
		return gossiper.PaginatedResult[ent.Record]{}, fmt.Errorf("failed to fetch records: %w", err)
	}

	return gossiper.NewPaginatedResult(records, int(count)), nil
}

// GetRecordsByPostId retrieves paginated records by post id from the database.
func (s RecordService) GetRecordsByPostId(ctx context.Context, postId int, filter gossiper.Filter[string]) (gossiper.PaginatedResult[ent.Record], error) {
	var records []ent.Record
	var count int64

	query := s.db.GetDB().Model(&ent.Record{}).Where("post_id = ?", postId)

	// Count total records
	if err := query.Count(&count).Error; err != nil {
		return gossiper.PaginatedResult[ent.Record]{}, fmt.Errorf("failed to count records: %w", err)
	}

	// Apply pagination
	query = query.Offset((filter.Pagination.Page - 1) * filter.Pagination.Length).Limit(filter.Pagination.Length)

	// Apply sorting dynamically
	if field := filter.Sort.Field; field != "" && gossiper.IsFieldValid(&ent.Record{}, field) {
		query = query.Order(fmt.Sprintf("%s %s", gossiper.ToSnakeCase(field), filter.Sort.Direction))
	}

	// Fetch data
	if err := query.Find(&records).Error; err != nil {
		return gossiper.PaginatedResult[ent.Record]{}, fmt.Errorf("failed to fetch records: %w", err)
	}

	return gossiper.NewPaginatedResult(records, int(count)), nil
}

// CreateRecord adds a new record to the database.
func (s RecordService) CreateRecord(ctx context.Context, record *ent.Record) (*ent.Record, error) {
	if err := s.db.GetDB().Create(record).Error; err != nil {
		return nil, fmt.Errorf("failed to create record: %w", err)
	}
	return record, nil
}
