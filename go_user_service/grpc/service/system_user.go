package service

import (
	"context"
	"user_service/config"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type SystemUserService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewSystemUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SystemUserService {
	return &SystemUserService{
		cfg:      cfg,	
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c *SystemUserService) Create(ctx context.Context, req *user_service.CreateSystemUser) (*user_service.SystemUserPrimaryKey, error) {

	c.log.Info("---CreateSystemUser--->>>", logger.Any("req", req))

	resp,err:= c.strg.SystemUser().Create(ctx, req)
	if err != nil {
		c.log.Error("---CreateSystemUser--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}	

func (c *SystemUserService) GetByID(ctx context.Context, req *user_service.SystemUserPrimaryKey) (*user_service.SystemUser, error) {
	c.log.Info("---GetByIdSystemUser--->>>", logger.Any("req", req))

	resp,err:= c.strg.SystemUser().GetByID(ctx, req)
	if err != nil {
		c.log.Error("---GetByIdSystemUser--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SystemUserService) GetList(ctx context.Context, req *user_service.GetListSystemUserRequest) (*user_service.GetListSystemUserResponse, error) {
	c.log.Info("---GetAllSystemUser--->>>", logger.Any("req", req))

	resp,err:= c.strg.SystemUser().GetList(ctx, req)
	if err != nil {
		c.log.Error("---GetAllSystemUser--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SystemUserService) Update(ctx context.Context,req *user_service.UpdateSystemUserRequest) (*user_service.UpdateSystemUserResponse,error) {
	c.log.Info("---UpdateSystemUser--->>>", logger.Any("req", req))

	resp,err:= c.strg.SystemUser().Update(ctx, req)
	if err != nil {
		c.log.Error("---UpdateSystemUser--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SystemUserService) Delete(ctx context.Context,req *user_service.SystemUserPrimaryKey) (*user_service.SystemUserEmpty,error) {
	c.log.Info("---DeleteSystemUser--->>>", logger.Any("req", req))

	resp,err:= c.strg.SystemUser().Delete(ctx, req)
	if err != nil {
		c.log.Error("---DeleteSystemUser--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (c *SystemUserService) GetByGmail(ctx context.Context,req *user_service.SystemUserGmail) (*user_service.SystemUserPrimaryKey,error) {
	c.log.Info("---GetByGmailSystemUser--->>>", logger.Any("req", req))

	resp, err := c.strg.SystemUser().GetByGmail(ctx, req)
	if err != nil {
		c.log.Error("---GetByGmailSystemUser--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}