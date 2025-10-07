package handler

import (
	"net/http"

	"github.com/rs/zerolog"
)

// HealthCheckHandler is an HTTP handler that responds to health check requests.
type HealthCheckHandler struct {
	logger *zerolog.Logger
}

// NewHealthCheckHandler creates a new HealthCheckHandler with the provided zerolog.Logger.
func NewHealthCheckHandler(logger *zerolog.Logger) *HealthCheckHandler {
	l := logger.With().Str("service", "health-check").Logger()
	return &HealthCheckHandler{logger: &l}
}

// HealthCheckHandler is an HTTP handler function that responds to health check requests.
// This function is typically used in web applications to verify that the server is operational.
// It listens for GET requests, sets the appropriate response headers, and returns a simple JSON payload
// indicating the server's status as "OK". This method ensures that the application provides a lightweight
// and reliable endpoint for monitoring and health checks.
func (h *HealthCheckHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET.
	// Only GET requests are processed, ensuring compliance with the expected behavior for health checks.
	if r.Method == http.MethodGet {
		// Set the Content-Type header to indicate the response format is JSON.
		// This ensures clients correctly interpret the response as JSON data.
		w.Header().Set("Content-Type", "application/json")
		// Write a 200 OK status to the response header.
		// This status code indicates that the server is healthy and operational.
		w.WriteHeader(http.StatusOK)
		// Write a JSON response body with the server status.
		// The status field is set to "OK" to signal that the health check was successful.
		_, _ = w.Write([]byte(`{"status": "OK"}`))
	}
}
