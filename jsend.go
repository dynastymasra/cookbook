package cookbook

import "encoding/json"

// JSend used JSend format with some modification
type JSend struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// Meta data used for JSON response
type Meta struct {
	Links *Links `json:"links,omitempty"`
}

// Links for meta data JSON response
type Links struct {
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

// NewMeta build new meta struct
func NewMeta(links *Links) *Meta {
	return &Meta{
		Links: links,
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
func SuccessResponse(meta *Meta) JSend {
	return JSend{Status: "success", Meta: meta}
}

// FailResponse is used to return response with JSON format if failure
func FailResponse(msg string, data interface{}, meta *Meta) JSend {
	return JSend{Status: "failed", Message: msg, Data: data, Meta: meta}
}

// SuccessDataResponse used to return response JSON format if have data value
func SuccessDataResponse(data interface{}, meta *Meta) JSend {
	return JSend{Status: "success", Data: data, Meta: meta}
}

// Stringify make JSend struct to string
func (j JSend) Stringify() string {
	toJSON, err := json.Marshal(j)
	if err != nil {
		return err.Error()
	}
	return string(toJSON)
}
