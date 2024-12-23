package pkg

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/post"
	"app/internal/pkg/record"
	"app/internal/pkg/route"

	"google.golang.org/grpc"
)

type Router struct {
	postModule   *post.Module
	recordModule *record.Module
	routeModule  *route.Module
}

func NewRouter() *Router {
	return &Router{
		postModule:   post.New(),
		recordModule: record.New(),
		routeModule:  route.New(),
	}
}

// InitGRPC initializes gRPC routes
func (r *Router) InitGRPC(grpcServer *grpc.Server) {
	// Register gRPC services
	pb.RegisterPostServiceServer(grpcServer, r.postModule.Controller)
	pb.RegisterRecordServiceServer(grpcServer, r.recordModule.Controller)
	pb.RegisterRouteServiceServer(grpcServer, r.routeModule.Controller)
}
