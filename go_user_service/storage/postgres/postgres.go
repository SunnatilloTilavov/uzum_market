package postgres

import (
	"context"
	"fmt"
	"log"
	"user_service/config"
	"user_service/storage"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	customer storage.CustomerRepoI
	shop storage.ShopRepoI
	seller storage.SellerRepoI
	branch storage.BranchRepoI
	systemuser storage.SystemUserRepoI
}


func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Customer() storage.CustomerRepoI {
	if s.customer == nil {
		s.customer = NewCustomerRepo(s.db)
	}

	return s.customer
}


func (s *Store) Shop() storage.ShopRepoI {
	if s.shop == nil {
		s.shop = NewShopRepo(s.db)
	}

	return s.shop
}


func (s *Store) Seller() storage.SellerRepoI {
	if s.seller == nil {
		s.seller = NewSellerRepo(s.db)
	}

	return s.seller
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}

func (s *Store) SystemUser() storage.SystemUserRepoI {
	if s.systemuser == nil {
		s.systemuser = NewSystemUserRepo(s.db)
	}

	return s.systemuser
}