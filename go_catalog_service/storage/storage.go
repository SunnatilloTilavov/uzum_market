package storage

import (
	"context"
	ct "microservice/genproto/catalog_service"
	// "github.com/google/uuid"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Product() ProductRepoI
	ProductCategory() ProductCategoryRepoI
	ProductReview() ProductReviewRepoI
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.Category, err error)
	GetByID(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.GetCategory, err error)
	Update(ctx context.Context, req *ct.UpdateCategory) (resp *ct.GetCategory, err error)
	Delete(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.Empty, err error)
	GetList(ctx context.Context, req *ct.GetListCategoryRequest) (resp *ct.GetListCategoryResponse, err error)
	GetCategoryWithProductId(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.GetCategory, err error)
}

type ProductRepoI interface {
	Create(ctx context.Context, req *ct.CreateProduct) (resp *ct.Product, err error)
	GetByID(ctx context.Context, req *ct.ProductPrimaryKey) (resp *ct.Product, err error)
	GetList(ctx context.Context, req *ct.GetListProductRequest) (resp *ct.GetListProductResponse, err error)
	Update(ctx context.Context, req *ct.UpdateProduct) (resp *ct.Product, err error)
	Delete(ctx context.Context, req *ct.ProductPrimaryKey) (resp *ct.Empty, err error)
}

type ProductCategoryRepoI interface {
	CreateProductCategory(ctx context.Context, req *ct.CreateProductCategoryRequest) (resp *ct.ProductCategory, err error)
	GetProductCategoriesByProductID(ctx context.Context, req *ct.GetProductCategoriesByProductIDRequest) (resp *ct.GetProductCategoriesByProductIDResponse, err error)
	GetProductCategoriesByCategoryID(ctx context.Context, req *ct.GetProductCategoriesByCategoryIDRequest) (resp *ct.GetProductCategoriesByCategoryIDResponse, err error)
	DeleteProductCategory(ctx context.Context, req *ct.DeleteProductCategoryRequest) (resp *ct.Empty2, err error)
}

type ProductReviewRepoI interface {
	CreateProductReview(ctx context.Context, req *ct.CreateProductReviewRequest) (resp *ct.ProductReview, err error)
	GetProductReviewByID(ctx context.Context, req *ct.ProductReviewPrimaryKey) (resp *ct.ProductReview, err error)
	GetProductReviewsByProductID(ctx context.Context, req *ct.GetProductReviewsByProductIDRequest) (resp *ct.GetProductReviewsByProductIDResponse, err error)
	GetProductReviewsByCustomerID(ctx context.Context, req *ct.GetProductReviewsByCustomerIDRequest) (resp *ct.GetProductReviewsByCustomerIDResponse, err error)
	UpdateProductReview(ctx context.Context, req *ct.UpdateProductReviewRequest) (resp *ct.ProductReview, err error)
	DeleteProductReview(ctx context.Context, req *ct.ProductReviewPrimaryKey) (resp *ct.Empty4, err error)
}
