package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "microservice/genproto/catalog_service"
	"microservice/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productReviewRepo struct {
	db *pgxpool.Pool
}

func NewProductReviewRepo(db *pgxpool.Pool) storage.ProductReviewRepoI {
	return &productReviewRepo{
		db: db,
	}
}

func (p *productReviewRepo) CreateProductReview(ctx context.Context, req *ct.CreateProductReviewRequest) (resp *ct.ProductReview, err error) {
	resp = &ct.ProductReview{
		Id:         uuid.NewString(),
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
		Text:       req.Text,
		Rating:     req.Rating,
		OrderId:    req.OrderId,
	}

	_, err = p.db.Exec(ctx, `
		INSERT INTO product_reviews (
			id,
			customer_id,
			product_id,
			text,
			rating,
			order_id,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, now()
		)`, resp.Id, resp.CustomerId, resp.ProductId, resp.Text, resp.Rating, resp.OrderId)

	if err != nil {
		log.Println("error while creating product review:", err)
		return nil, err
	}

	return resp, nil
}

func (p *productReviewRepo) GetProductReviewByID(ctx context.Context, req *ct.ProductReviewPrimaryKey) (resp *ct.ProductReview, err error) {
	resp = &ct.ProductReview{}
var OrderId,CreatedAt sql.NullString
	err = p.db.QueryRow(ctx, `
		SELECT
			id,
			customer_id,
			product_id,
			text,
			rating,
			order_id,
			created_at
		FROM product_reviews
		WHERE id = $1
	`, req.Id).Scan(&resp.Id, &resp.CustomerId, &resp.ProductId, &resp.Text, &resp.Rating,&OrderId, &CreatedAt)
resp.OrderId=OrderId.String
resp.CreatedAt=CreatedAt.String
	if err != nil {
		log.Println("error while getting product review by id:", err)
		return nil, err
	}

	return resp, nil
}

func (p *productReviewRepo) GetProductReviewsByProductID(ctx context.Context, req *ct.GetProductReviewsByProductIDRequest) (resp *ct.GetProductReviewsByProductIDResponse, err error) {
	resp = &ct.GetProductReviewsByProductIDResponse{}
	var OrderId,CreatedAt sql.NullString
	rows, err := p.db.Query(ctx, `
		SELECT
			id,
			customer_id,
			product_id,
			text,
			rating,
			order_id,
			created_at
		FROM product_reviews
		WHERE product_id = $1
	`, req.ProductId)

	if err != nil {
		log.Println("error while getting product reviews by product id:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		review := &ct.ProductReview{}
		err = rows.Scan(&review.Id, &review.CustomerId, &review.ProductId, &review.Text, &review.Rating, &OrderId, &CreatedAt)
		if err != nil {
			log.Println("error while scanning product review row:", err)
			return nil, err
		}
		review.OrderId=OrderId.String
		review.CreatedAt=CreatedAt.String
		resp.Reviews = append(resp.Reviews, review)
	}

	

	return resp, nil
}

func (p *productReviewRepo) GetProductReviewsByCustomerID(ctx context.Context, req *ct.GetProductReviewsByCustomerIDRequest) (resp *ct.GetProductReviewsByCustomerIDResponse, err error) {
	resp = &ct.GetProductReviewsByCustomerIDResponse{}
	var OrderId,CreatedAt sql.NullString
	rows, err := p.db.Query(ctx, `
		SELECT
			id,
			customer_id,
			product_id,
			text,
			rating,
			order_id,
			created_at
		FROM product_reviews
		WHERE customer_id = $1
	`, req.CustomerId)

	if err != nil {
		log.Println("error while getting product reviews by customer id:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		review := &ct.ProductReview{}
		err = rows.Scan(&review.Id, &review.CustomerId, &review.ProductId, &review.Text, &review.Rating, &OrderId, &CreatedAt)
		if err != nil {
			log.Println("error while scanning product review row:", err)
			return nil, err
		}
		review.OrderId=OrderId.String
		review.CreatedAt=CreatedAt.String
		resp.Reviews = append(resp.Reviews, review)
	}

	return resp, nil
}

func (p *productReviewRepo) UpdateProductReview(ctx context.Context, req *ct.UpdateProductReviewRequest) (resp *ct.ProductReview, err error) {
	resp = &ct.ProductReview{
		Id:     req.Id,
		Text:   req.Text,
		Rating: req.Rating,
	}

	_, err = p.db.Exec(ctx, `
		UPDATE product_reviews
		SET
			text = $1,
			rating = $2,
			updated_at = now()
		WHERE id = $3
	`, req.Text, req.Rating, req.Id)

	if err != nil {
		log.Println("error while updating product review:", err)
		return nil, err
	}

	review, err := p.GetProductReviewByID(ctx, &ct.ProductReviewPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting product review by id after update:", err)
		return nil, err
	}

	return review, nil
}

func (p *productReviewRepo) DeleteProductReview(ctx context.Context, req *ct.ProductReviewPrimaryKey) (resp *ct.Empty4, err error) {
	_, err = p.db.Exec(ctx, `
		DELETE FROM product_reviews
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting product review:", err)
		return nil, err
	}

	return &ct.Empty4{}, nil
}
