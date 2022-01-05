package api

import (
	"context"
	"github.com/morikuni/failure"
	"github.com/rea1shane/http-pro"
	"net/http"
	"strings"
)

var h = http_pro.GetHttp(3, 500)

const (
	RequestError failure.StringCode = "request error"
)

// DoLogin 登录账号
func DoLogin(p DoLoginParam, ctx context.Context) (string, error) {
	url := p.ServerUrl + "/userReq/doLogin"
	method := "POST"
	payload := strings.NewReader("loginEmail=" + p.LoginEmail + "&password=" + p.Password)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", failure.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := h.RequestWithStringResponse(req, ctx)
	if err != nil {
		return "", err
	}
	return response, nil
}
