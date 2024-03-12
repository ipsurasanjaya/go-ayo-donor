package repository

import (
	"context"
	"database/sql"
	"go-ayo-donor/model/domain"
)

type TransfusionUnitRepository interface {
	GetByProvinceID(ctx context.Context, provinceID int64) ([]domain.GetTransfusionUnitByProvinceIDOut, error)
}

type transfusionUnitRepo struct {
	db *sql.DB
}

var _ TransfusionUnitRepository = (*transfusionUnitRepo)(nil)

func NewTransfusionUnitRepo(db *sql.DB) *transfusionUnitRepo {
	return &transfusionUnitRepo{
		db: db,
	}
}
