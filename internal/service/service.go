package service

import (
	"context"
	"log/slog"

	"github.com/Bakhram74/gw-exchanger/internal/models"
	"github.com/Bakhram74/gw-exchanger/pkg"
)

type Rate interface {
	Rates(ctx context.Context) (models.RubRate, error)
	RateForCurrency(ctx context.Context, rateTable, rateColumn string) (float32, error)
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
	const op = "Rate.GetRates"

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

func (s *Service) GetRateForCurrency(ctx context.Context, fromCurrency, toCurrency string) (float32, error) {
	const op = "Rate.GetRateForCurrency"

	log := slog.With(
		slog.String("op", op),
	)
	log.Info("attempting to get rate for currency")

	var rateTable, rateColumn string

	switch fromCurrency {
	case "USD":
		rateTable = "usd_rates"
		if toCurrency == "EUR" {
			rateColumn = "eur"
		} else if toCurrency == "RUB" {
			rateColumn = "rub"
		}
	case "EUR":
		rateTable = "eur_rates"
		if toCurrency == "USD" {
			rateColumn = "usd"
		} else if toCurrency == "RUB" {
			rateColumn = "rub"
		}
	case "RUB":
		rateTable = "rub_rates"
		if toCurrency == "USD" {
			rateColumn = "usd"
		} else if toCurrency == "EUR" {
			rateColumn = "eur"
		}
	}

	rate, err := s.rate.RateForCurrency(ctx, rateTable, rateColumn)
	if err != nil {
		log.Error("error fetching exchange rate: ", pkg.Err(err))
		return 0, err
	}
	return rate, nil
}
