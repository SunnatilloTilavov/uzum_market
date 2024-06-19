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
		log.Error("grpc dial error", logger.Error(err))
	}
}

func main() {
	initDeps()

	server := api.New(api.Config{
		Logger:     log,
		GrpcClient: grpcClient,
		Cfg:        cfg,
	})

	server.Run(cfg.HTTPPort)
}

// go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"