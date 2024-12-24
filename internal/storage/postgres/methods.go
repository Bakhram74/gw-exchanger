package postgres

import (
	"context"
	"fmt"

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

func (p *Postgres) RateForCurrency(ctx context.Context, rateTable, rateColumn string) (float32, error) {

	query := fmt.Sprintf(`SELECT %s FROM %s ORDER BY created_at DESC LIMIT 1`, rateColumn, rateTable)

	var rate float32
	err := p.Pool.QueryRow(ctx, query).Scan(&rate)
	
	if err != nil {
		return 0,  err
	}

	return rate, nil
}
