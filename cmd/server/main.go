package main

import (
	"app/internal/core/cfg"
	"app/internal/pkg"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
)

func init() {
	// EXAMPLE, todo - finish this
	encryptedTenants := []gossiper.EncryptedTenant{
		{
			Namespace:   "someSchema",
			Credentials: "bIjCI1Y60pThND2uESKYeuWUyf3MIEOfFKn9",
		},
	}

	database, err := gossiper.NewDB(
		gossiper.PostgresDB,
		cfg.Inst().PostgresDatabaseDSN,
		false,
		[]any{},
	)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	tm, err := gossiper.NewTenantManager(database.GetDB(), cfg.Inst().Secret)
	if err != nil {
		log.Fatalf("Failed to create TenantManager: %v", err)
	}

	err = tm.SyncTenants(&encryptedTenants)
	if err != nil {
		log.Fatalf("Failed to sync tenants: %v", err)
	}
}

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
