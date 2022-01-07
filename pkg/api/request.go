package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/morikuni/failure"
	httpro "github.com/rea1shane/http-pro"
	"net/http"
	"strings"
)

var (
	h             = httpro.GetHttp(3, 500)
	sessionCookie *http.Cookie
)

const (
	SessionKey     string             = "JSESSIONID"
	RequestError   failure.StringCode = "request error"
	NoSessionError failure.StringCode = "No session id in cookies"
)

// DoLogin 登录
func DoLogin(p DoLoginParam, ctx context.Context) error {
	url := p.ServerUrl + "/userReq/doLogin"
	method := "POST"
	payloadStr := "loginEmail=" + p.LoginEmail + "&password=" + p.Password
	payload := strings.NewReader(payloadStr)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return failure.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := h.Request(req, ctx)
	if err != nil {
		return err
	}
	body, err := httpro.GetStringResponseBody(response)
	if err != nil {
		return err
	}
	if body != "\"success\"" {
		c := getReqFailureContext(method, url)
		c["payload"] = payloadStr
		c["reason"] = body
		return failure.New(RequestError, c)
	}
	for i, cookie := range response.Cookies() {
		cookieMap := make(map[string]string)
		cookieMap[cookie.Name] = cookie.Value
		if cookie.Name == SessionKey {
			sessionCookie = cookie
			break
		}
		if i == len(response.Cookies())-1 {
			c := getReqFailureContext(method, url)
			c["payload"] = payloadStr
			bytes, _ := json.Marshal(cookieMap)
			c["cookies"] = string(bytes)
			return failure.New(NoSessionError, c)
		}
	}
	return nil
}

// HomeCount 获取主页的统计数据
func HomeCount(p HomeCountParam, ctx context.Context) (Overview, error) {
	url := p.ServerUrl + "/home/count"
	method := "GET"
	responseStruct := Overviews{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return Overview{}, failure.Wrap(err)
	}
	req.AddCookie(sessionCookie)
	response, err := h.Request(req, ctx)
	if err != nil {
		return Overview{}, err
	}
	err = httpro.GetStructResponseBody(response, &responseStruct)
	if err != nil {
		return Overview{}, err
	}
	return responseStruct.AaData[0], nil
}

// HomeStatis 获取主页的图表数据
func HomeStatis(p HomeStatisParam, ctx context.Context) (Statis, error) {
	url := p.ServerUrl + "/home/statis"
	method := "POST"
	responseStruct := Statis{}
	payloadStr := "groupId=-1"
	payload := strings.NewReader(payloadStr)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return Statis{}, failure.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(sessionCookie)
	response, err := h.Request(req, ctx)
	if err != nil {
		return Statis{}, err
	}
	err = httpro.GetStructResponseBody(response, &responseStruct)
	if err != nil {
		return Statis{}, err
	}
	return responseStruct, nil
}

// GroupInitGroup 获取集群管理中分组信息
func GroupInitGroup(p GroupInitGroupParam, ctx context.Context) ([]Group, error) {
	url := p.ServerUrl + "/group/initGroup"
	method := "POST"
	responseStruct := Groups{}
	payloadStr := fmt.Sprintf(`{
    "draw": 1,
    "columns": [
        {
            "data": "id",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "groupName",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "groupDesc",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "groupState",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "lastRebalancedTime",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "generationId",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "createTime",
            "name": "",
            "searchable": true,
            "orderable": false,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "id",
            "name": "",
            "searchable": true,
            "orderable": false,
            "search": {
                "value": "",
                "regex": false
            }
        }
    ],
    "order": [
        {
            "column": 0,
            "dir": "asc"
        }
    ],
    "start": %d,
    "length": %d,
    "search": {
        "value": "",
        "regex": false
    }
}`, p.Start, p.Length)
	payload := strings.NewReader(payloadStr)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, failure.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(sessionCookie)
	response, err := h.Request(req, ctx)
	if err != nil {
		return nil, err
	}
	err = httpro.GetStructResponseBody(response, &responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct.AaData, nil
}

// WorkerInitWorker 获取集群管理中机器信息
func WorkerInitWorker(p WorkerInitWorkerParam, ctx context.Context) ([]Worker, error) {
	url := p.ServerUrl + "/worker/initWorker"
	method := "POST"
	responseStruct := Workers{}
	payloadStr := fmt.Sprintf(`{
    "draw": 1,
    "columns": [
        {
            "data": "id",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "workerName",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "workerState",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "workerAddress",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "restPort",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "groupName",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "startTime",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "createTime",
            "name": "",
            "searchable": true,
            "orderable": false,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "id",
            "name": "",
            "searchable": true,
            "orderable": false,
            "search": {
                "value": "",
                "regex": false
            }
        }
    ],
    "order": [
        {
            "column": 0,
            "dir": "asc"
        }
    ],
    "start": %d,
    "length": %d,
    "search": {
        "value": "",
        "regex": false
    },
    "groupId": "-1"
}`, p.Start, p.Length)
	payload := strings.NewReader(payloadStr)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, failure.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(sessionCookie)
	response, err := h.Request(req, ctx)
	if err != nil {
		return nil, err
	}
	err = httpro.GetStructResponseBody(response, &responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct.AaData, nil
}

// MysqlTaskMysqlTaskDatas 获取增量任务中的 MySQL task
func MysqlTaskMysqlTaskDatas(p MysqlTaskMysqlTaskDatasParam, ctx context.Context) ([]MysqlTask, error) {
	url := p.ServerUrl + "/mysqlTask/mysqlTaskDatas"
	method := "POST"
	responseStruct := MysqlTasks{}
	payloadStr := fmt.Sprintf(`{
    "draw": 1,
    "columns": [
        {
            "data": "id",
            "name": "",
            "searchable": true,
            "orderable": false,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "id",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "taskName",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "detail",
            "name": "",
            "searchable": true,
            "orderable": false,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "targetState",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "listenedState",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "groupId",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "workerId",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "currentTimeStamp",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "startTime",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "readerIp",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "latestEffectSyncLogFileName",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "latestEffectSyncLogFileOffset",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "taskSyncStatus",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "shadowCurrentTimeStamp",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "shadowLatestEffectSyncLogFileName",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "shadowLatestEffectSyncLogFileOffset",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": null,
            "name": "",
            "searchable": false,
            "orderable": false,
            "search": {
                "value": "",
                "regex": false
            }
        }
    ],
    "order": [
        {
            "column": 0,
            "dir": "asc"
        }
    ],
    "start": %d,
    "length": %d,
    "search": {
        "value": "",
        "regex": false
    },
    "readerMediaSourceId": "-1",
    "groupId": "-1",
    "id": "-1"
}`, p.Start, p.Length)
	payload := strings.NewReader(payloadStr)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, failure.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(sessionCookie)
	response, err := h.Request(req, ctx)
	if err != nil {
		return nil, err
	}
	err = httpro.GetStructResponseBody(response, &responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct.AaData, nil
}

// HbaseTaskInitHbaseTaskList 获取增量任务中的 Hbase task
func HbaseTaskInitHbaseTaskList(p HbaseTaskInitHbaseTaskListParam, ctx context.Context) ([]HbaseTask, error) {
	url := p.ServerUrl + "/hbaseTask/initHbaseTaskList"
	method := "POST"
	responseStruct := HbaseTasks{}
	payloadStr := fmt.Sprintf(`{
    "draw": 1,
    "columns": [
        {
            "data": "id",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "taskName",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "targetState",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "listenedState",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "groupId",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "workerId",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": "startTime",
            "name": "",
            "searchable": true,
            "orderable": true,
            "search": {
                "value": "",
                "regex": false
            }
        },
        {
            "data": null,
            "name": "",
            "searchable": false,
            "orderable": false,
            "search": {
                "value": "",
                "regex": false
            }
        }
    ],
    "order": [
        {
            "column": 0,
            "dir": "asc"
        }
    ],
    "start": %d,
    "length": %d,
    "search": {
        "value": "",
        "regex": false
    },
    "readerMediaSourceId": "-1",
    "groupId": "-1",
    "id": "-1"
}`, p.Start, p.Length)
	payload := strings.NewReader(payloadStr)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, failure.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(sessionCookie)
	response, err := h.Request(req, ctx)
	if err != nil {
		return nil, err
	}
	err = httpro.GetStructResponseBody(response, &responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct.AaData, nil
}

func getReqFailureContext(method, url string) failure.Context {
	return failure.Context{
		"method": method,
		"url":    url,
	}
}
