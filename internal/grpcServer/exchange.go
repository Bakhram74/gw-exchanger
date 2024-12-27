package grpcServer

import (
	"context"
	"strings"

	"github.com/Bakhram74/proto-exchange/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type exchangeApi struct {
	pb.UnimplementedExchangeServiceServer
	exchange Exchange
}

type Exchange interface {
	GetRates(ctx context.Context) (map[string]float32, error)
	GetRateForCurrency(ctx context.Context, fromCurrency, toCurrency string) (float32, error)
}

func (s *exchangeApi) GetExchangeRates(ctx context.Context, empt *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates, err := s.exchange.GetRates(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to fetch exchange rates")
	}

	return &pb.ExchangeRatesResponse{
		Rates: rates,
	}, nil

}

func (s *exchangeApi) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	fromCurrency := strings.ToUpper(req.GetFromCurrency())
	toCurrency := strings.ToUpper(req.GetToCurrency())

	if fromCurrency == toCurrency {
		return nil, status.Error(codes.InvalidArgument, "from_currency and to_currency cannot be the same")
	}
	rate, err := s.exchange.GetRateForCurrency(ctx, fromCurrency, toCurrency)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get exchange rate")
	}

	response := &pb.ExchangeRateResponse{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Rate:         rate,
	}
	return response, nil
}
