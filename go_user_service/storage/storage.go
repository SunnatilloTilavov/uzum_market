package storage

import (
	"context"
	ct "user_service/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	Customer() CustomerRepoI
	Shop() ShopRepoI
	Seller() SellerRepoI
	Branch() BranchRepoI
	SystemUser() SystemUserRepoI
}

type CustomerRepoI interface {
	Create(ctx context.Context,req *ct.CreateCustomer) (resp *ct.CustomerPrimaryKey,err error)
	GetByID(ctx context.Context, req *ct.CustomerPrimaryKey) (resp *ct.Customer, err error)
	GetList(ctx context.Context, req *ct.GetListCustomerRequest) (resp *ct.GetListCustomerResponse, err error)
	Update(ctx context.Context,req *ct.UpdateCustomerRequest) (resp *ct.UpdateCustomerResponse,err error)
	Delete(ctx context.Context,req *ct.CustomerPrimaryKey) (resp *ct.Empty,err error)
	GetByGmail(ctx context.Context,req *ct.CustomerGmail) (*ct.CustomerPrimaryKey,error)
}

type ShopRepoI interface{
	Create(ctx context.Context,req *ct.CreateShop) (resp *ct.ShopPrimaryKey,err error)
	GetById(ctx context.Context, req *ct.ShopPrimaryKey) (resp *ct.GetByID,err error)
	Update(ctx context.Context, req *ct.UpdateShopRequest) (resp *ct.ShopEmpty, err error)
	Delete(ctx context.Context,req *ct.ShopPrimaryKey) (resp *ct.ShopEmpty,err error)
	GetList(ctx context.Context,req *ct.GetListShopRequest) (resp *ct.GetListShopResponse,err error)
}

type SellerRepoI interface{
	Create(ctx context.Context,req *ct.CreateSeller) (resp *ct.SellerPrimaryKey,err error)	
	GetByID(ctx context.Context, req *ct.SellerPrimaryKey) (resp *ct.Seller, err error)
	Update(ctx context.Context, req *ct.UpdateSellerRequest) (resp *ct.UpdateSellerResponse, err error)
	Delete(ctx context.Context, req *ct.SellerPrimaryKey) (resp *ct.SellerEmpty, err error)
	GetList(ctx context.Context,req *ct.GetListSellerRequest) (resp *ct.GetListSellerResponse,err error)
	GetByGmail(ctx context.Context,req *ct.SellerGmail) (*ct.SellerPrimaryKey,error)
}

type BranchRepoI interface{
	Create(ctx context.Context,req *ct.CreateBranch) (resp *ct.BranchPrimaryKey,err error)
	GetByID(ctx context.Context, req *ct.BranchPrimaryKey) (resp *ct.Branch, err error)
	Update(ctx context.Context, req *ct.UpdateBranchRequest) (resp *ct.UpdateBranchResponse, err error)
	Delete(ctx context.Context, req *ct.BranchPrimaryKey) (resp *ct.BranchEmpty, err error)
	GetList(ctx context.Context,req *ct.GetListBranchRequest) (resp *ct.GetListBranchResponse,err error)
}

type SystemUserRepoI interface{
	Create(ctx context.Context,req *ct.CreateSystemUser) (resp *ct.SystemUserPrimaryKey,err error)
	GetByID(ctx context.Context, req *ct.SystemUserPrimaryKey) (resp *ct.SystemUser, err error)
	Update(ctx context.Context, req *ct.UpdateSystemUserRequest) (resp *ct.UpdateSystemUserResponse, err error)
	Delete(ctx context.Context, req *ct.SystemUserPrimaryKey) (resp *ct.SystemUserEmpty, err error)
	GetList(ctx context.Context,req *ct.GetListSystemUserRequest) (resp *ct.GetListSystemUserResponse,err error)
	GetByGmail(ctx context.Context,req *ct.SystemUserGmail) (*ct.SystemUserPrimaryKey,error)
}