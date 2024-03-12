package repository

import (
	"context"
	"go-ayo-donor/model/domain"
)

func (r *transfusionUnitRepo) GetByProvinceID(
	ctx context.Context,
	provinceID int64,
) ([]domain.GetTransfusionUnitByProvinceIDOut, error) {
	query := `SELECT id, name, code FROM transfusion_units WHERE province_id = $1`

	rows, err := r.db.QueryContext(ctx, query, provinceID)
	if err != nil {
		return nil, err
	}

	out := []domain.GetTransfusionUnitByProvinceIDOut{}
	for rows.Next() {
		tu := domain.GetTransfusionUnitByProvinceIDOut{}
		if err := rows.Scan(&tu.ID, &tu.Name, &tu.Code); err != nil {
			return nil, err
		}
		out = append(out, tu)
	}

	return out, nil
}
