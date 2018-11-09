package cookbook

import "encoding/json"

// Jsend used jsend format with some modification
type Jsend struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// Meta data used for JSON response
type Meta struct {
	RequestID interface{} `json:"request_id,omitempty"`
	Links     *Links      `json:"links,omitempty"`
}

// Links for meta data JSON response
type Links struct {
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

// NewMeta build new meta struct
func NewMeta(requestID string, links *Links) *Meta {
	return &Meta{
		RequestID: requestID,
		Links:     links,
	}
}

// NewLinks build new links struct
func NewLinks(next, prev string) *Links {
	return &Links{
		Next: next,
		Prev: prev,
	}
}

// SuccessResponse used to return response with JSON format success
func SuccessResponse(meta *Meta) Jsend {
	return Jsend{Status: "success", Meta: meta}
}

// FailResponse is used to return response with JSON format if failure
func FailResponse(msg string, data interface{}, meta *Meta) Jsend {
	return Jsend{Status: "failed", Message: msg, Data: data, Meta: meta}
}

// SuccessDataResponse used to return response JSON format if have data value
func SuccessDataResponse(data interface{}, meta *Meta) Jsend {
	return Jsend{Status: "success", Data: data, Meta: meta}
}

// Stringify make Jsend struct to string
func (j Jsend) Stringify() string {
	toJSON, err := json.Marshal(j)
	if err != nil {
		return err.Error()
	}
	return string(toJSON)
}
