package service

import (
	"context"
	"microservice/config"
	"microservice/genproto/catalog_service"
	"microservice/grpc/client"
	"microservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ProductReviewService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*catalog_service.UnimplementedProductReviewServiceServer
}

func NewProductReviewService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductReviewService {
	return &ProductReviewService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *ProductReviewService) CreateProductReview(ctx context.Context, req *catalog_service.CreateProductReviewRequest) (*catalog_service.ProductReview, error) {
	f.log.Info("---CreateProductReview--->>>", logger.Any("req", req))

	resp, err := f.strg.ProductReview().CreateProductReview(ctx, req)
	if err != nil {
		f.log.Error("---CreateProductReview--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (f *ProductReviewService) GetProductReviewByID(ctx context.Context, req *catalog_service.ProductReviewPrimaryKey) (*catalog_service.ProductReview, error) {
	f.log.Info("---GetProductReviewByID--->>>", logger.Any("req", req))

	resp, err := f.strg.ProductReview().GetProductReviewByID(ctx, req)
	if err != nil {
		f.log.Error("---GetProductReviewByID--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (f *ProductReviewService) UpdateProductReview(ctx context.Context, req *catalog_service.UpdateProductReviewRequest) (*catalog_service.ProductReview, error) {
	f.log.Info("---UpdateProductReview--->>>", logger.Any("req", req))

	resp, err := f.strg.ProductReview().UpdateProductReview(ctx, req)
	if err != nil {
		f.log.Error("---UpdateProductReview--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (f *ProductReviewService) DeleteProductReview(ctx context.Context, req *catalog_service.ProductReviewPrimaryKey) (*catalog_service.Empty4, error) {
	f.log.Info("---DeleteProductReview--->>>", logger.Any("req", req))

	_, err := f.strg.ProductReview().DeleteProductReview(ctx, req)
	if err != nil {
		f.log.Error("---DeleteProductReview--->>>", logger.Error(err))
		return nil, err
	}

	return &catalog_service.Empty4{}, nil
}


func (f *ProductReviewService) GetProductReviewsByProductID(ctx context.Context, req *catalog_service.GetProductReviewsByProductIDRequest) (*catalog_service.GetProductReviewsByProductIDResponse, error) {
	f.log.Info("---GetProductReviewsByProductID--->>>", logger.Any("req", req))

	resp, err := f.strg.ProductReview().GetProductReviewsByProductID(ctx, req)
	if err != nil {
		f.log.Error("---GetProductReviewsByProductID--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (f *ProductReviewService) GetProductReviewsByCustomerID(ctx context.Context, req *catalog_service.GetProductReviewsByCustomerIDRequest) (*catalog_service.GetProductReviewsByCustomerIDResponse, error) {
	f.log.Info("---GetProductReviewsByCustomerID--->>>", logger.Any("req", req))

	resp, err := f.strg.ProductReview().GetProductReviewsByCustomerID(ctx, req)
	if err != nil {
		f.log.Error("---GetProductReviewsByCustomerID--->>>", logger.Error(err))
		return nil, err
	}

	return resp, nil
}
