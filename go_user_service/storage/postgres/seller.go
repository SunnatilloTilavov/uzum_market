package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "user_service/genproto/user_service"
	"user_service/pkg/helper"
	"user_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type sellerRepo struct {
	db *pgxpool.Pool
}

func NewSellerRepo(db *pgxpool.Pool) storage.SellerRepoI {
	return &sellerRepo{
		db: db,
	}
}

func (c *sellerRepo) Create(ctx context.Context, req *ct.CreateSeller) (*ct.SellerPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.SellerPrimaryKey{Id: id}

	query := `INSERT INTO seller (
			phone,
			gmail,
			name,
			date_of_birth,
			shop_id,
			id,
			created_at) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				NOW()
			)`
	_, err := c.db.Exec(ctx, query, req.Phone, req.Gmail, req.Name, &req.DateOfBirth, req.ShopId, id)
	if err != nil {
		log.Println("error while creating seller")
		return nil, err
	}

	return resp, err
}


func (c *sellerRepo) GetByID(ctx context.Context, req *ct.SellerPrimaryKey) (*ct.Seller, error) {
	resp := &ct.Seller{}
	query := `SELECT id,
				   phone,
				   gmail,
				   name,
				   date_of_birth,
				   shop_id,
				   created_at,
				   updated_at
			FROM seller
			WHERE id=$1 AND deleted_at is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, dateOfBirth, createdAt sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&resp.Phone,
		&resp.Gmail,
		&resp.Name,
		&dateOfBirth,
		&resp.ShopId,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.DateOfBirth = helper.NullDateToString(dateOfBirth)
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *sellerRepo) Update(ctx context.Context, req *ct.UpdateSellerRequest) (*ct.UpdateSellerResponse, error) {
	resp := &ct.UpdateSellerResponse{Message: "Seller updated successfully"}
	query := `UPDATE seller SET phone=$1,
								 gmail=$2,
								 name=$3,
								 date_of_birth=$4,
								 shop_id=$5,
								 updated_at=NOW()
								 WHERE id=$6 AND deleted_at is null`
	_, err := c.db.Exec(ctx, query, req.Phone, req.Gmail, req.Name, req.DateOfBirth, req.ShopId, req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (c *sellerRepo) Delete(ctx context.Context, req *ct.SellerPrimaryKey) (*ct.SellerEmpty, error) {
	resp:=&ct.SellerEmpty{}
	query := `UPDATE seller SET
							 deleted_at=NOW()
							 WHERE id=$1 AND deleted_at is null RETURNING created_at`

	var createdAt sql.NullTime
	err := c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	if err=helper.DeleteChecker(createdAt);err!=nil {
		return resp,nil
	}

	return resp, nil
}


func (c *sellerRepo) GetList(ctx context.Context,req *ct.GetListSellerRequest) (*ct.GetListSellerResponse,error) {
	resp := &ct.GetListSellerResponse{}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
				id,
				phone,
				gmail,
				name,
				date_of_birth,
				shop_id,
				created_at,
				updated_at
			FROM seller
        	WHERE deleted_at is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		seller := &ct.Seller{}
		var createdAt, updatedAt, dateOfBirth sql.NullTime
		if err := rows.Scan(
			&seller.Id,
			&seller.Phone,
			&seller.Gmail,
			&seller.Name,
			&dateOfBirth,
			&seller.ShopId,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}
		seller.DateOfBirth = helper.NullDateToString(dateOfBirth)
		seller.CreatedAt = helper.NullTimeStampToString(createdAt)
		seller.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Seller = append(resp.Seller, seller)
	}

	queryCount := `SELECT COUNT(*) FROM seller WHERE deleted_at is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *sellerRepo) GetByGmail(ctx context.Context,req *ct.SellerGmail) (*ct.SellerPrimaryKey,error) {
	query:=`SELECT id FROM seller WHERE gmail=$1`
	var id string
	err:=c.db.QueryRow(ctx,query,req.Gmail).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &ct.SellerPrimaryKey{Id: id},nil
}