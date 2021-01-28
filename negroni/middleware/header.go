package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dynastymasra/cookbook"

	"github.com/dynastymasra/cookbook/message"
	"github.com/urfave/negroni/v2"
)

// AcceptMediaTypeJSON is a middleware to check content-type and accept value in HTTP header is application/json
func AcceptMediaTypeJSON() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		accept := r.Header.Get(cookbook.HAccept)
		mediaType := r.Header.Get(cookbook.HContentType)

		var errs []message.FailedMessage
		if !strings.Contains(mediaType, cookbook.HJSONType) {
			msg := message.FailedUnsupportedMediaType()
			errs = append(errs, message.FailedMessage{
				Code:    msg.Code,
				Title:   msg.Title,
				Message: msg.Message,
			})
		}

		if !strings.Contains(accept, cookbook.HJSONType) {
			msg := message.FailedRequestNotAcceptable()
			errs = append(errs, message.FailedMessage{
				Code:    msg.Code,
				Title:   msg.Title,
				Message: msg.Message,
			})
		}

		if len(errs) > 0 {
			w.Header().Set(cookbook.HContentType, cookbook.HJSONTypeUTF8)
			w.Header().Set(cookbook.HAccept, cookbook.HJSONType)

			w.WriteHeader(message.HTTPStatusCode(errs[0].Code))
			fmt.Fprint(w, cookbook.FailResponse(errs).Stringify())
			return
		}

		next(w, r.WithContext(r.Context()))
	}
}
