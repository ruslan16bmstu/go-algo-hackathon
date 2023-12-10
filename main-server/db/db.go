package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	pb "pocket-trader-backend/pb/gen"
)

type DB struct {
	industries map[int]string
	stocks     map[string]int
}

type jsonStock struct {
	SecId    string `json:"sec_id"`
	Industry string `json:"industry"`
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

func Init(industriesSource string) *DB {
	db := DB{}
	db.stocks = make(map[string]int)
	db.industries = make(map[int]string)

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
		db.industries[industryId] = ind
		tempMap[ind] = industryId
		industryId++
	}

	for _, stock := range arr {
		db.stocks[stock.SecId] = tempMap[stock.Industry]
	}

	return &db
}
