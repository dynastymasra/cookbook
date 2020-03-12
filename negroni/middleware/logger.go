package middleware

import (
	"net/http"
	"time"

	"github.com/dynastymasra/cookbook"

	"github.com/sirupsen/logrus"

	"github.com/urfave/negroni"
)

func LogrusLog(name string) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now().UTC()
		requestID := r.Header.Get(cookbook.XRequestID)

		next(w, r)

		responseTime := time.Now().UTC()
		deltaTime := responseTime.Sub(startTime)

		logrus.WithFields(logrus.Fields{
			"request_time":     startTime.Format(time.RFC3339),
			"delta_time":       deltaTime,
			"response_time":    responseTime.Format(time.RFC3339),
			"request_proxy":    r.RemoteAddr,
			"url":              r.URL.Path,
			"method":           r.Method,
			"request_source":   r.Header.Get("X-FORWARDED-FOR"),
			"headers":          r.Header,
			cookbook.RequestID: requestID,
		}).Infoln("HTTP Request", name)
	}
}
