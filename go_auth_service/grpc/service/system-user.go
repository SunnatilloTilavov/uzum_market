package service

import (
	"auth/config"
	"auth/genproto/auth_service"
	"auth/genproto/user_service"
	"auth/grpc/client"
	"auth/pkg/hash"
	smtp "auth/pkg/helper"
	"auth/pkg/jwt"
	"auth/storage"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/saidamir98/udevs_pkg/logger"
)

type SystemUserService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.IStorage
	services client.ServiceManagerI
	*auth_service.UnimplementedSystemUserAuthServer
}

func NewSystemUserAuthService(cfg config.Config, log logger.LoggerI, strg storage.IStorage, srvs client.ServiceManagerI) *SystemUserService {
	return &SystemUserService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (c SystemUserService) SystemUserLoginByPassword(ctx context.Context, req *auth_service.SystemUserLoginRequest) (*auth_service.SystemUserLoginResponse, error) {

	c.log.Info("---SystemUserLoginByPassword-->>>", logger.Any("req", req))

	resp:= &auth_service.SystemUserLoginResponse{}

	data, err := c.strg.SystemUser().SystemUserGmailCheck(ctx, &auth_service.SystemUserGmailCheckRequest{Gmail: req.Gmail})
	if err != nil {
		c.log.Error("---SystemUserLogin--->>>", logger.Error(err))
		return nil, err
	}

	if err = hash.CompareHashAndPassword(data.Password, req.Password); err != nil {
		c.log.Error("---SystemUserLogin--->>>", logger.Error(err))
		return nil, err
	}

	m := make(map[interface{}]interface{})
	id,err:=c.services.SystemUserService().GetByGmail(ctx,&user_service.SystemUserGmail{Gmail: req.Gmail})
	if err != nil {
		return nil, err
	}
	m["user_id"] = id
	m["user_role"] = config.SYSTEM_TYPE
	accesstoken, refreshtoken, err := jwt.GenJWT(m)
	if err != nil {
		c.log.Error("---SystemUserLogin--->>>", logger.Error(err))
		return nil, err
	}

	resp.Accesstoken = accesstoken
	resp.Refreshtoken = refreshtoken

	return resp, nil
}

func (c SystemUserService) SystemUserRegisterByMail(ctx context.Context, req *auth_service.SystemUserGmailCheckRequest) (*auth_service.SystemUserEmpty, error) {
	c.log.Info("---SystemUserRegisterByMail--->>>", logger.Any("req", req))
	resp:=&auth_service.SystemUserEmpty{}

	password, _ := c.strg.SystemUser().SystemUserGmailCheck(ctx,&auth_service.SystemUserGmailCheckRequest{Gmail: req.Gmail})
	if password==nil {
		otp := smtp.GenerateOTP()
		msg := fmt.Sprintf("Your OTP: %v. DON'T give anyone", otp)
		err := c.strg.Redis().SetX(ctx, req.Gmail, otp, time.Minute*2)
		if err != nil {
			return resp,err
		}

		err = smtp.SendMail(req.Gmail, msg)
		if err != nil {
			return resp,err
		}
	}else {
		return resp,errors.New("you are already registered")
	}

	return resp, nil
}

func (c SystemUserService) SystemUserRegisterByMailConfirm(ctx context.Context, req *auth_service.SystemUserRConfirm) (*auth_service.RespRegSeller, error) {
	resp:=&auth_service.RespRegSeller{}
	validOtp:=c.strg.Redis().Get(ctx,req.Gmail)
	if validOtp!=req.Otp {
		c.log.Error("---SystemUserConfirmByMail--->>>", logger.Error(errors.New("wrong otp")))
		return nil,errors.New("wrong otp")
	}
	hashedPassword,err:=hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	_,err=c.strg.SystemUser().SystemUserCreate(ctx,&auth_service.SystemUserCreateRequest{Password:hashedPassword,Gmail: req.Gmail })
	if err != nil {
		c.log.Error("---SystemUserConfirmByMail--->>>", logger.Error(err))
		return nil, err
	}

	primaryKey,err:=c.services.SystemUserService().Create(ctx,&user_service.CreateSystemUser{
		Gmail: req.Gmail,
		Role: req.Role,
	})

	resp=&auth_service.RespRegSeller{Id: primaryKey.Id}

	return resp,nil
}

func (c SystemUserService) SystemUserLoginByGmail(ctx context.Context,req *auth_service.SystemUserGmailCheckRequest) (*auth_service.SystemUserEmpty, error) {
	_,err:=c.strg.SystemUser().SystemUserGmailCheck(ctx,&auth_service.SystemUserGmailCheckRequest{Gmail: req.Gmail})
	resp:=&auth_service.SystemUserEmpty{}
	if err!=sql.ErrNoRows {
		otp:=smtp.GenerateOTP()
		err:=c.strg.Redis().SetX(ctx,req.Gmail,otp,time.Minute*2)
		if err != nil {
			return resp, err
		}
		msg := fmt.Sprintf("Your OTP: %v. DON'T give anyone", otp)
		err = smtp.SendMail(req.Gmail, msg)
		if err != nil {
			return resp,err
		}
	}

	return resp,nil
}


func (c SystemUserService) SystemUserLoginByGmailComfirm(ctx context.Context,req *auth_service.SystemUserLoginByGmailRequest) (*auth_service.SystemUserLoginResponse,error) {
	resp:=&auth_service.SystemUserLoginResponse{}
	_,err:=c.strg.SystemUser().SystemUserGmailCheck(ctx,&auth_service.SystemUserGmailCheckRequest{Gmail: req.Gmail})
	if err==sql.ErrNoRows {
		return nil,errors.New("you are not registered")
	}

	validOtp:=c.strg.Redis().Get(ctx,req.Gmail)
	if validOtp!=req.Otp {
		return nil,errors.New("wrong otp")
	}
	id,err:=c.services.SystemUserService().GetByGmail(ctx,&user_service.SystemUserGmail{Gmail: req.Gmail})
	if err != nil {
		return nil, err
	}

	m := make(map[interface{}]interface{})
	m["user_id"] = id
	m["user_role"] = config.SYSTEM_TYPE
	accesstoken, refreshtoken, err := jwt.GenJWT(m)
	if err != nil {
		c.log.Error("---SystemUserLoginByMailConfirm--->>>", logger.Error(err))
		return nil, err
	}

	resp.Accesstoken = accesstoken
	resp.Refreshtoken = refreshtoken

	return resp,nil
}

func (c SystemUserService) SystemUserResetPassword(ctx context.Context, req *auth_service.SystemUserGmailCheckRequest) (*auth_service.SystemUserEmpty, error) {
	c.log.Info("---SystemUserResetPassword--->>>", logger.Any("req", req))
	resp:=&auth_service.SystemUserEmpty{}

	otp := smtp.GenerateOTP()
	msg := fmt.Sprintf("Your OTP: %v. DON'T give anyone", otp)
	err := c.strg.Redis().SetX(ctx, req.Gmail, otp, time.Minute*2)
	if err != nil {
		return resp, err
	}

	err = smtp.SendMail(req.Gmail, msg)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c SystemUserService) SystemUserResetPasswordConfirm(ctx context.Context, req *auth_service.SystemUserPasswordConfirm) (*auth_service.SystemUserEmpty, error) {
	resp := &auth_service.SystemUserEmpty{}
	validOtp := c.strg.Redis().Get(ctx, req.Gmail)
	if validOtp != req.Otp {
		return resp, errors.New("invalid otp")
	}

	resp, err := c.strg.SystemUser().SystemUserUpdatePassword(ctx, &auth_service.SystemUserCreateRequest{Gmail: req.Gmail, Password: req.Password})
	if err != nil {
		return resp, nil
	}

	return resp, nil
}