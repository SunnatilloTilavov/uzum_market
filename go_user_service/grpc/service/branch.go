package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type BranchService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BranchService {
	return &BranchService{
		cfg:      cfg,	
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *BranchService) Create(ctx context.Context, req *user_service.CreateBranch) (*user_service.BranchPrimaryKey, error) {

	c.log.Info("---CreateBranch--->>>", logger.Any("req", req))

	resp, err := c.strg.Branch().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateBranch--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}	

func (c *BranchService) GetByID(ctx context.Context, req *user_service.BranchPrimaryKey) (*user_service.Branch, error) {
	c.log.Info("---GetByIdBranch--->>>", logger.Any("req", req))

	resp, err := c.strg.Branch().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdBranch--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *BranchService) GetList(ctx context.Context, req *user_service.GetListBranchRequest) (*user_service.GetListBranchResponse, error) {
	c.log.Info("---GetAllBranch--->>>", logger.Any("req", req))

	resp, err := c.strg.Branch().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllBranch--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *BranchService) Update(ctx context.Context,req *user_service.UpdateBranchRequest) (*user_service.UpdateBranchResponse,error) {
	c.log.Info("---UpdateBranch--->>>", logger.Any("req", req))

	resp, err := c.strg.Branch().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateBranch--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *BranchService) Delete(ctx context.Context,req *user_service.BranchPrimaryKey) (*user_service.BranchEmpty,error) {
	c.log.Info("---DeleteBranch--->>>", logger.Any("req", req))

	resp, err := c.strg.Branch().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteBranch--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}