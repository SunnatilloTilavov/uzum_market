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

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepoI {
	return &branchRepo{
		db: db,
	}
}

func (c *branchRepo) Create(ctx context.Context, req *ct.CreateBranch) (*ct.BranchPrimaryKey, error) {
	id := uuid.NewString()
	resp := &ct.BranchPrimaryKey{Id: id}

	open := helper.TimeToSecond(req.OpenTime)
	close := helper.TimeToSecond(req.CloseTime)

	query := `INSERT INTO branch (
				phone,
				name,
				addres,
				id,
				open_time,
				close_time,
				active,
				location,
				created_at
			) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				ST_SetSRID(ST_MakePoint($8, $9), 4326),
				NOW()
			);
`
	_, err := c.db.Exec(ctx, query, req.Phone, req.Name, req.Address, id,
		open, close, req.Active, req.Location.Longitude, req.Location.Latitude)
	if err != nil {
		log.Println("error while creating branch")
		return nil, err
	}

	return resp, err
}

func (c *branchRepo) GetByID(ctx context.Context, req *ct.BranchPrimaryKey) (*ct.Branch, error) {
	resp := &ct.Branch{}

	query := `SELECT phone,
                name,
                addres,
                id,
                open_time,
                close_time,
                active,
				ST_Y(location) AS latitude, 
      			ST_X(location) AS longitude,
                created_at,
                updated_at
            FROM branch
            WHERE id=$1 AND deleted_at IS NULL`

	row := c.db.QueryRow(ctx, query, req.Id)

	var (createdAt, updatedAt sql.NullTime
		longitude, latitude      float64)
	var open, close int
	err := row.Scan(
		&resp.Phone,
		&resp.Name,
		&resp.Address,
		&resp.Id,
		&open,
		&close,
		&resp.Active,
		&latitude,
		&longitude,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	resp.Location=&ct.Location{Longitude: longitude,Latitude: latitude}
	resp.OpenTime = helper.SecondToTime(open)
	resp.CloseTime = helper.SecondToTime(close)
	resp.CreatedAt = helper.NullTimeStampToString(createdAt)
	resp.UpdatedAt = helper.NullTimeStampToString(updatedAt)

	return resp, nil
}

func (c *branchRepo) Update(ctx context.Context, req *ct.UpdateBranchRequest) (*ct.UpdateBranchResponse, error) {
	resp := &ct.UpdateBranchResponse{Message: "Branch updated successfully"}
	open := helper.TimeToSecond(req.OpenTime)
	close := helper.TimeToSecond(req.CloseTime)

	query := `UPDATE branch SET  phone=$1,
								 name=$2,
								 addres=$3,
								 open_time=$4,
								 close_time=$5,
								 active=$6,
								 location = ST_SetSRID(ST_MakePoint($7, $8), 4326),
								 updated_at=NOW()
								 WHERE id=$9 AND deleted_at is null`
	_, err := c.db.Exec(ctx, query, req.Phone, req.Name, req.Address, open, close, req.Active, 
		req.Location.Longitude,req.Location.Latitude, req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *branchRepo) Delete(ctx context.Context, req *ct.BranchPrimaryKey) (*ct.BranchEmpty, error) {
	resp:=&ct.BranchEmpty{}
	query := `UPDATE branch SET
							 deleted_at=NOW()
							 WHERE id=$1 AND deleted_at is null RETURNING created_at`

	var createdAt sql.NullTime
	err := c.db.QueryRow(ctx, query, req.Id).Scan(&createdAt)
	if err != nil {
		return nil, err
	}

	if err = helper.DeleteChecker(createdAt); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *branchRepo) GetList(ctx context.Context, req *ct.GetListBranchRequest) (*ct.GetListBranchResponse, error) {
	resp := &ct.GetListBranchResponse{}

	filter := ""
	offset := (req.Offset - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT 
				phone,
				name,
				addres,
				id,
				open_time,
				close_time,
				active,
				ST_Y(location) AS latitude, 
      			ST_X(location) AS longitude,
				created_at,
				updated_at
			FROM branch
			WHERE deleted_at is null AND TRUE ` + filter + `
			OFFSET $1 LIMIT $2
    `

	rows, err := c.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		branch := &ct.Branch{}
		var (
			createdAt, updatedAt sql.NullTime
			open, close          int
			longitude, latitude      float64
			
		)
		if err := rows.Scan(
			&branch.Phone,
			&branch.Name,
			&branch.Address,
			&branch.Id,
			&open,
			&close,
			&branch.Active,
			&latitude,
			&longitude,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		branch.Location=&ct.Location{Longitude: longitude,Latitude: latitude}
		branch.OpenTime = helper.SecondToTime(open)
		branch.CloseTime = helper.SecondToTime(close)
		branch.CreatedAt = helper.NullTimeStampToString(createdAt)
		branch.UpdatedAt = helper.NullTimeStampToString(updatedAt)
		resp.Branches = append(resp.Branches, branch)
	}

	queryCount := `SELECT COUNT(*) FROM branch WHERE deleted_at is null AND TRUE ` + filter + ``
	err = c.db.QueryRow(ctx, queryCount).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
