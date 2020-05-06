package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/gommon/random"

	"github.com/dynastymasra/cookbook"
	"github.com/urfave/negroni"
)

// RequestID middleware for get request id from client, if request id not found it will set to UUID
func RequestID() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		requestID := r.Header.Get(cookbook.XRequestID)
		if len(requestID) < 1 {
			requestID = random.New().String(12, random.Alphanumeric)
		}
		next(w, r.WithContext(context.WithValue(r.Context(), cookbook.RequestID, requestID)))
	}
}
