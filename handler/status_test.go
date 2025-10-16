package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	t.Parallel()

	// Initialize a zerolog logger with a specific context for the health-check service.
	// This logger is used by the HealthCheckHandler to log relevant information during request processing.
	logger := zerolog.New(nil).With().Str("service", "health-check").Logger()

	// Create a new HealthCheckHandler instance with the configured logger.
	// This handler will be used to process HTTP requests in the test scenarios.
	handler := NewHealthCheckHandler(&logger)

	// ValidRequest tests the behavior of the HealthCheckHandler method when a GET request is received.
	// This test ensures that the method correctly processes a GET request to the health check endpoint,
	// returning a 200 OK status, setting the appropriate Content-Type header, and providing a JSON payload
	// indicating the server status as "OK". The goal is to verify that the health check endpoint functions
	// correctly for monitoring and operational verification.
	t.Run("ValidRequest", func(t *testing.T) {
		// Create a new HTTP GET request to simulate a health check request to the "/health" endpoint.
		// This request mimics a typical health check call made by monitoring systems to verify server status.
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		// Assert that no error occurred during request creation to ensure the test setup is valid.
		assert.NoError(t, err, "Expected no error when creating GET request for health check")

		// Create a ResponseRecorder to capture the handler's response.
		// This allows inspection of the status code, headers, and response body for verification.
		rr := httptest.NewRecorder()

		// Call the HealthCheckHandler method with the response recorder and request.
		// This simulates the server processing a health check request and writing the response.
		handler.HealthCheckHandler(rr, req)

		// Assert that the response status code is 200 OK.
		// This confirms that the handler correctly processes GET requests and indicates a healthy server.
		assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200 OK for GET request")

		// Assert that the Content-Type header is set to application/json.
		// This ensures the response format is correctly indicated for clients expecting JSON.
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Expected Content-Type header to be application/json")

		// Assert that the response body contains the expected JSON payload.
		// The payload should be {"status": "OK"}, indicating a successful health check.
		assert.Equal(t, `{"status": "OK"}`, rr.Body.String(), "Expected response body to be {\"status\": \"OK\"}")
	})
}
