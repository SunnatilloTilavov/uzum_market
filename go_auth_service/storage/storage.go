package storage

import (
	auth "auth/genproto/auth_service"
	"context"
	"time"
)

type IStorage interface {
	CloseDB()
	Customer() CustomerStorage
	Seller() SellerStorage
	SystemUser() SystemUserStorage
	Redis() IRedisStorage
}

type CustomerStorage interface {
	GmailCheck(ctx context.Context,req *auth.GmailCheckRequest) (resp *auth.GmailCheckResponse,err error)
	Create(ctx context.Context,req *auth.CreateRequest) (resp *auth.Empty,err error)
	UpdatePassword(ctx context.Context, req * auth.CreateRequest) (resp *auth.Empty,err error)
}

type IRedisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}

type SellerStorage interface{
	SellerGmailCheck(ctx context.Context,req *auth.SellerGmailCheckRequest) (resp *auth.SellerGmailCheckResponse,err error)
	SellerCreate(ctx context.Context,req *auth.SellerCreateRequest) (resp *auth.SellerEmpty,err error)
	SellerUpdatePassword(ctx context.Context, req * auth.SellerCreateRequest) (resp *auth.SellerEmpty,err error)
}

type SystemUserStorage interface{
	SystemUserGmailCheck(ctx context.Context,req *auth.SystemUserGmailCheckRequest) (resp *auth.SystemUserGmailCheckResponse,err error)
	SystemUserCreate(ctx context.Context,req *auth.SystemUserCreateRequest) (resp *auth.SystemUserEmpty,err error)
	SystemUserUpdatePassword(ctx context.Context, req * auth.SystemUserCreateRequest) (resp *auth.SystemUserEmpty,err error)
}