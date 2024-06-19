package client

import (
	"auth/config"
	"auth/genproto/user_service"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	CustomerService() user_service.CustomerServiceClient
	SystemUserService() user_service.SystemUserServiceClient
	SellerService() user_service.SellerServiceClient
}

type grpcClients struct {
	customerService user_service.CustomerServiceClient
	system_userService user_service.SystemUserServiceClient
	sellerService user_service.SellerServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	// Правильное формирование адреса сервиса
	connCustomerService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort), // Исправлено
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		customerService: user_service.NewCustomerServiceClient(connCustomerService),
		system_userService: user_service.NewSystemUserServiceClient(connCustomerService),
		sellerService: user_service.NewSellerServiceClient(connCustomerService),
	}, nil
}

func (g *grpcClients) CustomerService() user_service.CustomerServiceClient {
	return g.customerService
}

func (g *grpcClients) SystemUserService() user_service.SystemUserServiceClient {
	return g.system_userService
}

func (g *grpcClients) SellerService() user_service.SellerServiceClient {
	return g.sellerService
}