package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/route/svc"
	"context"
)

type RouteController struct {
	routeService *svc.RouteService
	pb.UnimplementedRouteServiceServer
}

func NewRouteController(service *svc.RouteService) *RouteController {
	return &RouteController{routeService: service}
}

func (r RouteController) GetRoutes(ctx context.Context, request *pb.GetRoutesRequest) (*pb.GetRoutesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RouteController) GetRoute(ctx context.Context, request *pb.GetRouteRequest) (*pb.Route, error) {
	//TODO implement me
	panic("implement me")
}

func (r RouteController) CreateRoute(ctx context.Context, request *pb.CreateRouteRequest) (*pb.Route, error) {
	//TODO implement me
	panic("implement me")
}

func (r RouteController) DeleteRoute(ctx context.Context, request *pb.DeleteRouteRequest) (*pb.Route, error) {
	//TODO implement me
	panic("implement me")
}
