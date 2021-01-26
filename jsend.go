package cookbook

import "encoding/json"

// JSend used JSend format with some modification
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"meta": {
//		"page": {
//			"number": "page number",
//			"size": "page size",
//			"total": "total page"
//		}
//	},
//	"links": {
//		"first": "first page",
//		"next": "next page",
//		"prev": "previous page",
//		"last": "last page"
//	},
//	"status": "message status return with success, failed, or error",
//	"message": "return message if error",
//	"data": "return data if success with data or failed",
//	"code": "error status code"
//}
type JSend struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
	Link    *Links      `json:"links,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}

// Meta data used for JSON response
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"meta": {
//		"page": {
//			"number": "page number",
//			"size": "page size",
//			"total": "total page"
//		}
//	}
//}
type Meta struct {
	Page *Page `json:"page,omitempty"`
}

// Page pagination information data
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
type Page struct {
	Number int `json:"number,omitempty"`
	Size   int `json:"size,omitempty"`
	Total  int `json:"total,omitempty"`
}

// NewPage of return data with total of data, size per page, and number page
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"page": {
//		"number": "page number",
//		"size": "page size",
//		"total": "total page"
//	}
//}
func NewPage(number, size, total int) *Page {
	return &Page{
		Number: number,
		Size:   size,
		Total:  total,
	}
}

// Links for meta data JSON response
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"first": "first page",
//	"next": "next page",
//	"prev": "previous page",
//	"last": "last page"
//}
type Links struct {
	First string `json:"first,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Last  string `json:"last,omitempty"`
}

// NewMeta build new meta struct
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"meta": {
//		"page": {
//			"number": "page number",
//			"size": "page size",
//			"total": "total page"
//		}
//	}
//}
func NewMeta(page *Page) *Meta {
	return &Meta{
		Page: page,
	}
}

// NewLinks build new links struct
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"first": "first page",
//	"next": "next page",
//	"prev": "previous page",
//	"last": "last page"
//}
func NewLinks(next, prev, first, last string) *Links {
	return &Links{
		Next:  next,
		Prev:  prev,
		First: first,
		Last:  last,
	}
}

// SuccessResponse used to return response with JSON format success
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"status": "success"
//}
func SuccessResponse() JSend {
	return JSend{Status: "success"}
}

// FailResponse is used to return response with JSON format if failure from client side
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"status": "failed",
//	"data": [
//		{
//			"field": "field",
//			"code": "failed code",
//			"message": "error message"
//		}
//	],
//}
func FailResponse(data interface{}) JSend {
	return JSend{Status: "failed", Data: data}
}

// ErrorResponse is used return response with JSON format if failure in server side
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"status": "error",
//	"message": "error message",
//	"code": "error code"
//}
func ErrorResponse(msg string, code interface{}) JSend {
	return JSend{Status: "error", Message: msg, Code: code}
}

// SuccessDataResponse used to return response JSON format if have data value
// See https://jsonapi.org/ and https://github.com/omniti-labs/jsend for references
//{
//	"status": "success",
//	"data": "response data",
//	"meta": "optional",
//	"links": "optional"
//}
func SuccessDataResponse(data interface{}, meta *Meta, links *Links) JSend {
	return JSend{Status: "success", Data: data, Meta: meta, Link: links}
}

// Stringify make JSend struct to string
func (j JSend) Stringify() string {
	toJSON, err := json.Marshal(j)
	if err != nil {
		return err.Error()
	}
	return string(toJSON)
}
