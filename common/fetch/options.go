package fetch

import (
	"net/http"
	"time"
)

type ClientOptions struct {
	DialTimeout           time.Duration
	TLSHandshakeTimeout   time.Duration
	ResponseHeaderTimeout time.Duration
	ExpectContinueTimeout time.Duration
	IdleConnTimeout       time.Duration
	MaxIdleConnect        int
	MaxIdleConnectPerHost int
	MaxConnectPerHost     int
	CheckRedirect         func(req *http.Request, via []*http.Request) error
}
