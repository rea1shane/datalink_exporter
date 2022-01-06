package api

import (
	"context"
	"fmt"
	"testing"
)

func TestDoLogin(t *testing.T) {
	params := DoLoginParam{
		ServerUrl:  "http://172.16.88.35:18888",
		LoginEmail: "admin",
		Password:   "admin",
	}
	err := DoLogin(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
}
