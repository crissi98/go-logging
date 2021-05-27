package logging

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

var CorrelationIdHeaderName = "X-Correlation-ID"

type ContextKey string

const (
	CorrelationIdContextKey ContextKey = "correlationId"
)

type CorrelationIdHandler struct {
	Wrapped http.Handler
}

func (handler *CorrelationIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	correlationId := r.Header.Get(CorrelationIdHeaderName)
	if correlationId == "" {
		correlationId = uuid.NewString()
	}
	r.WithContext(context.WithValue(r.Context(), CorrelationIdContextKey, correlationId))
	w.Header().Add(CorrelationIdHeaderName, correlationId)
	handler.Wrapped.ServeHTTP(w, r)

}
