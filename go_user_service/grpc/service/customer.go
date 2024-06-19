package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type CustomerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewCustomerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CustomerService {
	return &CustomerService{
		cfg:      cfg,	
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *CustomerService) Create(ctx context.Context, req *user_service.CreateCustomer) (*user_service.CustomerPrimaryKey, error) {

	c.log.Info("---CreateCustomer--->>>", logger.Any("req", req))

	resp, err := c.strg.Customer().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateCustomer--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}	

func (c *CustomerService) GetByID(ctx context.Context, req *user_service.CustomerPrimaryKey) (*user_service.Customer, error) {
	c.log.Info("---GetByIdCustomer--->>>", logger.Any("req", req))

	resp, err := c.strg.Customer().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdCustomer--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *CustomerService) GetList(ctx context.Context, req *user_service.GetListCustomerRequest) (*user_service.GetListCustomerResponse, error) {
	c.log.Info("---GetAllCustomer--->>>", logger.Any("req", req))

	resp, err := c.strg.Customer().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllCustomer--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *CustomerService) Update(ctx context.Context,req *user_service.UpdateCustomerRequest) (*user_service.UpdateCustomerResponse,error) {
	c.log.Info("---UpdateCustomer--->>>", logger.Any("req", req))

	resp, err := c.strg.Customer().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateCustomer--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *CustomerService) Delete(ctx context.Context,req *user_service.CustomerPrimaryKey) (*user_service.Empty,error) {
	c.log.Info("---DeleteCustomer--->>>", logger.Any("req", req))

	resp, err := c.strg.Customer().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteCustomer--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *CustomerService) GetByGmail(ctx context.Context,req *user_service.CustomerGmail) (*user_service.CustomerPrimaryKey,error) {
	c.log.Info("---GetByGmailCustomer--->>>", logger.Any("req", req))

	resp, err := c.strg.Customer().GetByGmail(ctx, req)
	if err != nil {
		c.log.Error("---GetByGmailCustomer--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}