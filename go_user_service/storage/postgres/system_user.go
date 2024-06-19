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
	
type system_userRepo struct {
	db *pgxpool.Pool
}

func NewSystemUserRepo(db *pgxpool.Pool) storage.SystemUserRepoI {
	return &system_userRepo{
		db: db,
	}
}

func (c *system_userRepo) Create(ctx context.Context, req *ct.CreateSystemUser) (*ct.SystemUserPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.SystemUserPrimaryKey{Id: id}

	query := `INSERT INTO system_user (
			phone,
			name,
			gmail,
			id,
			role,
			created_at
			) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				NOW()
			)`
	_, err := c.db.Exec(ctx, query, req.Phone, req.Name, req.Gmail, id,req.Role)
	if err != nil {
		log.Println("error while creating system_user")
		return nil, err
	}

	return resp, err
}


func (c *system_userRepo) GetByID(ctx context.Context, req *ct.SystemUserPrimaryKey) (*ct.SystemUser, error) {
    resp := &ct.SystemUser{}

    query := `SELECT phone,
				name,
				gmail,
				id,
				role,
                created_at,
                updated_at
            FROM system_user
            WHERE id=$1 AND deleted_at IS null`

    row := c.db.QueryRow(ctx, query, req.Id)

    var createdAt, updatedAt sql.NullTime
    err := row.Scan(
        &resp.Phone,
        &resp.Name,
        &resp.Gmail,
        &resp.Id,
        &resp.Role,
        &createdAt,
        &updatedAt,
    )
    if err != nil {
		return nil, err
	}

    resp.CreatedAt = helper.NullTimeStampToString(createdAt)
    resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

    return resp, nil
}

func (c *system_userRepo) Update(ctx context.Context, req *ct.UpdateSystemUserRequest) (*ct.UpdateSystemUserResponse, error) {
	resp := &ct.UpdateSystemUserResponse{Message: "System User updated successfully"}
	query := `UPDATE system_user SET  phone=$1,
								 name=$2,
								 gmail=$3,
								 role=$4,
								 updated_at=NOW()
								 WHERE id=$5 AND deleted_at is null`
	_, err := c.db.Exec(ctx, query, req.Phone, req.Name, req.Gmail, req.Role,req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (c *system_userRepo) Delete(ctx context.Context, req *ct.SystemUserPrimaryKey) (*ct.SystemUserEmpty, error) {
	resp:=&ct.SystemUserEmpty{}
	query := `UPDATE system_user SET
							 deleted_at=NOW()
							 WHERE id=$1 AND deleted_at is null RETURNING created_at`

	var createdAt sql.NullTime
	err := c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	if err=helper.DeleteChecker(createdAt);err!=nil {
		return nil,err
	}

	return resp, nil
}

func (c *system_userRepo) GetList(ctx context.Context,req *ct.GetListSystemUserRequest) (*ct.GetListSystemUserResponse,error) {
	resp := &ct.GetListSystemUserResponse{}

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
				role,
				created_at,
				updated_at
			FROM system_user
			WHERE deleted_at is null AND TRUE ` + filter + `
			OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query,offset,req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var createdAt, updatedAt sql.NullTime
		system_user :=&ct.SystemUser{}
		if err := rows.Scan(
			&system_user.Id,
			&system_user.Phone,
			&system_user.Gmail,
			&system_user.Name,
			&system_user.Role,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		system_user.CreatedAt = helper.NullTimeStampToString(createdAt)
		system_user.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.SystemUser = append(resp.SystemUser, system_user)
	}

	queryCount := `SELECT COUNT(*) FROM system_user WHERE deleted_at is null AND TRUE ` + filter +``
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *system_userRepo) GetByGmail(ctx context.Context,req *ct.SystemUserGmail) (*ct.SystemUserPrimaryKey,error) {
	query:=`SELECT id FROM system_user WHERE gmail=$1`
	var id string
	err:=c.db.QueryRow(ctx,query,req.Gmail).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &ct.SystemUserPrimaryKey{Id: id},nil	
}