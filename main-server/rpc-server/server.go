package rpc_server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

import pb "pocket-trader-backend/pb/gen"
import db "pocket-trader-backend/db"

type RPCServer struct {
	pb.UnimplementedTraderServer
	DB db.Api // interface for database
}

func (s RPCServer) GetGlobalRating(context.Context, *pb.GlobalRatingRequest) (*pb.GlobalRating, error) {
	// TODO implement
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalRating not implemented")
}

func (s RPCServer) GetStock(context.Context, *pb.StockRequest) (*pb.Stock, error) {
	// TODO implement
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}

func (s RPCServer) GetIndustries(context.Context, *pb.IndustriesRequest) (*pb.Industries, error) {
	return s.DB.GetIndustries(), nil
}
