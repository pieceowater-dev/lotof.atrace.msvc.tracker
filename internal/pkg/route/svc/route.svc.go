package svc

import (
	"app/internal/pkg/route/ent"
	"context"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type RouteService struct {
	db gossiper.Database
}

func NewRouteService(db gossiper.Database) *RouteService {
	return &RouteService{db: db}
}

func (s RouteService) GetRoutes(ctx context.Context, filter gossiper.Filter[string]) (gossiper.PaginatedResult[ent.Route], error) {
	//TODO implement me
	panic("implement me")
}

func (s RouteService) GetRoute(ctx context.Context, id int) (*ent.Route, error) {
	//TODO implement me
	panic("implement me")
}

func (s RouteService) CreateRoute(ctx context.Context, request *ent.Route) (*ent.Route, error) {
	//TODO implement me
	panic("implement me")
}

func (s RouteService) DeleteRoute(ctx context.Context, id int) (*ent.Route, error) {
	//TODO implement me
	panic("implement me")
}
