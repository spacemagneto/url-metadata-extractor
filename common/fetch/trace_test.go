package fetch

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTraceRequest(t *testing.T) {
	t.Parallel()

	// SuccessCase verifies that TraceRequest correctly captures all timing phases
	// during a successful HTTP request to a test server with artificial response delay.
	// This test ensures that GetConn, ConnectStart/ConnectDone, and GotFirstResponseByte
	// hooks are triggered in the correct order and record accurate durations.
	t.Run("SuccessCase", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(80 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ok"))
		}))
		defer server.Close()

		// Initialize a new TraceRequest instance to capture timings.
		tracer := &TraceRequest{}

		// Use the test server's transport to ensure requests hit our controlled endpoint.
		client := server.Client()

		// Create the base HTTP request.
		baseRequest, err := http.NewRequest("GET", server.URL, nil)
		assert.NoError(t, err, "Failed to create base HTTP request")

		tracedRequest := tracer.TraceRequest(baseRequest)

		requestStartTime := time.Now()

		response, err := client.Do(tracedRequest)
		requestEndTime := time.Now()

		assert.NoError(t, err, "HTTP request should complete without error")
		assert.NotNil(t, response, "HTTP response should not be nil")
		assert.Equal(t, http.StatusOK, response.StatusCode, "Expected HTTP 200 OK status")

		assert.False(t, tracer.start.IsZero(), "GetConn must set start time")
		assert.True(t, tracer.start.After(requestStartTime.Add(-50*time.Millisecond)), "Start time should be close to actual request dispatch time")
		assert.True(t, tracer.start.Before(requestEndTime), "Start time must be before request completion")

		assert.False(t, tracer.connect.IsZero(), "ConnectStart must set connect time")
		assert.NotZero(t, tracer.requestDuration, "ConnectDone must record connection duration")
		assert.Less(t, tracer.requestDuration, 100*time.Millisecond, "Connection establishment on localhost should be fast")

		assert.NotZero(t, tracer.firstByteDuration, "GotFirstResponseByte must record TTFB")
		assert.GreaterOrEqual(t, tracer.firstByteDuration, 70*time.Millisecond, "Time to first byte must reflect the 80ms server sleep")
		assert.Less(t, tracer.firstByteDuration, 300*time.Millisecond, "Time to first byte should not be excessively long")

		totalToFirstByte := tracer.requestDuration + tracer.firstByteDuration
		assert.GreaterOrEqual(t, totalToFirstByte, 75*time.Millisecond, "Total time from connect to first byte must include server processing delay")
	})
}
