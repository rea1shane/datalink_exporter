package api

import (
	"context"
	"encoding/json"
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

func getReqFailureContext(method, url string) failure.Context {
	return failure.Context{
		"method": method,
		"url":    url,
	}
}
