package service

import (
	"context"
	"microservice/config"
	"microservice/genproto/catalog_service"
	"microservice/grpc/client"
	"microservice/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type CategoryService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*catalog_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CategoryService {
	return &CategoryService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}
func (f *CategoryService) Create(ctx context.Context, req *catalog_service.CreateCategory) (resp *catalog_service.Category, err error) {

	f.log.Info("---CreateCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateCategory--->>>", logger.Error(err))
		return &catalog_service.Category{}, err
	}

	return resp, nil
}

func (f *CategoryService) GetByID(ctx context.Context, req *catalog_service.CategoryPrimaryKey) (resp *catalog_service.GetCategory, err error) {
	f.log.Info("---GetSingleCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().GetByID(ctx,req)
	if err != nil {
		f.log.Error("---GetSingleCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}


func (f *CategoryService) GetAll(ctx context.Context, req *catalog_service.GetListCategoryRequest) (resp *catalog_service.GetListCategoryResponse, err error) {
	f.log.Info("---GetAllCategories--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().GetList(ctx, req)
	if err != nil {
		f.log.Error("---GetAllCategories--->>>", logger.Error(err))
		return &catalog_service.GetListCategoryResponse{}, err
	}

	return resp, nil
}

func (f *CategoryService) Update(ctx context.Context, req *catalog_service.UpdateCategory) (resp *catalog_service.GetCategory, err error) {
	f.log.Info("---UpdateCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}

func (f *CategoryService) Delete(ctx context.Context, req *catalog_service.CategoryPrimaryKey) (resp *catalog_service.Empty, err error) {
	f.log.Info("---DeleteCategory--->>>", logger.Any("req", req))

	_,err = f.strg.Category().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteCategory--->>>", logger.Error(err))
		return &catalog_service.Empty{}, err
	}

	return &catalog_service.Empty{}, nil
}


func (f *CategoryService) GetList(ctx context.Context, req *catalog_service.GetListCategoryRequest) (resp *catalog_service.GetListCategoryResponse, err error) {
    f.log.Info("---GetAllCategories--->>>", logger.Any("req", req))

    // strg.Category().GetList metodi ochilmasa, uni ochish uchun bitta client.ServiceManagerI obyektini qo'llab-quvvatlaydigan qadam qo'shing
    listResp, err := f.strg.Category().GetList(ctx,req)
    if err != nil {
        f.log.Error("---GetAllCategories--->>>", logger.Error(err))
        return &catalog_service.GetListCategoryResponse{}, err
    }

    return listResp, nil
}


func (f *CategoryService) GetCategoryWithProductId(ctx context.Context, req *catalog_service.CategoryPrimaryKey) (resp *catalog_service.GetCategory, err error){
	f.log.Info("---GetSingleCategory--->>>", logger.Any("req", req))

	resp, err = f.strg.Category().GetCategoryWithProductId(ctx,req)
	if err != nil {
		f.log.Error("---GetSingleCategory--->>>", logger.Error(err))
		return &catalog_service.GetCategory{}, err
	}

	return resp, nil
}





