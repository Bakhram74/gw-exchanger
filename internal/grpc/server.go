package grpc

import (
	"context"

	"github.com/Bakhram74/proto-exchange/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExchangeServiceServer
	exchange Exchange
}

type Exchange interface {
	GetRates(ctx context.Context) (map[string]float32, error)
}

func Register(gRPCServer *grpc.Server, exchange Exchange) {
	pb.RegisterExchangeServiceServer(gRPCServer, &server{exchange: exchange})

}
