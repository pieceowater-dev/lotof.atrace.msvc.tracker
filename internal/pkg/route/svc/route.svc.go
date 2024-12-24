package svc

import (
	"app/internal/pkg/route/ent"
	"context"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"gorm.io/gorm"
)

type RouteService struct {
	db gossiper.Database
}

// NewRouteService initializes a new RouteService.
func NewRouteService(db gossiper.Database) *RouteService {
	return &RouteService{db: db}
}

// GetRoutes retrieves paginated routes with optional search and sorting.
func (s RouteService) GetRoutes(ctx context.Context, filter gossiper.Filter[string]) (gossiper.PaginatedResult[ent.Route], error) {
	var routes []ent.Route
	var count int64

	query := s.db.GetDB().Model(&ent.Route{}).Preload("Milestones")

	// Apply search filters
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("title LIKE ?", search)
	}

	// Count total records
	if err := query.Count(&count).Error; err != nil {
		return gossiper.PaginatedResult[ent.Route]{}, fmt.Errorf("failed to count routes: %w", err)
	}

	// Apply pagination
	query = query.Offset((filter.Pagination.Page - 1) * filter.Pagination.Length).Limit(filter.Pagination.Length)

	// Apply sorting dynamically
	if field := filter.Sort.Field; field != "" && gossiper.IsFieldValid(&ent.Route{}, field) {
		query = query.Order(fmt.Sprintf("%s %s", gossiper.ToSnakeCase(field), filter.Sort.Direction))
	}

	// Fetch data
	if err := query.Find(&routes).Error; err != nil {
		return gossiper.PaginatedResult[ent.Route]{}, fmt.Errorf("failed to fetch routes: %w", err)
	}

	return gossiper.NewPaginatedResult(routes, int(count)), nil
}

// GetRoute retrieves a single route by ID.
func (s RouteService) GetRoute(ctx context.Context, id int) (ent.Route, error) {
	var route ent.Route
	result := s.db.GetDB().Preload("Milestones").First(&route, "id = ?", id)
	if result.Error != nil {
		return ent.Route{}, result.Error
	}
	return route, nil
}

// CreateRoute adds a new route along with its milestones.
func (s RouteService) CreateRoute(ctx context.Context, route *ent.Route) (ent.Route, error) {
	err := s.db.GetDB().Transaction(func(tx *gorm.DB) error {
		// Save the route
		if err := tx.Create(route).Error; err != nil {
			return fmt.Errorf("failed to create route: %w", err)
		}

		// Set RouteID for milestones and reset IDs
		for i := range route.Milestones {
			route.Milestones[i].RouteID = route.ID
			route.Milestones[i].ID = 0 // Reset ID to avoid conflicts
		}

		return nil
	})
	if err != nil {
		return ent.Route{}, err
	}
	return *route, nil
}

// DeleteRoute removes a route and its milestones from the database.
func (s RouteService) DeleteRoute(ctx context.Context, id int) (ent.Route, error) {
	var route ent.Route
	result := s.db.GetDB().Preload("Milestones").First(&route, "id = ?", id)
	if result.Error != nil {
		return ent.Route{}, result.Error
	}

	err := s.db.GetDB().Transaction(func(tx *gorm.DB) error {
		// Delete the route
		if err := tx.Delete(&ent.Route{}, id).Error; err != nil {
			return fmt.Errorf("failed to delete route: %w", err)
		}

		// Milestones will be automatically deleted due to the CASCADE constraint
		return nil
	})
	if err != nil {
		return ent.Route{}, err
	}

	return route, nil
}
