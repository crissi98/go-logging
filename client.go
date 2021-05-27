package logging

import "net/http"

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}

func (c *HttpClient) Do(r *http.Request) (*http.Response, error) {
	correlationID := r.Context().Value(CorrelationIdContextKey)
	if correlationID != nil {
		r.Header.Add(CorrelationIdHeaderName, correlationID.(string))
	}
	return c.client.Do(r)
}
