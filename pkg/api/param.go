package api

import "net/http"

type DoLoginParam struct {
	ServerUrl  string
	LoginEmail string
	Password   string
}

type HomeCountParam struct {
	ServerUrl     string
	SessionCookie *http.Cookie
}

type HomeStatisParam struct {
	ServerUrl     string
	SessionCookie *http.Cookie
}

type GroupInitGroupParam struct {
	ServerUrl     string
	Start         int
	Length        int
	SessionCookie *http.Cookie
}

type WorkerInitWorkerParam struct {
	ServerUrl     string
	Start         int
	Length        int
	SessionCookie *http.Cookie
}

type MysqlTaskMysqlTaskDatasParam struct {
	ServerUrl     string
	Start         int
	Length        int
	SessionCookie *http.Cookie
}

type HbaseTaskInitHbaseTaskListParam struct {
	ServerUrl     string
	Start         int
	Length        int
	SessionCookie *http.Cookie
}

type RabbitmqTaskInitRabbitmqTaskListParam struct {
	ServerUrl     string
	Start         int
	Length        int
	SessionCookie *http.Cookie
}

type TaskMonitorInitTaskMonitorParam struct {
	ServerUrl     string
	Start         int
	Length        int
	SessionCookie *http.Cookie
}
