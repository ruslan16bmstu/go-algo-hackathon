package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	pb "pocket-trader-backend/pb/gen"
	"sort"
)

type DB struct {
	industries  map[int]string
	stocks      map[string]int
	predictions []*pb.Stock
}

type jsonStock struct {
	SecId    string `json:"sec_id"`
	Industry string `json:"industry"`
}

type jsonPrediction struct {
	SecId string  `json:"sec_id"`
	Delta float32 `json:"delta"`
	Price float32 `json:"price"`
}

func (local *DB) GetIndustries() *pb.Industries {
	industries := pb.Industries{}

	for key, value := range local.industries {
		industries.Industries = append(industries.Industries, &pb.Industries_Industry{
			Id:   uint32(key),
			Name: value,
		})
	}

	return &industries
}

func (local *DB) GetStock(stockId string) *pb.Stock {
	for _, val := range local.predictions {
		if val.SecId == stockId {
			return val
		}
	}
	return nil
}

func (local *DB) GetRating(skip, limit int) []*pb.Stock {
	begin, end := skip, skip+limit
	if begin > len(local.predictions) {
		return make([]*pb.Stock, 0, 1)
	}
	if end > len(local.predictions) {
		end = len(local.predictions)
	}
	return local.predictions[begin:end]
}

func Init(industriesSource, predictionsSource string) *DB {
	db := DB{}
	db.stocks = make(map[string]int)
	db.industries = make(map[int]string)

	db.parseStockIndustries(industriesSource)
	db.parsePredictions(predictionsSource)

	return &db
}

func (local *DB) parseStockIndustries(industriesSource string) {
	jsonFile, err := os.Open(industriesSource)
	if err != nil {
		log.Fatalf("error opening db file: %v", err)
	}
	defer jsonFile.Close()

	var arr []jsonStock
	bytes, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal([]byte(bytes), &arr)

	industriesSet := make(map[string]any)
	tempMap := make(map[string]int)

	for _, stock := range arr {
		industriesSet[stock.Industry] = nil
	}

	industryId := 1
	for ind := range industriesSet {
		local.industries[industryId] = ind
		tempMap[ind] = industryId
		industryId++
	}

	for _, stock := range arr {
		local.stocks[stock.SecId] = tempMap[stock.Industry]
	}
}

func (local *DB) parsePredictions(predictionsSource string) {
	jsonFile, err := os.Open(predictionsSource)
	if err != nil {
		log.Fatalf("error opening db file: %v", err)
	}
	defer jsonFile.Close()

	var arr []jsonPrediction
	bytes, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(bytes), &arr)
	if err != nil {
		log.Println(err)
	}

	for _, p := range arr {
		local.predictions = append(local.predictions, &pb.Stock{
			Industry:   uint32(local.stocks[p.SecId]),
			SecId:      p.SecId,
			StockPrice: p.Price,
			Delta:      p.Delta,
		})
	}

	sort.Slice(local.predictions, func(i, j int) bool {
		return local.predictions[i].Delta > local.predictions[j].Delta
	})
}
