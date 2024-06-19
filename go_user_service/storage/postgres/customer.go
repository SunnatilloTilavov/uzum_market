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

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) storage.CustomerRepoI {
	return &customerRepo{
		db: db,
	}
}

func (c *customerRepo) Create(ctx context.Context, req *ct.CreateCustomer) (*ct.CustomerPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.CustomerPrimaryKey{Id: id}

	query := `INSERT INTO customers (
			phone,
			gmail,
			language,
			date_of_birth,
			gender,
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
	_, err := c.db.Exec(ctx, query, req.Phone, req.Gmail, req.Language, &req.DateOfBirth, req.Gender, id)
	if err != nil {
		log.Println("error while creating customer")
		return nil, err
	}

	return resp, err
}

func (c *customerRepo) GetByID(ctx context.Context, req *ct.CustomerPrimaryKey) (*ct.Customer, error) {
	resp := &ct.Customer{}
	query := `SELECT id,
				   phone,
				   gmail,
				   language,
				   date_of_birth,
				   gender,
				   created_at,
				   updated_at
			FROM customers
			WHERE id=$1 AND deleted_at is null`

	row := c.db.QueryRow(ctx, query, req.Id)

	var updatedAt, dateOfBirth, createdAt sql.NullTime
	if err := row.Scan(
		&resp.Id,
		&resp.Phone,
		&resp.Gmail,
		&resp.Language,
		&dateOfBirth,
		&resp.Gender,
		&createdAt,
		&updatedAt); err != nil {
		return nil, err
	}

	resp.DateOfBirth = helper.NullDateToString(dateOfBirth)
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *customerRepo) GetList(ctx context.Context, req *ct.GetListCustomerRequest) (*ct.GetListCustomerResponse, error) {
	resp := &ct.GetListCustomerResponse{}

	filter := ""
    offset := (req.Offset - 1) * req.Limit

    if req.Search != "" {
        filter = ` AND gender ILIKE '%` + req.Search + `%' `
    }

	query := `SELECT 
				id,
				phone,
				gmail,
				language,
				date_of_birth,
				gender,
				created_at,
				updated_at
			FROM customers
        	WHERE deleted_at is null AND TRUE ` + filter + `
           OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		customer := &ct.Customer{}
		var createdAt, updatedAt, dateOfBirth sql.NullTime
		if err := rows.Scan(
			&customer.Id,
			&customer.Phone,
			&customer.Gmail,
			&customer.Language,
			&dateOfBirth,
			&customer.Gender,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}
		customer.DateOfBirth = helper.NullDateToString(dateOfBirth)
		customer.CreatedAt = helper.NullTimeStampToString(createdAt)
		customer.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Customer = append(resp.Customer, customer)
	}

	queryCount := `SELECT COUNT(*) FROM customers WHERE deleted_at is null AND TRUE ` + filter
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *customerRepo) Update(ctx context.Context, req *ct.UpdateCustomerRequest) (*ct.UpdateCustomerResponse, error) {
	resp := &ct.UpdateCustomerResponse{Message: "Customer updated successfully"}
	query := `UPDATE customers SET phone=$1,
								 gmail=$2,
								 language=$3,
								 date_of_birth=$4,
								 gender=$5,
								 updated_at=NOW()
								 WHERE id=$6 AND deleted_at is null`
	_, err := c.db.Exec(ctx, query, req.Phone, req.Gmail, req.Language, req.DateOfBirth, req.Gender, req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *customerRepo) Delete(ctx context.Context, req *ct.CustomerPrimaryKey) (*ct.Empty, error) {
	resp:=&ct.Empty{}
	query := `UPDATE customers SET
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

func (c *customerRepo) GetByGmail(ctx context.Context,req *ct.CustomerGmail) (*ct.CustomerPrimaryKey,error) {
	query:=`SELECT id FROM customers WHERE gmail=$1`
	var id string
	err:=c.db.QueryRow(ctx,query,req.Gmail).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &ct.CustomerPrimaryKey{Id: id},nil	
}