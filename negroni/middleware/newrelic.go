package middleware

import (
	"net/http"

	"github.com/dynastymasra/cookbook"
	"github.com/newrelic/go-agent"
	"github.com/urfave/negroni"
)

// NewRelicHandler use new relic handler in middleware
func NewRelicHandler(app newrelic.Application, pattern string) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		_, nextHandler := newrelic.WrapHandleFunc(app, pattern, next)
		nextHandler(w, r)
	})
}

// NewRelicInstrumentation used for add request id to new relic attribute
func NewRelicInstrumentation() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if trx := newrelic.FromContext(r.Context()); trx != nil {
			trx.AddAttribute(cookbook.XRequestID, r.Context().Value(cookbook.RequestID))

			next(w, newrelic.RequestWithTransactionContext(r, trx))
		}
	})
}
