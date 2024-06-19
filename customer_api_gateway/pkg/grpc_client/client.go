package grpc_client

import (
	pc "backend_course/customer_api_gateway/genproto/catalog_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"backend_course/customer_api_gateway/config"
)

// GrpcClientI ...
type GrpcClientI interface {
	CategoryService() pc.CategoryServiceClient
	ProductService() pc.ProductServiceClient
	ProductReviewService() pc.ProductReviewServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port:%s err: %s",
			cfg.CatalogServiceHost, cfg.CatalogServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"category_service":       pc.NewCategoryServiceClient(connCatalog),
			"product_service":        pc.NewProductServiceClient(connCatalog),
			"product_review_service": pc.NewProductReviewServiceClient(connCatalog),
		},
	}, nil
}

func (g *GrpcClient) CategoryService() pc.CategoryServiceClient {
	return g.connections["category_service"].(pc.CategoryServiceClient)
}

func (g *GrpcClient) ProductService() pc.ProductServiceClient {
	return g.connections["product_service"].(pc.ProductServiceClient)
}

func (g *GrpcClient) ProductReviewService() pc.ProductReviewServiceClient {
	return g.connections["product_review_service"].(pc.ProductReviewServiceClient)
}
