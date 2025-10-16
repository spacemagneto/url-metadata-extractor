package fetch

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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
}
