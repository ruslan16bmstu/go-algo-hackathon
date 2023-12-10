package rpc_server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

import pb "pocket-trader-backend/pb/gen"
import db "pocket-trader-backend/db"

type RPCServer struct {
	pb.UnimplementedTraderServer
	DB db.Api // interface for database
}

func (s RPCServer) GetGlobalRating(ctx context.Context, req *pb.GlobalRatingRequest) (*pb.GlobalRating, error) {
	rating := s.DB.GetRating(int(req.GetSkip()), int(req.GetLimit()))
	return &pb.GlobalRating{
		Datetime: time.Now().Format(time.RFC3339),
		Stocks:   rating,
	}, nil
}

func (s RPCServer) GetStock(ctx context.Context, req *pb.StockRequest) (*pb.Stock, error) {
	stock := s.DB.GetStock(req.GetSecId())
	if stock != nil {
		return stock, nil
	}
	return nil, status.Errorf(codes.NotFound, "no such stock found")
}

func (s RPCServer) GetIndustries(context.Context, *pb.IndustriesRequest) (*pb.Industries, error) {
	return s.DB.GetIndustries(), nil
}
