package postgres

import (
	"context"
	"log"

	ct "microservice/genproto/catalog_service"
	"microservice/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productCategoryRepo struct {
	db *pgxpool.Pool
}

func NewProductCategoryRepo(db *pgxpool.Pool) storage.ProductCategoryRepoI {
	return &productCategoryRepo{
		db: db,
	}
}

func (p *productCategoryRepo) CreateProductCategory(ctx context.Context, req *ct.CreateProductCategoryRequest) (resp *ct.ProductCategory, err error) {
	resp = &ct.ProductCategory{
		Id:         uuid.NewString(),
		ProductId:  req.ProductId,
		CategoryId: req.CategoryId,
	}

	_, err = p.db.Exec(ctx, `
		INSERT INTO product_categories (
			id,
			product_id,
			category_id
		) VALUES (
			$1, $2, $3
		)`, resp.Id, resp.ProductId, resp.CategoryId)

	if err != nil {
		log.Println("error while creating product category:", err)
		return nil, err
	}

	return resp, nil
}

func (p *productCategoryRepo) GetProductCategoriesByProductID(ctx context.Context, req *ct.GetProductCategoriesByProductIDRequest) (resp *ct.GetProductCategoriesByProductIDResponse, err error) {
	resp = &ct.GetProductCategoriesByProductIDResponse{}
	rows, err := p.db.Query(ctx, `
		SELECT
			id,
			product_id,
			category_id
		FROM product_categories
		WHERE product_id = $1
	`, req.ProductId)

	if err != nil {
		log.Println("error while getting product categories by product id:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		category := &ct.ProductCategory{}
		err = rows.Scan(&category.Id, &category.ProductId, &category.CategoryId)
		if err != nil {
			log.Println("error while scanning product category row:", err)
			return nil, err
		}
		resp.Categories = append(resp.Categories, category)
	}

	return resp, nil
}

func (p *productCategoryRepo) GetProductCategoriesByCategoryID(ctx context.Context, req *ct.GetProductCategoriesByCategoryIDRequest) (resp *ct.GetProductCategoriesByCategoryIDResponse, err error) {
	resp = &ct.GetProductCategoriesByCategoryIDResponse{}
	rows, err := p.db.Query(ctx, `
		SELECT
			id,
			product_id,
			category_id
		FROM product_categories
		WHERE category_id = $1
	`, req.CategoryId)

	if err != nil {
		log.Println("error while getting product categories by category id:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		category := &ct.ProductCategory{}
		err = rows.Scan(&category.Id, &category.ProductId, &category.CategoryId)
		if err != nil {
			log.Println("error while scanning product category row:", err)
			return nil, err
		}
		resp.Categories = append(resp.Categories, category)
	}

	return resp, nil
}

func (p *productCategoryRepo) DeleteProductCategory(ctx context.Context, req *ct.DeleteProductCategoryRequest) (resp *ct.Empty2, err error) {
	_, err = p.db.Exec(ctx, `
		DELETE FROM product_categories
		WHERE product_id = $1 
	`, req.ProductId)

	if err != nil {
		log.Println("error while deleting product category:", err)
		return nil, err
	}

	return &ct.Empty2{}, nil
}
