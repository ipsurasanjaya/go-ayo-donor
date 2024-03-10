package repository

import (
	"context"
	"database/sql"
	"go-ayo-donor/model/domain"
)

type ProvinceRepository interface {
	Get(ctx context.Context, in domain.GetProvinceIn) ([]domain.GetProvinceOut, error)
}

var _ ProvinceRepository = (*provincesRepo)(nil)

type provincesRepo struct {
	db *sql.DB
}

func NewProvincesRepo(db *sql.DB) *provincesRepo {
	return &provincesRepo{
		db: db,
	}
}
