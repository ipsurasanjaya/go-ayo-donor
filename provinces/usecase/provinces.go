package usecase

import (
	"context"
	"go-ayo-donor/model/domain"
	"go-ayo-donor/provinces/repository"
)

type ProvinceUsecase interface {
	Get(ctx context.Context, in domain.GetProvinceIn) ([]domain.GetProvinceOut, error)
}

type provinceUsecase struct {
	repo repository.ProvinceRepository
}

var _ ProvinceUsecase = (*provinceUsecase)(nil)

func NewUsecase(repo repository.ProvinceRepository) *provinceUsecase {
	return &provinceUsecase{
		repo: repo,
	}
}

func (p *provinceUsecase) Get(
	ctx context.Context,
	in domain.GetProvinceIn,
) ([]domain.GetProvinceOut, error) {
	out, err := p.repo.Get(ctx, in)
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, domain.ErrDataNotFound
	}

	return out, nil
}
