package db

import pb "pocket-trader-backend/pb/gen"

type Api interface {
	GetIndustries() *pb.Industries
	//GetStock(stockId string)
}
