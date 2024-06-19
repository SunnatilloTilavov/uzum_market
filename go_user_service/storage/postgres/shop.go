package postgres

import (
	"context"
	"database/sql"
	"log"
	ct "user_service/genproto/user_service"
	"user_service/pkg/helper"
	"user_service/storage"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type shopRepo struct {
	db *pgxpool.Pool
}

func NewShopRepo(db *pgxpool.Pool) storage.ShopRepoI {
	return &shopRepo{
		db: db,
	}
}

func (c *shopRepo) Create(ctx context.Context, req *ct.CreateShop) (*ct.ShopPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.ShopPrimaryKey{Id: id}
	slug := slug.Make(req.NameEn)

	query := `INSERT INTO shop (
			slug,
			phone,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
			location,
			currency,
			payment_types,
			id,
			created_at) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8,
				ST_SetSRID(ST_MakePoint($9, $10), 4326),
				$11,
				$12,
				$13,
				NOW()
			)`
	_, err := c.db.Exec(ctx, query, slug, req.Phone, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz,
		req.DescriptionRu, req.DescriptionEn, req.Location.Longitude, req.Location.Latitude, req.Currency, pq.Array(req.PaymentTypes), id)
	if err != nil {
		log.Println("error while creating shop")
		return nil, err
	}

	return resp, nil
}

func (c *shopRepo) GetById(ctx context.Context, req *ct.ShopPrimaryKey) (*ct.GetByID, error) {
	resp := &ct.GetByID{}

	query := `SELECT 
			slug,
			phone,
			name_uz,
			name_ru,
			name_en,	
			description_uz,
			description_ru,
			description_en,
			currency,
			id,
			COALESCE(payment_types, '{}'),
			ST_Y(location) AS latitude, 
      		ST_X(location) AS longitude,
			created_at,
			updated_at
			FROM shop
			WHERE id=$1 AND deleted_at is null`

	row := c.db.QueryRow(ctx, query, req.Id)
	var (createdAt, updatedAt sql.NullTime
		longitude, latitude      float64)
	if err := row.Scan(
		&resp.Slug,
		&resp.Phone,
		&resp.NameUz,
		&resp.NameRu,
		&resp.NameEn,
		&resp.DescriptionUz,
		&resp.DescriptionRu,
		&resp.DescriptionEn,
		&resp.Currency,
		&resp.Id,
		&resp.PaymentTypes,
		&longitude,
		&latitude,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.Location=&ct.LocationShop{Longitude: longitude,Latitude: latitude}
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *shopRepo) Update(ctx context.Context, req *ct.UpdateShopRequest) (*ct.ShopEmpty, error) {
	resp := &ct.ShopEmpty{}
	query := `UPDATE shop SET 	phone=$1,
								name_uz=$2,
								name_ru=$3,
								name_en=$4,	
								description_uz=$5,
								description_ru=$6,
								description_en=$7,
								location=$8,
								currency=$9,
								payment_types=$10,
								location = ST_SetSRID(ST_MakePoint($11, $12), 4326)
								 updated_at=NOW()
								 WHERE id=$13 AND deleted_at is null`
	_, err := c.db.Exec(ctx, query, req.Phone, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz,
		req.DescriptionRu, req.DescriptionEn, req.Location, req.Currency, pq.Array(req.PaymentTypes),req.Location.Longitude,req.Location.Latitude, req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *shopRepo) Delete(ctx context.Context, req *ct.ShopPrimaryKey) (*ct.ShopEmpty, error) {
	resp := &ct.ShopEmpty{}
	query := `UPDATE shop SET
							 deleted_at=NOW()
							 WHERE id=$1 AND deleted_at is null RETURNING created_at`

	var createdAt sql.NullTime
	err := c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	if err = helper.DeleteChecker(createdAt); err != nil {
		return resp, nil
	}

	return resp, nil
}

func (c *shopRepo) GetList(ctx context.Context, req *ct.GetListShopRequest) (*ct.GetListShopResponse, error) {
	resp := &ct.GetListShopResponse{}
	shop := &ct.Shop{}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND description_uz ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
				slug,
				phone,
				name_uz,
				name_ru,
				name_en,	
				description_uz,
				description_ru,
				description_en,
				currency,
				id,
				COALESCE(payment_types, '{}'),
				ST_Y(location) AS latitude, 
      			ST_X(location) AS longitude,
				created_at,
				updated_at
			FROM shop
			WHERE deleted_at is null AND TRUE ` + filter + `
			OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (createdAt, updatedAt sql.NullTime
			longitude, latitude      float64
		)
		if err := rows.Scan(
			&shop.Slug,
			&shop.Phone,
			&shop.NameUz,
			&shop.NameRu,
			&shop.NameEn,
			&shop.DescriptionUz,
			&shop.DescriptionRu,
			&shop.DescriptionEn,
			&shop.Currency,
			&shop.Id,
			&shop.PaymentTypes,
			&longitude,
			&latitude,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		shop.Location=&ct.LocationShop{Longitude: longitude,Latitude: latitude}
		shop.CreatedAt = helper.NullTimeStampToString(createdAt)
		shop.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Shop = append(resp.Shop, shop)
	}

	queryCount := `SELECT COUNT(*) FROM shop WHERE deleted_at is null AND TRUE ` + filter + ``
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
