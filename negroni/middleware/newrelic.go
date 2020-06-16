package middleware

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/cookbook"
	"github.com/gorilla/mux"
	newrelic "github.com/newrelic/go-agent"
	"github.com/urfave/negroni"
)

// NewRelicHandler use new relic handler in middleware
func NewRelicHandler(app newrelic.Application) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		_, nextHandler := newrelic.WrapHandleFunc(app, fmt.Sprintf("%s - %s", r.Method, getRoutePath(r)), next)
		nextHandler(w, r)
	})
}

// NewRelicInstrumentation used for add request id to new relic attribute
func NewRelicInstrumentation(reqID string) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if trx := newrelic.FromContext(r.Context()); trx != nil {
			trx.AddAttribute(cookbook.XRequestID, r.Context().Value(reqID))

			next(w, newrelic.RequestWithTransactionContext(r, trx))
		}
	})
}

func getRoutePath(r *http.Request) string {
	if route := mux.CurrentRoute(r); route != nil {
		if path, err := route.GetPathTemplate(); err != nil {
			return path
		}
	}
	return r.URL.Path
}
