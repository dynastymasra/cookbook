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

		var errs []message.ErrorMessage
		if !strings.Contains(mediaType, cookbook.HJSONType) {
			w.Header().Set(cookbook.HContentType, cookbook.HJSONType)
			w.Header().Set(cookbook.HAccept, cookbook.HJSONType)

			msg := message.ErrUnsupportedMediaTypeCode.ErrorMessage()
			errs = append(errs, message.ErrorMessage{
				Code:  msg.Code,
				Title: msg.Title,
				Error: msg.Error,
			})
		}

		if !strings.Contains(accept, cookbook.HJSONType) {
			w.Header().Set(cookbook.HContentType, cookbook.HJSONType)
			w.Header().Set(cookbook.HAccept, cookbook.HJSONType)

			msg := message.ErrRequestNotAcceptableCode.ErrorMessage()
			errs = append(errs, message.ErrorMessage{
				Code:  msg.Code,
				Title: msg.Title,
				Error: msg.Error,
			})
		}

		if len(errs) > 0 {
			w.WriteHeader(errs[0].Code.HTTPErrorMessage())
			fmt.Fprint(w, cookbook.FailResponse(message.ErrorMessageToJSONList(errs)).Stringify())
			return
		}

		next(w, r.WithContext(r.Context()))
	}
}
