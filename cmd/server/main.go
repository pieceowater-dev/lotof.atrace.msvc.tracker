package main

import (
	"app/internal/core/cfg"
	gtw "app/internal/core/grpc/generated"
	"app/internal/core/middleware"
	"app/internal/pkg"
	"context"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
)

func init() {
	// EXAMPLE, todo: make tenants receiving from hub, then schema switching based on token

	// todo: dump this in independent init func later
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().HubApplicationAddr,
	)

	clientConstructor := gtw.NewGatewayServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	c := client.(gtw.GatewayServiceClient)

	ctx := context.Background()

	response, err := grpcTransport.Send(ctx, c, "GatewayNamespacesByApp", &gtw.GatewayNamespacesByAppRequest{AppBundle: cfg.Inst().AppBundleName})
	if err != nil {
		log.Printf("Error sending request: %v", err)
	}

	res, ok := response.(*gtw.GatewayNamespacesByAppResponse)
	if !ok {
		log.Printf("Error converting response to GatewayNamespacesByAppResponse")
	}

	log.Printf("CreateUser response: %v", response)

	var encryptedTenants []gossiper.EncryptedTenant
	for _, tenant := range res.Tenants {
		encryptedTenants = append(encryptedTenants, gossiper.EncryptedTenant{
			Namespace:   tenant.Namespace,
			Credentials: tenant.Credentials,
		})
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

	tm, err := gossiper.NewTenantManager(database.GetDB(), cfg.Inst().AppBundleSecret)
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

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			middleware.GrpcMiddleware(
				middleware.NewMetadataMiddleware(),
			).MiddlewareMethod(),
		),
	)
	reflection.Register(grpcServer)

	serverManager := gossiper.NewServerManager()
	serverManager.AddServer(gossiper.NewGRPCServ(appCfg.GrpcPort, grpcServer, appRouter.InitGRPC))

	serverManager.StartAll()
	defer serverManager.StopAll()
}
