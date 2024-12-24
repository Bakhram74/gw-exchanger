package service

import (
	"context"
	"log/slog"

	"github.com/Bakhram74/gw-exchanger/internal/models"
	"github.com/Bakhram74/gw-exchanger/pkg"
)

type Rate interface {
	Rates(ctx context.Context) (models.Rates, error)
}

type Service struct {
	rate Rate
}

func NewService(rate Rate) *Service {
	return &Service{
		rate: rate,
	}
}

func (s *Service) GetRates(ctx context.Context) (map[string]float32, error) {
	const op = "Exchange.GetRates"

	log := slog.With(
		slog.String("op", op),
	)
	log.Info("attempting to get rates")

	rates, err := s.rate.Rates(ctx)
	if err != nil {
		log.Error("failed to get rates", pkg.Err(err))
		return nil, err
	}

	respMap := map[string]float32{"USD": rates.Usd, "EUR": rates.Eur}

	return respMap, nil
}
