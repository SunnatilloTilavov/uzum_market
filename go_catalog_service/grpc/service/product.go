package service

import (
	"context"
	"microservice/config"
	"microservice/genproto/catalog_service"
	"microservice/grpc/client"
	"microservice/storage"
 
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	db *pgxpool.Pool
	services client.ServiceManagerI
	*catalog_service.UnimplementedProductServiceServer
}

func NewProductService(db *pgxpool.Pool,cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductService {
	return &ProductService{
		cfg:      cfg,
		log:      log,
		db:       db,
		strg:     strg,
		services: srvs,
	}
}

func (f *ProductService) Create(ctx context.Context, req *catalog_service.CreateProduct) (*catalog_service.Product, error) {
    f.log.Info("---CreateProduct--->>>", logger.Any("req", req))

    // Start a new transaction
    tx, err := f.db.Begin(ctx)
    if err != nil {
        f.log.Error("---CreateProduct--->>>", logger.Error(err))
        return nil, err
    }
    defer tx.Rollback(ctx) // Rollback transaction if not committed

    // Create the product
    resp, err := f.strg.Product().Create(ctx, req)
    if err != nil {
        f.log.Error("---CreateProduct--->>>", logger.Error(err))
        return nil, err
    }

    // Iterate through category IDs and create associations
    for _, CategoryId := range req.CategoriesId {
        req1 := &catalog_service.CreateProductCategoryRequest{
            ProductId:  resp.Id,
            CategoryId: CategoryId,
        }
        _, err := f.strg.ProductCategory().CreateProductCategory(ctx, req1)
        if err != nil {
            // Rollback transaction if category creation fails
            f.log.Error("---CreateProduct--->>>", logger.Error(err))
            if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
                f.log.Error("---CreateProduct--->>> Rollback failed", logger.Error(rollbackErr))
            }
            return nil, err
        }
    }

    // Commit transaction if all operations are successful
    if err = tx.Commit(ctx); err != nil {
        f.log.Error("---CreateProduct--->>>", logger.Error(err))
        return nil, err
    }

    return resp, nil
}

func (f *ProductService) GetByID(ctx context.Context, req *catalog_service.ProductPrimaryKey) (resp *catalog_service.Product, err error) {
	f.log.Info("---GetSingleProduct--->>>", logger.Any("req", req))

	resp, err = f.strg.Product().GetByID(ctx,req)
	if err != nil {
		f.log.Error("---GetSingleProduct--->>>", logger.Error(err))
		return &catalog_service.Product{}, err
	}

	return resp, nil
}


func (f *ProductService) GetAll(ctx context.Context, req *catalog_service.GetListProductRequest) (resp *catalog_service.GetListProductResponse, err error) {
	f.log.Info("---GetAllCategories--->>>", logger.Any("req", req))

	resp, err = f.strg.Product().GetList(ctx, req)
	if err != nil {
		f.log.Error("---GetAllCategories--->>>", logger.Error(err))
		return &catalog_service.GetListProductResponse{}, err
	}

	return resp, nil
}

// func (f *ProductService) Update(ctx context.Context, req *catalog_service.UpdateProduct) (*catalog_service.Product, error) {
// 	f.log.Info("---UpdateProduct--->>>", logger.Any("req", req))

// 	// Update the product
// 	resp, err := f.strg.Product().Update(ctx, req)
// 	if err != nil {
// 		f.log.Error("---UpdateProduct--->>>", logger.Error(err))
// 		return nil, err
// 	}


// 	req1 := &catalog_service.DeleteProductCategoryRequest{
// 		ProductId: req.Id,
// 	}
// 	_, err = f.strg.ProductCategory().DeleteProductCategory(ctx, req1)
// 	if err != nil {
// 		f.log.Error("---DeleteProductCategory--->>>", logger.Error(err))
// 		return nil, err
// 	}

// 	for _, CategoryId := range req.CategoriesId {
// 		req2 := &catalog_service.CreateProductCategoryRequest{
// 			ProductId:  resp.Id,
// 			CategoryId: CategoryId,
// 		}
// 		_, err := f.strg.ProductCategory().CreateProductCategory(ctx, req2)
// 		if err != nil {
// 			// Rollback transaction if category creation fails
// 			f.log.Error("---CreateProductCategory--->>>", logger.Error(err))
// 			return nil, err
// 		}
// 	}

// 	return resp, nil
// }


func (f *ProductService) Update(ctx context.Context, req *catalog_service.UpdateProduct) (*catalog_service.Product, error) {
	f.log.Info("---UpdateProduct--->>>", logger.Any("req", req))

	// Start a new transaction
	tx, err := f.db.Begin(ctx)
	if err != nil {
		f.log.Error("---UpdateProduct--->>>", logger.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			f.log.Error("---UpdateProduct--->>>", logger.Error(err))
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				f.log.Error("---UpdateProduct--->>>", logger.Error(err))
			}
		}
	}()

	// Update the product
	resp, err := f.strg.Product().Update(ctx, req)
	if err != nil {
		return nil, err
	}

	// Delete existing product categories
	req1 := &catalog_service.DeleteProductCategoryRequest{
		ProductId: req.Id,
	}
	_, err = f.strg.ProductCategory().DeleteProductCategory(ctx, req1)
	if err != nil {
		return nil, err
	}

	// Create new product categories
	for _, CategoryId := range req.CategoriesId {
		req2 := &catalog_service.CreateProductCategoryRequest{
			ProductId:  resp.Id,
			CategoryId: CategoryId,
		}
		_, err := f.strg.ProductCategory().CreateProductCategory(ctx, req2)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}


func (f *ProductService) Delete(ctx context.Context, req *catalog_service.ProductPrimaryKey) (*catalog_service.Empty1, error) {
	f.log.Info("---DeleteProduct--->>>", logger.Any("req", req))

	// Start a new transaction
	tx, err := f.db.Begin(ctx)
	if err != nil {
		f.log.Error("---DeleteProduct--->>>", logger.Error(err))
		return &catalog_service.Empty1{}, err
	}
	defer tx.Rollback(ctx) // Rollback transaction if not committed

	// Delete the product
	_, err = f.strg.Product().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteProduct--->>>", logger.Error(err))
		return &catalog_service.Empty1{}, err
	}

	// Delete associated product categories
	req1 := &catalog_service.DeleteProductCategoryRequest{
		ProductId: req.Id,
	}
	_, err = f.strg.ProductCategory().DeleteProductCategory(ctx, req1)
	if err != nil {
		f.log.Error("---DeleteProductCategory--->>>", logger.Error(err))
		return &catalog_service.Empty1{}, err
	}

	// Commit the transaction
	if err = tx.Commit(ctx); err != nil {
		f.log.Error("---DeleteProduct--->>>", logger.Error(err))
		return &catalog_service.Empty1{}, err
	}

	return &catalog_service.Empty1{}, nil
}



func (f *ProductService) GetList(ctx context.Context, req *catalog_service.GetListProductRequest) (resp *catalog_service.GetListProductResponse, err error) {
    f.log.Info("---GetAllCategories--->>>", logger.Any("req", req))

    listResp, err := f.strg.Product().GetList(ctx,req)
    if err != nil {
        f.log.Error("---GetAllCategories--->>>", logger.Error(err))
        return &catalog_service.GetListProductResponse{}, err
    }

    return listResp, nil
}







