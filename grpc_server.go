package main

import (
	"context"
	"net"

	microservices "github.com/sushant102004/microservices/proto"
	"google.golang.org/grpc"
)

func makeGRPCServer(listenAddr string, svc PriceFetcher) error {
	grpcPriceFetcher := NewGRPCPriceFetcher(svc)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	server := grpc.NewServer(&grpc.EmptyServerOption{})
	microservices.RegisterPriceFetcherServer(server, grpcPriceFetcher)

	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}

type GRPCPriceFetcher struct {
	svc PriceFetcher
	microservices.UnimplementedPriceFetcherServer
}

func NewGRPCPriceFetcher(svc PriceFetcher) *GRPCPriceFetcher {
	return &GRPCPriceFetcher{
		svc: svc,
	}
}

func (s *GRPCPriceFetcher) FetchPrice(ctx context.Context, req *microservices.PriceRequest) (*microservices.PriceResponse, error) {
	price, err := s.svc.FetchPrice(ctx, req.Crypto)
	if err != nil {
		return nil, err
	}

	res := &microservices.PriceResponse{
		Price:  float32(price),
		Crypto: req.Crypto,
	}

	return res, nil
}
