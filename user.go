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

func PrintBeautyStruct(value interface{}) string {
	marshalStruct, _ := json.MarshalIndent(value, "", "\n")
	return string(marshalStruct)
}

func (user User_info) print_info() string {
	return PrintBeautyStruct(user)
}
