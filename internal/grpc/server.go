package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/Bakhram74/proto-exchange/pb"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
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

func Register(gRPCServer *grpc.Server, exchange Exchange) {
	pb.RegisterExchangeServiceServer(gRPCServer, &exchangeApi{exchange: exchange})

}

type Server struct {
	gRPCServer *grpc.Server
	port       string
}

func New(exchange Exchange, port string) *Server {

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			slog.Error("Recovered from panic", slog.Any("panic", p))
			return status.Error(codes.Internal, "internal error")
		}),
	}

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived,
			logging.PayloadSent,
		),
	}
	log := &slog.Logger{}
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
	))
	Register(gRPCServer, exchange)

	return &Server{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(level), msg, fields...)
	})
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "grpcapp.Run"
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	slog.Info("grpc server started", slog.String("addr", l.Addr().String()))
	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Server) Stop() {
	const op = "grpcapp.Stop"

	slog.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.String("port", s.port))

	s.gRPCServer.Stop()
}
