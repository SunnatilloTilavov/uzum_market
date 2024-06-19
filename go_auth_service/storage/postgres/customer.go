package postgres

import (
	auth "auth/genproto/auth_service"
	"auth/pkg/hash"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomer(db *pgxpool.Pool) customerRepo {
	return customerRepo{
		db: db,
	}
}


func (c *customerRepo) GmailCheck(ctx context.Context,req *auth.GmailCheckRequest) (*auth.GmailCheckResponse,error) {
	resp:=&auth.GmailCheckResponse{}
	query:=`SELECT 
			 password
			 	FROM customer
				WHERE gmail=$1 `

	var password string
	err:=c.db.QueryRow(ctx,query,req.Gmail).Scan(&password)
	if err != nil {
		return nil, err
	}

	resp.Password=password

	return resp,nil
}

func (c *customerRepo) Create(ctx context.Context,req *auth.CreateRequest) (*auth.Empty,error) {
	resp:=&auth.Empty{}
	query:=`INSERT INTO customer (password, gmail, created_at)
						VALUES ($1, $2, NOW())
	`
	_,err:=c.db.Exec(ctx,query,req.Password,req.Gmail)
	if err != nil {
		return resp, err
	}

	return resp,nil
}


func (c *customerRepo) UpdatePassword(ctx context.Context, req * auth.CreateRequest) (*auth.Empty,error) {
	resp:=&auth.Empty{}
	hashedPassword,err:=hash.HashPassword(req.Password)
	if err != nil {
		return resp, err
	}
	resp=&auth.Empty{}
	query:=`UPDATE customer SET
			password=$1,updated_at=NOW()
			WHERE gmail=$2`
	_,err=c.db.Exec(ctx,query,hashedPassword,req.Gmail)
	if err != nil {
		return resp, err
	}

	return resp,err
}