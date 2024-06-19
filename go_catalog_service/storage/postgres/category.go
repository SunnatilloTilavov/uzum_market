package postgres

import (
	"context"
	"database/sql"
	"fmt"

	// "fmt"
	"log"
	ct "microservice/genproto/catalog_service"
	"microservice/storage"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v4/pgxpool"
	// "errors"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) storage.CategoryRepoI {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) Create(ctx context.Context, req *ct.CreateCategory) (resp *ct.Category, err error) {

	resp = &ct.Category{}
	enText := slug.MakeLang(req.NameEn, "en")
	id := uuid.NewString()

	if req.ParentId == "" {
		req.ParentId = id
	}

	_, err = c.db.Exec(ctx, `
		INSERT INTO category (
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			active,
			order_no,
			parent_id
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8
		) `, id, enText, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, req.ParentId)

	if err != nil {
		log.Println("error while creating category")
		return nil, err
	}

	// Obyektni to'ldirish
	resp.Id = id
	resp.Slug = enText
	resp.NameUz = req.NameUz
	resp.NameRu = req.NameRu
	resp.NameEn = req.NameEn
	resp.Active = req.Active
	resp.OrderNo = req.OrderNo
	resp.ParentId = req.ParentId

	return resp, nil
}

func (c *categoryRepo) GetByID(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.GetCategory, err error) {

	resp = &ct.GetCategory{}

	var ParentId sql.NullString

	err = c.db.QueryRow(ctx, `
		SELECT
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			active,
			order_no,
			parent_id
		FROM category
		WHERE id = $1 
	`, req.Id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.Active, &resp.OrderNo, &ParentId)

	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}
	resp.ChildCategories = []*ct.GetChildId{}

	Child_Id := []string{}

	Child_Id, err = c.GetByParentID(ctx, resp.Id)
	if err != nil {
		log.Println("error while getting category by id")
		return nil, err
	}

	for _, ChildId := range Child_Id {
		resp.ChildCategories = append(resp.ChildCategories, &ct.GetChildId{Id: ChildId})
	}

	resp.ParentId = ParentId.String

	return resp, nil
}

func (c *categoryRepo) Update(ctx context.Context, req *ct.UpdateCategory) (resp *ct.GetCategory, err error) {

	resp = &ct.GetCategory{}
	enText := slug.MakeLang(req.NameEn, "en")
	_, err = c.db.Exec(ctx, `
		UPDATE category
		SET
			slug = $1,
			name_uz = $2,
			name_ru = $3,
			name_en = $4,
			active = $5,
			order_no = $6,
			parent_id = $7,
			updated_at = now()
		WHERE id = $8
	`, enText, req.NameUz, req.NameRu, req.NameEn, req.Active, req.OrderNo, req.ParentId, req.Id)

	if err != nil {
		log.Println("error while updating category")
		return nil, err
	}

	category, err := c.GetByID(ctx, &ct.CategoryPrimaryKey{Id: req.Id})
	if err != nil {
		log.Println("error while getting category by id after update")
		return nil, err
	}

	return category, nil
}

func (c *categoryRepo) Delete(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.Empty, err error) {

	_, err = c.db.Exec(ctx, `
		UPDATE category
		SET
			deleted_at = now()
		WHERE id = $1
	`, req.Id)

	if err != nil {
		log.Println("error while deleting category")
		return nil, err
	}

	return &ct.Empty{}, nil
}

func (c *categoryRepo) GetByParentID(ctx context.Context, parentID string) ([]string, error) {
	Child_Id := []string{}
	var categoryId string
	rows, err := c.db.Query(ctx, `
        SELECT
            id
        FROM category
        WHERE parent_id = $1
    `, parentID)
	if err != nil {
		log.Println("error while getting categories by parent id:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&categoryId)
		if err != nil {
			log.Println("error while scanning category:", err)
			return nil, err
		}
		Child_Id = append(Child_Id, categoryId)
	}

	if err = rows.Err(); err != nil {
		log.Println("error after iterating over rows:", err)
		return nil, err
	}

	return Child_Id, nil
}

func (c *categoryRepo) GetList(ctx context.Context, req *ct.GetListCategoryRequest) (*ct.GetListCategoryResponse, error) {
	var (
		parentId   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
		deleted_at sql.NullString
		active     sql.NullBool
	)

	offset := (req.Page - 1) * req.Limit

	filter := "  WHERE deleted_at IS NULL" // Filter for undeleted categories

	if req.Search != "" {
		filter += fmt.Sprintf(` AND (slug ILIKE '%%%v%%' OR name_uz ILIKE '%%%v%%' OR name_ru ILIKE '%%%v%%' OR name_en ILIKE '%%%v%%')`, req.Search, req.Search, req.Search, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	query := `
        SELECT
            id,
            slug,
            name_uz,
            name_ru,
            name_en,
            active,
            order_no,
            parent_id,
            created_at,
            updated_at,
            deleted_at
        FROM category
    ` + filter

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	resp := &ct.GetListCategoryResponse{}
	var count int = 0
	for rows.Next() {
		count++
		category := &ct.GetCategory{}

		err := rows.Scan(
			&category.Id, &category.Slug,
			&category.NameUz, &category.NameRu,
			&category.NameEn, &active,
			&category.OrderNo, &parentId,
			&created_at, &updated_at, &deleted_at,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		category.ParentId = parentId.String
		category.UpdatedAt = updated_at.String
		category.CreatedAt = created_at.String
		category.DeletedAt = deleted_at.String
		category.Active = active.Bool

		category.ChildCategories = []*ct.GetChildId{}

		Child_Id := []string{}

		Child_Id, err = c.GetByParentID(ctx, category.Id)
		if err != nil {
			log.Println("error while getting category by id")
			return nil, err
		}

		for _, ChildId := range Child_Id {
			category.ChildCategories = append(category.ChildCategories, &ct.GetChildId{Id: ChildId})
		}

		resp.Categorys = append(resp.Categorys, category)
	}
	resp.Count = int64(count)
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return resp, nil
}

// func (c *categoryRepo) GetCategoryWithProduct(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.GetCategory, err error){
// 	resp = &ct.GetCategory{}

// 	var ParentId sql.NullString

// 	err = c.db.QueryRow(ctx, `
// 		SELECT
// 			id,
// 			slug,
// 			name_uz,
// 			name_ru,
// 			name_en,
// 			active,
// 			order_no,
// 			parent_id
// 		FROM category
// 		WHERE id = $1
// 	`, req.Id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.Active, &resp.OrderNo, &ParentId)

// 	if err != nil {
// 		log.Println("error while getting category by id")
// 		return nil, err
// 	}
// 	resp.ChildCategories = []*ct.GetChildId{}

// 	Child_Id := []string{}

// 	Child_Id, err = c.GetByParentID(ctx, resp.Id)
// 	if err != nil {
// 		log.Println("error while getting category by id")
// 		return nil, err
// 	}
// 	productIDs := []string{}
// 	product:=&ct.Product1{}
// 	for _, ChildId := range Child_Id {
// 		resp.ChildCategories = append(resp.ChildCategories, &ct.GetChildId{Id: ChildId})
// 		productIDs, err = c.GetProductIDByCategoryID(ctx, resp.Id)
// 		for _, ProductIDs := range productIDs {
// 			product, err = c.GetByParentID(ctx,ProductIDs)
// 			if err != nil {
// 				log.Println("error while getting category by id")
// 				return nil, err
// 			}
// 			resp.ChildCategories = append(resp.ChildCategories, &ct.GetChildId{Products: product})

// 		}

// 	}

// 	resp.ParentId = ParentId.String

// 	return resp, nil
// }

func (c *categoryRepo) GetProductIDByCategoryID(ctx context.Context, req string) ([]string, error) {
	productIDs := []string{}

	rows, err := c.db.Query(ctx, `
		SELECT
			product_id
		FROM product_categories
		WHERE category_id = $1
	`, req)

	if err != nil {
		log.Println("error while getting product categories by category id:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productID string
		if err := rows.Scan(&productID); err != nil {
			log.Println("error while scanning product category row:", err)
			return nil, err
		}
		productIDs = append(productIDs, productID)
	}

	return productIDs, nil
}

// func (c *categoryRepo) GetByIDProduct(ctx context.Context, id string) (resp *ct.Product1, err error) {
// 	resp = &ct.Product1{}

// 	err = c.db.QueryRow(ctx, `
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
// 			image,
// 			created_at,
// 			updated_at
// 		FROM product
// 		WHERE id = $1
// 	`, id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.DescriptionUz, &resp.DescriptionRu, &resp.DescriptionEn, &resp.Active, &resp.OrderNo, &resp.InPrice, &resp.OutPrice, &resp.LeftCount, &resp.DiscountPercent, &resp.Image, &resp.CreatedAt, &resp.UpdatedAt)

// 	if err != nil {
// 		log.Println("error while getting product by id")
// 		return nil, err
// 	}

// 	return resp, nil
// }

func (c *categoryRepo) GetCategoryWithProductId(ctx context.Context, req *ct.CategoryPrimaryKey) (resp *ct.GetCategory, err error) {
	resp = &ct.GetCategory{}
	var ParentId sql.NullString
	var created_at sql.NullString
	var updated_at sql.NullString

	// Fetch the category details
	err = c.db.QueryRow(ctx, `
		SELECT
			id,
			slug,
			name_uz,
			name_ru,
			name_en,
			active,
			order_no,
			parent_id,
			created_at,
			updated_at
		FROM category
		WHERE id = $1 
	`, req.Id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.Active, &resp.OrderNo, &ParentId, &created_at, &updated_at)

	if err != nil {
		log.Println("error while getting category by id:", err)
		return nil, err
	}
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String

	resp.ChildCategories = []*ct.GetChildId{}
	resp.ParentId = ParentId.String

	// Get child categories
	childIDs, err := c.GetByParentID(ctx, resp.Id)
	if err != nil {
		log.Println("error while getting child categories by parent id:", err)
		return nil, err
	}

	// Fetch products for each child category
	for _, childID := range childIDs {
		childCategory := &ct.GetChildId{Id: childID}

		productIDs, err := c.GetProductIDByCategoryID(ctx, childID)
		if err != nil {
			log.Println("error while getting product ids by category id:", err)
			return nil, err
		}

		for _, productID := range productIDs {
			product, err := c.GetByIDProduct(ctx, productID)
			if err != nil {
				log.Println("error while getting product by id:", err)
				return nil, err
			}
			childCategory.Products = append(childCategory.Products, product)
		}

		resp.ChildCategories = append(resp.ChildCategories, childCategory)
	}

	return resp, nil
}

func (c *categoryRepo) GetByIDProduct(ctx context.Context, id string) (*ct.Product1, error) {
	resp := &ct.Product1{}
	var discount_percent sql.NullFloat64
	var created_at sql.NullString
	var updated_at sql.NullString

	err := c.db.QueryRow(ctx, `
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
		WHERE id = $1
	`, id).Scan(&resp.Id, &resp.Slug, &resp.NameUz, &resp.NameRu, &resp.NameEn, &resp.DescriptionUz, &resp.DescriptionRu, &resp.DescriptionEn, &resp.Active, &resp.OrderNo, &resp.InPrice, &resp.OutPrice, &resp.LeftCount, &discount_percent, &created_at, &updated_at)

	resp.DiscountPercent = float32(discount_percent.Float64)
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String

	if err != nil {
		log.Println("error while getting product by id:", err)
		return nil, err
	}

	return resp, nil
}
