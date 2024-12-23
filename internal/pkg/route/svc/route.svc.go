package svc

import gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"

type RouteService struct {
	db gossiper.Database
}

func NewRouteService(db gossiper.Database) *RouteService {
	return &RouteService{db: db}
}
