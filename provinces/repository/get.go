package repository

import (
	"context"
	"go-ayo-donor/model/domain"
	"strings"
)

func (r *provincesRepo) Get(ctx context.Context, in domain.GetProvinceIn) ([]domain.GetProvinceOut, error) {
	query := `SELECT id, name FROM provinces `
	args := []interface{}{}

	if in.Search != "" {
		in.Search = strings.ToLower(in.Search)
		query += `WHERE LOWER(name) LIKE '%' || $1 || '%' `
		args = append(args, in.Search)
	}

	if in.Limit != 0 {
		query += `LIMIT $2`
		args = append(args, in.Limit)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	out := []domain.GetProvinceOut{}
	for rows.Next() {
		p := domain.GetProvinceOut{}
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}

		out = append(out, p)
	}

	return out, nil
}
