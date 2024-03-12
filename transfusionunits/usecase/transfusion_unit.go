package usecase

import (
	"context"
	"fmt"
	"go-ayo-donor/model/domain"
	"go-ayo-donor/transfusionunits/repository"
)

type TransfusionUnitUsecase interface {
	GetByProvinceID(ctx context.Context, provinceID int64) ([]domain.GetTransfusionUnitByProvinceIDOut, error)
}

type transfusionUnitUsecase struct {
	repo repository.TransfusionUnitRepository
}

var _ TransfusionUnitUsecase = (*transfusionUnitUsecase)(nil)

func NewUsecase(repo repository.TransfusionUnitRepository) *transfusionUnitUsecase {
	return &transfusionUnitUsecase{
		repo: repo,
	}
}

func (t *transfusionUnitUsecase) GetByProvinceID(
	ctx context.Context,
	provinceID int64,
) ([]domain.GetTransfusionUnitByProvinceIDOut, error) {
	out, err := t.repo.GetByProvinceID(ctx, provinceID)
	if err != nil {
		return nil, err
	}
	fmt.Println(out)

	if len(out) == 0 {
		return nil, domain.ErrDataNotFound
	}

	return out, nil
}
