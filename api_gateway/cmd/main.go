package main

import (
	"backend_course/customer_api_gateway/api"
	"backend_course/customer_api_gateway/config"
	"backend_course/customer_api_gateway/pkg/grpc_client"
	"backend_course/customer_api_gateway/pkg/logger"
)

var (
	log        logger.Logger
	cfg        config.Config
	grpcClient *grpc_client.GrpcClient
)

func initDeps() {
	var err error
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "customer-api-gateway")

	grpcClient, err = grpc_client.New(cfg)
	if err != nil {
		log.Fatal("failed to initialize gRPC clients", logger.Error(err))
	}
	log.Info("Dependencies initialized successfully")
}

func main() {
	initDeps()

	server := api.New(api.Config{
		Logger:     log,
		GrpcClient: grpcClient,
		Cfg:        cfg,
	})

	log.Info("Starting server on port", logger.String("port", cfg.HTTPPort))
	server.Run(cfg.HTTPPort)
}
