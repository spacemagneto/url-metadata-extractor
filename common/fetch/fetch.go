package fetch

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rs/zerolog"
)

type FetcherOptions func(*Fetcher)

func Transport(transport http.RoundTripper) FetcherOptions {
	return func(f *Fetcher) {
		f.transport = transport
	}
}

func UserAgent(userAgent string) FetcherOptions {
	return func(f *Fetcher) {
		if f.headers == nil {
			customHeaders := make(http.Header)
			f.headers = &customHeaders
		}

		f.headers.Set("User-Agent", userAgent)
	}
}

func Headers(headers map[string]string) FetcherOptions {
	return func(f *Fetcher) {
		customHeaders := make(http.Header)
		for header, value := range headers {
			customHeaders.Add(header, value)
		}

		f.headers = &customHeaders
	}
}

func Client(client *http.Client) FetcherOptions {
	return func(f *Fetcher) {
		if client != nil {
			f.client = client
		}
	}
}

type Fetcher struct {
	client    *http.Client
	transport http.RoundTripper
	logger    *zerolog.Logger
	headers   *http.Header
}

func New(log *zerolog.Logger) *Fetcher {
	l := log.With().Str("service", "http").Logger()

	return &Fetcher{logger: &l, client: &http.Client{}}
}

func (f *Fetcher) Fetch(ctx context.Context, link string) (*http.Response, error) {
	l := f.logger.With().Str("method", "Fetch").Logger()

	u, err := url.Parse(link)
	if err != nil {
		l.Error().Err(err).Msg("failed parse link")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := f.doRequest(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (f *Fetcher) doRequest(req *http.Request) (*http.Response, error) {
	return f.client.Do(req)
}
