package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	ct "microservice/genproto/catalog_service"
	"microservice/storage"
	"github.com/gosimple/slug"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.ProductRepoI {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(ctx context.Context, req *ct.CreateProduct) (resp *ct.Product, err error) {
	resp = &ct.Product{}

	id := uuid.NewString()
	enText := slug.MakeLang(req.NameEn, "en")

	_, err = p.db.Exec(ctx, `
		INSERT INTO product (
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
			active,
			order_no,
			in_price,
			out_price,
			left_count,
			discount_percent,
			image,
			created_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14,
			$15,
			now()
		) `, id, enText, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz, req.DescriptionRu, req.DescriptionEn, req.Active, req.OrderNo, req.InPrice, req.OutPrice, req.LeftCount, req.DiscountPercent, req.Image)

	if err != nil {
		log.Println("error while creating product")
		return nil, err
	}

	product, err := p.GetByID(ctx, &ct.ProductPrimaryKey{Id: id})
	if err != nil {
		log.Println("error while getting product by id")
		return nil, err
	}

	return product, nil
}

func (p *productRepo) GetByID(ctx context.Context, req *ct.ProductPrimaryKey) (resp *ct.Product, err error) {
	resp = &ct.Product{}
var  UpdatedAt sql.NullString
var CreatedAt sql.NullString
	err =   p.db.QueryRow(ctx, `
		SELECT
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			description_uz,
			description_ru,
			description_en,
			active,
			order_no,
			in_price,
			out_price,
			left_count,
			discount_percent,
			created_at,
			updated_at
		FROM product
		WHERE id = $1 and  deleted_at IS NULL
	`, req.Id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.DescriptionUz, &resp.DescriptionRu, &resp.DescriptionEn, &resp.Active, &resp.OrderNo, &resp.InPrice, &resp.OutPrice, &resp.LeftCount, &resp.DiscountPercent,&CreatedAt, &UpdatedAt)

	resp.UpdatedAt=UpdatedAt.String
	resp.CreatedAt=CreatedAt.String

	if err != nil {
		log.Println("error while getting product by id")
		return nil, err
	}

	return resp, nil
}

func (p *productRepo) Update(ctx context.Context, req *ct.UpdateProduct) (resp *ct.Product, err error) {
	resp = &ct.Product{}

	_, err = p.db.Exec(ctx, `
		UPDATE product
		SET
			slug = $1,
			name_uz = $2,
			name_ru = $3,
			name_en = $4,
			description_uz = $5,
			description_ru = $6,
			description_en = $7,
			active = $8,
			order_no = $9,
			in_price = $10,
			out_price = $11,
			left_count = $12,
			discount_percent = $13,
			image = $14,
			updated_at = now()
		WHERE id = $15 and  deleted_at IS NULL
	`, req.Slug, req.NameUz, req.NameRu, req.NameEn, req.DescriptionUz, req.DescriptionRu, req.DescriptionEn, req.Active, req.OrderNo, req.InPrice, req.OutPrice, req.LeftCount, req.DiscountPercent, req.Image, req.Id)

	if err != nil {
		log.Println("error while updating product")
		return nil, err
	}

	product, err := p.GetByID(ctx, &ct.ProductPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting product by id after update")
		return nil, err
	}

	return product, nil
}

func (p *productRepo) Delete(ctx context.Context, req *ct.ProductPrimaryKey) (resp *ct.Empty, err error) {
	_, err = p.db.Exec(ctx, `
		UPDATE product
		SET
			deleted_at = now()
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting product")
		return nil, err
	}

	return &ct.Empty{}, nil
}

// func (p *productRepo) GetList(ctx context.Context, req *ct.GetListProductRequest) (resp *ct.GetListProductResponse, err error) {
// 	resp = &ct.GetListProductResponse{}
// 	resp.Products = []*ct.Product{}
// 	var UpdatedAt,CreatedAt sql.NullString

// 	offset := (req.Page - 1) * req.Limit

// 	filter := " WHERE deleted_at IS NULL" // Filter for undeleted categories

// 	if req.Search != "" {
// 		filter += fmt.Sprintf(` AND (slug ILIKE '%%%v%%' OR name_uz ILIKE '%%%v%%' OR name_ru ILIKE '%%%v%%' OR name_en ILIKE '%%%v%%')`, req.Search, req.Search, req.Search, req.Search)
// 	}

// 	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)


// 	query := `
// 		SELECT
// 			id,
// 			slug,
// 			name_uz,
// 			name_ru,
// 			name_en,
// 			description_uz,
// 			description_ru,
// 			description_en,
// 			active,
// 			order_no,
// 			in_price,
// 			out_price,
// 			left_count,
// 			discount_percent,
// 			created_at,
// 			updated_at
// 		FROM product
// 		WHERE deleted_at IS NULL
// 	`+filter

// 	rows, err := p.db.Query(ctx, query)
// 	if err != nil {
// 		log.Println("error while getting product list")
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var count int = 0
// 	for rows.Next() {
// 		count++
// 		product := &ct.Product{}
// 		err = rows.Scan(&product.Id, &product.Slug, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn, &product.Active, &product.OrderNo, &product.InPrice, &product.OutPrice, &product.LeftCount, &product.DiscountPercent,  &CreatedAt, &UpdatedAt)
// 		if err != nil {
// 			log.Println("error while scanning product row")
// 			return nil, err
// 		}

// 		product.UpdatedAt=UpdatedAt.String
// 		product.CreatedAt=CreatedAt.String
// 		resp.Products = append(resp.Products, product)
// 	}
// 	resp.Count = int64(count)
// 	err = p.db.QueryRow(ctx, `SELECT COUNT(*) FROM product WHERE deleted_at IS NULL`).Scan(&resp.Count)
// 	if err != nil {
// 		log.Println("error while getting product count")
// 		return nil, err
// 	}

// 	return resp, nil
// }

func (p *productRepo) GetList(ctx context.Context, req *ct.GetListProductRequest) (*ct.GetListProductResponse, error) {
    resp := &ct.GetListProductResponse{}
    resp.Products = []*ct.Product{}
    var UpdatedAt, CreatedAt sql.NullString

    // Calculate OFFSET based on page and limit
    offset := (req.Page - 1) * req.Limit

    // Initialize filter string
    filter := " WHERE deleted_at IS NULL"

    // Add search condition if req.Search is not empty
    if req.Search != "" {
        filter += fmt.Sprintf(` AND (slug ILIKE '%%%v%%' OR name_uz ILIKE '%%%v%%' OR name_ru ILIKE '%%%v%%' OR name_en ILIKE '%%%v%%')`, req.Search, req.Search, req.Search, req.Search)
    }

    // Construct the final query with filter, limit, and offset
    query := fmt.Sprintf(`
        SELECT
            id,
            slug,
            name_uz,
            name_ru,
            name_en,
            description_uz,
            description_ru,
            description_en,
            active,
            order_no,
            in_price,
            out_price,
            left_count,
            discount_percent,
            created_at,
            updated_at
        FROM product
        %s
        ORDER BY created_at DESC
        LIMIT %v OFFSET %v
    `, filter, req.Limit, offset)

    // Execute the query
    rows, err := p.db.Query(ctx, query)
    if err != nil {
        log.Println("error while executing product query:", err)
        return nil, err
    }
    defer rows.Close()

    // Iterate over the rows and populate response
    var count int64 = 0
    for rows.Next() {
        count++
        product := &ct.Product{}
        err := rows.Scan(
            &product.Id,
            &product.Slug,
            &product.NameUz,
            &product.NameRu,
            &product.NameEn,
            &product.DescriptionUz,
            &product.DescriptionRu,
            &product.DescriptionEn,
            &product.Active,
            &product.OrderNo,
            &product.InPrice,
            &product.OutPrice,
            &product.LeftCount,
            &product.DiscountPercent,
            &CreatedAt,
            &UpdatedAt,
        )
        if err != nil {
            log.Println("error while scanning product row:", err)
            return nil, err
        }

        // Assign updated and created timestamps
        product.UpdatedAt = UpdatedAt.String
        product.CreatedAt = CreatedAt.String

        // Append product to response
        resp.Products = append(resp.Products, product)
    }

    // Assign total count of products
    resp.Count = count

    // Query total count of products (without limit and offset) for pagination
    if req.Search == "" {
        err := p.db.QueryRow(ctx, `SELECT COUNT(*) FROM product WHERE deleted_at IS NULL`).Scan(&resp.Count)
        if err != nil {
            log.Println("error while getting product count:", err)
            return nil, err
        }
    }

    return resp, nil
}

