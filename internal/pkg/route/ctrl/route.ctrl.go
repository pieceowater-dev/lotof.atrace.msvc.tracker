package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/route/ent"
	"app/internal/pkg/route/svc"
	"context"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type RouteController struct {
	routeService *svc.RouteService
	pb.UnimplementedRouteServiceServer
}

func NewRouteController(service *svc.RouteService) *RouteController {
	return &RouteController{routeService: service}
}

// GetRoutes retrieves paginated routes.
func (r RouteController) GetRoutes(ctx context.Context, request *pb.GetRoutesRequest) (*pb.GetRoutesResponse, error) {
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

	paginatedResult, err := r.routeService.GetRoutes(ctx, filter)
	if err != nil {
		return nil, err
	}

	var routes []*pb.Route
	for _, route := range paginatedResult.Rows {
		var milestones []*pb.RouteMilestone
		for _, milestone := range route.Milestones {
			milestones = append(milestones, &pb.RouteMilestone{
				PostId:   uint64(milestone.PostID),
				Priority: int32(milestone.Priority),
			})
		}

		routes = append(routes, &pb.Route{
			Id:         uint64(route.ID),
			Title:      route.Title,
			Milestones: milestones,
		})
	}

	return &pb.GetRoutesResponse{
		Routes: routes,
		PaginationInfo: &pb.PaginationInfo{
			Count: int32(paginatedResult.Info.Count),
		},
	}, nil
}

// GetRoute retrieves a specific route by ID.
func (r RouteController) GetRoute(ctx context.Context, request *pb.GetRouteRequest) (*pb.Route, error) {
	route, err := r.routeService.GetRoute(ctx, int(request.GetId()))
	if err != nil {
		return nil, err
	}

	var milestones []*pb.RouteMilestone
	for _, milestone := range route.Milestones {
		milestones = append(milestones, &pb.RouteMilestone{
			PostId:   uint64(milestone.PostID),
			Priority: int32(milestone.Priority),
		})
	}

	return &pb.Route{
		Id:         uint64(route.ID),
		Title:      route.Title,
		Milestones: milestones,
	}, nil
}

// CreateRoute creates a new route.
func (r RouteController) CreateRoute(ctx context.Context, request *pb.CreateRouteRequest) (*pb.Route, error) {
	var milestones []ent.RouteMilestone
	for _, milestone := range request.GetMilestones() {
		milestones = append(milestones, ent.RouteMilestone{
			PostID:   uint(milestone.PostId),
			Priority: uint(milestone.Priority),
		})
	}

	route := &ent.Route{
		Title:      request.GetTitle(),
		Milestones: milestones,
	}

	createdRoute, err := r.routeService.CreateRoute(ctx, route)
	if err != nil {
		return nil, fmt.Errorf("failed to create route: %w", err)
	}

	var pbMilestones []*pb.RouteMilestone
	for _, milestone := range createdRoute.Milestones {
		pbMilestones = append(pbMilestones, &pb.RouteMilestone{
			PostId:   uint64(milestone.PostID),
			Priority: int32(milestone.Priority),
		})
	}

	return &pb.Route{
		Id:         uint64(createdRoute.ID),
		Title:      createdRoute.Title,
		Milestones: pbMilestones,
	}, nil
}

// todo: implement UpdateRoute later!

// DeleteRoute deletes a route by ID.
func (r RouteController) DeleteRoute(ctx context.Context, request *pb.DeleteRouteRequest) (*pb.Route, error) {
	deletedRoute, err := r.routeService.DeleteRoute(ctx, int(request.GetId()))
	if err != nil {
		return nil, fmt.Errorf("failed to delete route: %w", err)
	}

	var milestones []*pb.RouteMilestone
	for _, milestone := range deletedRoute.Milestones {
		milestones = append(milestones, &pb.RouteMilestone{
			PostId:   uint64(milestone.PostID),
			Priority: int32(milestone.Priority),
		})
	}

	return &pb.Route{
		Id:         uint64(deletedRoute.ID),
		Title:      deletedRoute.Title,
		Milestones: milestones,
	}, nil
}
