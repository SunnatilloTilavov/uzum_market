package postgres

import (
	auth "auth/genproto/auth_service"
	"auth/pkg/hash"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type systemuserRepo struct {
	db *pgxpool.Pool
}

func NewSystemUser(db *pgxpool.Pool) systemuserRepo {
	return systemuserRepo{
		db: db,
	}
}


func (c *systemuserRepo) SystemUserGmailCheck(ctx context.Context,req *auth.SystemUserGmailCheckRequest) (resp *auth.SystemUserGmailCheckResponse,err error) {
	resp=&auth.SystemUserGmailCheckResponse{}
	query:=`SELECT 
			 password
			 	FROM system_user
				WHERE gmail=$1`

	var password string
	err=c.db.QueryRow(ctx,query,req.Gmail).Scan(&password)
	if err != nil {
		return nil, err
	}

	resp.Password=password

	return resp,nil
}

func (c *systemuserRepo) SystemUserCreate(ctx context.Context,req *auth.SystemUserCreateRequest) (resp *auth.SystemUserEmpty,err error) {
	resp=&auth.SystemUserEmpty{}
	query:=`INSERT INTO system_user (password, gmail, created_at)
						VALUES ($1, $2, NOW())
	`
	_,err=c.db.Exec(ctx,query,req.Password,req.Gmail)
	if err != nil {
		return resp, err
	}

	return resp,nil
}


func (c *systemuserRepo) SystemUserUpdatePassword(ctx context.Context, req * auth.SystemUserCreateRequest) (resp *auth.SystemUserEmpty,err error) {
	resp=&auth.SystemUserEmpty{}
	hashedPassword,err:=hash.HashPassword(req.Password)
	if err != nil {
		return resp, err
	}
	resp=&auth.SystemUserEmpty{}
	query:=`UPDATE system_user SET
			password=$1,updated_at=NOW()
			WHERE gmail=$2`
	_,err=c.db.Exec(ctx,query,hashedPassword,req.Gmail)
	if err != nil {
		return resp, err
	}

	return resp,err
}