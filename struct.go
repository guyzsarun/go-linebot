package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User_info struct {
	UserID         string
	DisplayName    string
	Language       string
	ProfilePicture string
	StatusMessage  string
}

type Miner struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	LastSeen         int64   `json:"lastSeen"`
	HashRate         float32 `json:"currentHashrate"`
	ReportedHashRate float32 `json:"reportedHashrate"`
	ValidShare       int     `json:"validShares"`
	StaleShares      int     `json:"staleShares"`
}

type ReturnStruct struct {
	MinerStatus      string
	Lastseen         string
	CurrentHashrate  string
	ReportedHashrate string
}

func PrintBeautyStruct(value interface{}) string {
	marshalStruct, _ := json.MarshalIndent(value, "", "")
	return string(marshalStruct)
}

func (user User_info) print_info() string {
	return PrintBeautyStruct(user)
}

func (returnStruct ReturnStruct) print_info() string {
	return PrintBeautyStruct(returnStruct)
}

func (miner *Miner) postProcess() ReturnStruct {
	returnStruct := ReturnStruct{}
	returnStruct.Lastseen = fmt.Sprintf("%s", time.Unix(miner.Data.LastSeen, 0))
	returnStruct.CurrentHashrate = fmt.Sprintf("%.2f MH/s", miner.Data.HashRate/1000000)
	returnStruct.ReportedHashrate = fmt.Sprintf("%.2f MH/s", miner.Data.ReportedHashRate/1000000)
	returnStruct.MinerStatus = miner.Status
	return returnStruct
}
