package api

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

const (
	url      = "http://172.16.88.35:18888"
	username = "admin"
	password = "admin"
)

func TestDoLogin(t *testing.T) {
	params := DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}
	err := DoLogin(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
}

func TestHomeCount(t *testing.T) {
	TestDoLogin(t)
	params := HomeCountParam{
		ServerUrl: url,
	}
	overview, err := HomeCount(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	bytes, err := json.Marshal(overview)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}


func TestHomeStatis(t *testing.T) {
	TestDoLogin(t)
	params := HomeStatsParam{
		ServerUrl: url,
	}
	overview, err := HomeStatis(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	bytes, err := json.Marshal(overview)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}