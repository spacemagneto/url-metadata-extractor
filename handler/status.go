package handler

import "github.com/rs/zerolog"

type HealthCheckHandler struct {
	logger *zerolog.Logger
}

func New(logger *zerolog.Logger) *HealthCheckHandler {
	l := logger.With().Str("service", "health-check").Logger()
	return &HealthCheckHandler{logger: &l}
}
