package http

import (
	"github.com/rs/zerolog"
)

type Client struct {
	logger *zerolog.Logger
}

func New(log *zerolog.Logger) *Client {
	l := log.With().Str("service", "http").Logger()

	return &Client{logger: &l}
}
