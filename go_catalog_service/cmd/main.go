package main

import (
	"context"
	"microservice/config"
	"microservice/grpc"
	"microservice/grpc/client"
	"microservice/storage/postgres"
	"net"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set logger level based on environment
	var loggerLevel string
	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	// Initialize logger
	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	// Initialize PostgreSQL store
	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	// Initialize gRPC clients
	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	// Get the pgxpool.Pool from pgStore
	db := pgStore.(*postgres.Store).DB()

	// Set up the gRPC server
	grpcServer := grpc.SetUpServer(db, cfg, log, pgStore, svcs)

	// Listen on the specified port
	lis, err := net.Listen("tcp", cfg.ContentGRPCPort)
	if err != nil {
		log.Panic("net.Listen", logger.Error(err))
	}

	log.Info("GRPC: Server is being started...", logger.String("port", cfg.ContentGRPCPort))

	// Serve gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic("grpcServer.Serve", logger.Error(err))
	}
}
