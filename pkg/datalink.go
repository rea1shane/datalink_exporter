package pkg

import "net/http"

type Server struct {
	Protocol string  `yaml:"protocol"`
	Host     string  `yaml:"host"`
	Port     string  `yaml:"port"`
	Session  Session `yaml:"session"`
	url      string
}

type Session struct {
	Timeout       int          `yaml:"timeout"` // Timeout 单位 ms
	sessionCookie *http.Cookie // sessionId 默认有效期 24 小时
	authTimestamp int64
}
