package grpc

import (
	"context"

	"github.com/Bakhram74/proto-exchange/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetExchangeRates(ctx context.Context, empt *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates, err := s.exchange.GetRates(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to fetch exchange rates")
	}

	return &pb.ExchangeRatesResponse{
		Rates: rates,
	}, nil

}

func (s *server) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	return nil, nil
}
