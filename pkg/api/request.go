package api

import (
	"context"
	"github.com/morikuni/failure"
	"github.com/rea1shane/http-pro"
	"net/http"
	"strings"
)

var (
	h      = http_pro.GetHttp(3, 500)
	cookie []*http.Cookie
)

const (
	RequestError failure.StringCode = "request error"
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
	body, err := http_pro.GetStringResponseBody(response)
	if err != nil {
		return err
	}
	if body != "\"success\"" {
		c := getReqFailureContext(method, url)
		c["payload"] = payloadStr
		c["reason"] = body
		return failure.New(RequestError, c)
	}
	cookie = response.Cookies()
	return nil
}

func getReqFailureContext(method, url string) failure.Context {
	return failure.Context{
		"method": method,
		"url":    url,
	}
}
