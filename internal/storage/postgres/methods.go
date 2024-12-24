package postgres

import (
	"context"

	"github.com/Bakhram74/gw-exchanger/internal/models"
)

func (p *Postgres) Rates(ctx context.Context) (models.Rates, error) {

	query := `SELECT id, usd, eur, created_at FROM rub_rates ORDER BY created_at DESC LIMIT 1`
	row := p.Pool.QueryRow(ctx, query)

	var rates models.Rates

	err := row.Scan(
		&rates.ID,
		&rates.Usd,
		&rates.Eur,
		&rates.CreatedAt,
	)
	if err != nil {
		return models.Rates{}, err
	}
	return rates, nil
}
