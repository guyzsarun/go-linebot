package main

import (
	"encoding/json"
)

type User_info struct {
	UserID         string
	DisplayName    string
	Language       string
	ProfilePicture string
	StatusMessage  string
}

type Miner struct {
	WorkerOnline    int    `json:"workersOnline"`
	CurrentHashRate int    `json:"currentHashrate"`
	CurrentLuck     string `json:"currentluck"`
	LastShare       Worker `json:"workers"`
}

type Worker struct {
	WorkerStatus WorkerStatus `json:"0"`
}

type WorkerStatus struct {
	LastBeat  int  `json:"lastBeat"`
	HashRate  int  `json:"hr"`
	Offline   bool `json:"offline`
	HashRate2 int  `json:"hr2`
}

func PrintBeautyStruct(value interface{}) string {
	marshalStruct, _ := json.MarshalIndent(value, "", "")
	return string(marshalStruct)
}

func (user User_info) print_info() string {
	return PrintBeautyStruct(user)
}

func (miner Miner) print_info() string {
	return PrintBeautyStruct(miner)
}
