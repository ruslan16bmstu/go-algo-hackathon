package rpc_server

import (
	"context"
)

import pb "pocket-trader-backend/pb/gen"
import db "pocket-trader-backend/db"

type RPCServer struct {
	pb.UnimplementedTraderServer
	DB db.Api // interface for database
}

func (s RPCServer) GetGlobalRating(context.Context, *pb.GlobalRatingRequest) (*pb.GlobalRatingResponse, error) {
	// TODO implement
	return nil, nil
}

func (s RPCServer) GetStock(context.Context, *pb.StockRequest) (*pb.StockResponse, error) {
	// TODO implement
	return nil, nil
}
