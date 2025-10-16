package fetch

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestFetcher(t *testing.T) {
	t.Parallel()

	// Initialize the logger instance using zerolog to output logs to the standard output.
	// This logger will be used for logging purposes throughout the test run.
	logger := zerolog.New(os.Stdout)

	t.Run("Success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}))
		defer server.Close()

		fetcher := New(&logger)

		response, err := fetcher.Fetch(context.Background(), server.URL)
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("RealRequest", func(t *testing.T) {
		fetcher := &Fetcher{
			client: &http.Client{
				Timeout: 10 * time.Second,
			},
			logger: &logger,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		response, err := fetcher.Fetch(ctx, "https://www.nytimes.com")
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})
}
