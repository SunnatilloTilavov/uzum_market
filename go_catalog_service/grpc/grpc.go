package grpc

import (
	"microservice/config"
	"microservice/genproto/catalog_service"
	"microservice/grpc/client"
	"microservice/grpc/service"
	"microservice/storage"
"github.com/jackc/pgx/v4/pgxpool"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(db *pgxpool.Pool,cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	catalog_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg, srvc))
	catalog_service.RegisterProductServiceServer(grpcServer, service.NewProductService(db,cfg, log, strg, srvc))
	catalog_service.RegisterProductReviewServiceServer(grpcServer, service.NewProductReviewService(cfg, log, strg, srvc))
	reflection.Register(grpcServer)
	return
}
