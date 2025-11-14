package fetch

import (
	"net/http"
	"net/http/httptrace"
	"time"
)

// TraceRequest captures precise timing information during an HTTP request's network lifecycle.
// It uses Go's http-trace package to hook into connection and response events, providing
// accurate measurements of connection establishment time and time-to-first-byte.
// This is an important part of analyzing delays, monitoring performance, and debugging
// slow network requests.
type TraceRequest struct {
	start             time.Time
	connect           time.Time
	requestDuration   time.Duration
	firstByteDuration time.Duration
}

// trace creates and returns a fully configured httptrace.ClientTrace that records
// timing information directly into this TraceRequest instance as the request progresses.
// All relevant network and response events are captured at the exact moments defined
// by the Go httptrace specification, ensuring high-precision, real-world timing data.
func (t *TraceRequest) trace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn:              func(hostPort string) { t.start = time.Now() },
		GotFirstResponseByte: func() { t.firstByteDuration = time.Since(t.connect) },
		ConnectStart:         func(network, addr string) { t.connect = time.Now() },
		ConnectDone:          func(network, addr string, err error) { t.requestDuration = time.Since(t.connect) },
	}
}

// TraceRequest instruments the provided HTTP request with client-side tracing using this instance.
// It attaches the configured trace to the request context, enabling all timing callbacks
// to execute during the actual network operation.
//
// Returns a new request with tracing enabled. The original request remains unchanged.
// Multiple calls are safe — only the last applied trace will be active.
func (t *TraceRequest) TraceRequest(req *http.Request) *http.Request {
	return req.WithContext(httptrace.WithClientTrace(req.Context(), t.trace()))
}
