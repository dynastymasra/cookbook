package middleware

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/dynastymasra/cookbook"

	"github.com/sirupsen/logrus"

	"github.com/urfave/negroni"
)

func RequestID(reqID string) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now().UTC()

		requestID := r.Header.Get(cookbook.XRequestID)
		if len(requestID) < 1 {
			entropy := rand.New(rand.NewSource(rand.Int63n(startTime.UnixNano())))
			requestID = ulid.MustNew(ulid.Timestamp(startTime), entropy).String()
		}
		next(w, r.WithContext(context.WithValue(r.Context(), reqID, requestID)))
	}
}

// LogrusLog middleware function for log HTTP request. combine with RequestID middleware first to add request id in log
func LogrusLog(name, reqID string) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now().UTC()

		requestID := r.Header.Get(cookbook.XRequestID)
		if len(requestID) < 1 {
			entropy := rand.New(rand.NewSource(rand.Int63n(startTime.UnixNano())))
			requestID = ulid.MustNew(ulid.Timestamp(startTime), entropy).String()
		}
		next(w, r.WithContext(context.WithValue(r.Context(), reqID, requestID)))

		responseTime := time.Now().UTC()
		deltaTime := responseTime.Sub(startTime)

		logrus.WithFields(logrus.Fields{
			"service": name,
			"start":   startTime.Format(time.RFC3339),
			"delta":   deltaTime,
			"finish":  responseTime.Format(time.RFC3339),
			"proxy":   r.RemoteAddr,
			"url":     r.URL.Path,
			"method":  r.Method,
			"headers": r.Header,
			reqID:     requestID,
		}).Infoln("HTTP Request")
	}
}
