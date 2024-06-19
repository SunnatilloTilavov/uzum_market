package grpc_client

import (
	ps "backend_course/customer_api_gateway/genproto/auth_service"
	pc "backend_course/customer_api_gateway/genproto/catalog_service"
	pa "backend_course/customer_api_gateway/genproto/order_notes"
	pd "backend_course/customer_api_gateway/genproto/order_product_service"
	pb "backend_course/customer_api_gateway/genproto/order_service"
	pq "backend_course/customer_api_gateway/genproto/user_service"
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
	OrderService() pb.OrderServiceClient
	OrderProductsService() pd.OrderProductsClient
	OrderProductNotesService() pa.OrderStatusNotesClient
	BranchBranch() pq.BranchServiceClient
	CustomerService() pq.CustomerServiceClient
	SellerAuthService() ps.SellerAuthClient
	CustomerAuthService() ps.CustomerAuthClient
	SystemUserAuthService() ps.SystemUserAuthClient
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

	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("order service dial host: %s port:%s err: %s",
			cfg.OrderServiceHost, cfg.OrderServicePort, err)
	}

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	connAuth, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.AuthServiceHost, cfg.AuthServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
			cfg.AuthServiceHost, cfg.AuthServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"category_service":       pc.NewCategoryServiceClient(connCatalog),
			"product_service":        pc.NewProductServiceClient(connCatalog),
			"product_review_service": pc.NewProductReviewServiceClient(connCatalog),
			"order_service":          pb.NewOrderServiceClient(connOrder),
			"order_product_service":  pd.NewOrderProductsClient(connOrder),
			"order_notes":            pa.NewOrderStatusNotesClient(connOrder),
			"branch":                 pq.NewBranchServiceClient(connUser),
			"customerservice":        pq.NewCustomerServiceClient(connUser),
			"sellerservice":          pq.NewSellerServiceClient(connUser),
			"shopservice":            pq.NewShopServiceClient(connUser),
			"SystemUserService":      pq.NewSystemUserServiceClient(connUser),
			"SellerAuthClient":       ps.NewCustomerAuthClient(connAuth),
			"CustomerAuthService":    ps.NewSellerAuthClient(connAuth),
			"SystemUserAuthService":  ps. NewSystemUserAuthClient(connAuth),
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

func (g *GrpcClient) OrderService() pb.OrderServiceClient {
	return g.connections["order_service"].(pb.OrderServiceClient)
}

func (g *GrpcClient) OrderProductsService() pd.OrderProductsClient {
	return g.connections["order_product_service"].(pd.OrderProductsClient)
}

func (g *GrpcClient) OrderProductNotesService() pa.OrderStatusNotesClient {
	return g.connections["order_notes"].(pa.OrderStatusNotesClient)
}

func (g *GrpcClient) BranchBranch() pq.BranchServiceClient {
	return g.connections["branch"].(pq.BranchServiceClient)
}

func (g *GrpcClient) CustomerService() pq.CustomerServiceClient {
	return g.connections["customerservice"].(pq.CustomerServiceClient)
}

func (g *GrpcClient) SellerService() pq.SellerServiceClient {
	return g.connections["sellerservice"].(pq.SellerServiceClient)
}

func (g *GrpcClient) ShopService() pq.ShopServiceClient {
	return g.connections["shopservice"].(pq.ShopServiceClient)
}

func (g *GrpcClient) SystemUserService() pq.SystemUserServiceClient {
	return g.connections["SystemUserService"].(pq.SystemUserServiceClient)
}

func (g *GrpcClient) SellerAuthService() ps.SellerAuthClient {
	return g.connections["SellerAuthClient"].(ps.SellerAuthClient)
}

func (g *GrpcClient) CustomerAuthService() ps.CustomerAuthClient {
	return g.connections["CustomerAuthService"].(ps.CustomerAuthClient)
}

func (g *GrpcClient) SystemUserAuthService() ps.SystemUserAuthClient {
	return g.connections["SystemUserAuthService"].(ps.SystemUserAuthClient)
}
