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
func AcceptMediaTypeJSON(request string) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		requestID := r.Header.Get(cookbook.XRequestID)
		accept := r.Header.Get(cookbook.HAccept)
		mediaType := r.Header.Get(cookbook.HContentType)

		if len(requestID) < 1 {
			requestID = fmt.Sprintf("%v", r.Context().Value(request))
		}

		if !strings.Contains(accept, cookbook.HJSONType) {
			w.Header().Set(cookbook.HContentType, cookbook.HJSONType)
			w.Header().Set(cookbook.HAccept, cookbook.HJSONType)
			w.Header().Set(cookbook.XRequestID, requestID)

			res := message.ErrRequestNotAcceptableCode.ErrorMessage()
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprint(w, cookbook.FailResponse([]cookbook.JSON{
				{
					"code":    res.Code,
					"title":   res.Title,
					"message": res.Error.Error(),
				},
			}).Stringify())
			return
		}

		if !strings.Contains(mediaType, cookbook.HJSONType) {
			w.Header().Set(cookbook.HContentType, cookbook.HJSONType)
			w.Header().Set(cookbook.HAccept, cookbook.HJSONType)
			w.Header().Set(cookbook.XRequestID, requestID)

			res := message.ErrUnsupportedMediaTypeCode.ErrorMessage()
			w.WriteHeader(http.StatusUnsupportedMediaType)
			fmt.Fprint(w, cookbook.FailResponse([]cookbook.JSON{
				{
					"code":    res.Code,
					"title":   res.Title,
					"message": res.Error.Error(),
				},
			}).Stringify())
			return
		}

		next(w, r.WithContext(r.Context()))
	}
}
