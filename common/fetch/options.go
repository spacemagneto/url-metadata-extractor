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

func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		DialTimeout:           10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 2 * time.Second,
		IdleConnTimeout:       90 * time.Second,
		MaxIdleConnect:        100,
		MaxIdleConnectPerHost: 25,
		MaxConnectPerHost:     50,
	}
}
