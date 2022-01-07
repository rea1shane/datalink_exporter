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
	sessionCookie, err := DoLogin(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Println(sessionCookie.Name + ": " + sessionCookie.Value)
}

func TestHomeCount(t *testing.T) {
	sessionCookie, err := DoLogin(DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	params := HomeCountParam{
		ServerUrl:     url,
		SessionCookie: sessionCookie,
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
	sessionCookie, err := DoLogin(DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	params := HomeStatisParam{
		ServerUrl:     url,
		SessionCookie: sessionCookie,
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

func TestGroupInitGroup(t *testing.T) {
	sessionCookie, err := DoLogin(DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	params := GroupInitGroupParam{
		ServerUrl:     url,
		Start:         0,
		Length:        10,
		SessionCookie: sessionCookie,
	}
	overview, err := GroupInitGroup(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	bytes, err := json.Marshal(overview)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestWorkerInitWorker(t *testing.T) {
	sessionCookie, err := DoLogin(DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	params := WorkerInitWorkerParam{
		ServerUrl:     url,
		Start:         0,
		Length:        10,
		SessionCookie: sessionCookie,
	}
	overview, err := WorkerInitWorker(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	bytes, err := json.Marshal(overview)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMysqlTaskMysqlTaskDatas(t *testing.T) {
	sessionCookie, err := DoLogin(DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	params := MysqlTaskMysqlTaskDatasParam{
		ServerUrl:     url,
		Start:         0,
		Length:        10,
		SessionCookie: sessionCookie,
	}
	overview, err := MysqlTaskMysqlTaskDatas(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	bytes, err := json.Marshal(overview)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestHbaseTaskInitHbaseTaskList(t *testing.T) {
	sessionCookie, err := DoLogin(DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	params := HbaseTaskInitHbaseTaskListParam{
		ServerUrl:     url,
		Start:         0,
		Length:        10,
		SessionCookie: sessionCookie,
	}
	overview, err := HbaseTaskInitHbaseTaskList(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	bytes, err := json.Marshal(overview)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestRabbitmqTaskInitRabbitmqTaskList(t *testing.T) {
	sessionCookie, err := DoLogin(DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	params := RabbitmqTaskInitRabbitmqTaskListParam{
		ServerUrl:     url,
		Start:         0,
		Length:        10,
		SessionCookie: sessionCookie,
	}
	overview, err := RabbitmqTaskInitRabbitmqTaskList(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	bytes, err := json.Marshal(overview)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestTaskMonitorInitTaskMonitor(t *testing.T) {
	sessionCookie, err := DoLogin(DoLoginParam{
		ServerUrl:  url,
		LoginEmail: username,
		Password:   password,
	}, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	params := TaskMonitorInitTaskMonitorParam{
		ServerUrl:     url,
		Start:         0,
		Length:        10,
		SessionCookie: sessionCookie,
	}
	overview, err := TaskMonitorInitTaskMonitor(params, context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}
	bytes, err := json.Marshal(overview)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
