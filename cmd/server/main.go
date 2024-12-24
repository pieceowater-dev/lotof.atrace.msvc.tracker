package main

import (
	"app/internal/core/cfg"
	"app/internal/pkg"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	appCfg := cfg.Inst()
	appRouter := pkg.NewRouter()

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	serverManager := gossiper.NewServerManager()
	serverManager.AddServer(gossiper.NewGRPCServ(appCfg.GrpcPort, grpcServer, appRouter.InitGRPC))

	serverManager.StartAll()
	defer serverManager.StopAll()
}
