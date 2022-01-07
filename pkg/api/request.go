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

// DoLogin 登录账号
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
	responseStruct := Overview{}
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
	return responseStruct, nil
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

func GroupInitGroup(p GroupInitGroupParam, ctx context.Context) (Group, error) {
	url := p.ServerUrl + "/group/initGroup"
	method := "POST"
	responseStruct := Group{}
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
		return Group{}, failure.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(sessionCookie)
	response, err := h.Request(req, ctx)
	if err != nil {
		return Group{}, err
	}
	err = httpro.GetStructResponseBody(response, &responseStruct)
	if err != nil {
		return Group{}, err
	}
	return responseStruct, nil
}

func getReqFailureContext(method, url string) failure.Context {
	return failure.Context{
		"method": method,
		"url":    url,
	}
}
