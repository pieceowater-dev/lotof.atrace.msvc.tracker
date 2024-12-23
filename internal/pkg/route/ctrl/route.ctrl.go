package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/route/svc"
)

type RouteController struct {
	routeService *svc.RouteService
	pb.UnimplementedRouteServiceServer
}

func NewRouteController(service *svc.RouteService) *RouteController {
	return &RouteController{routeService: service}
}
