package cookbook_test

import (
	"testing"

	"github.com/dynastymasra/cookbook"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type JSendSuite struct {
	suite.Suite
}

func Test_JSendSuite(t *testing.T) {
	suite.Run(t, new(JSendSuite))
}

func (j *JSendSuite) Test_SuccessResponse() {
	resp := cookbook.SuccessResponse()

	assert.EqualValues(j.T(), cookbook.JSend{
		Status: "success",
	}, resp)
}

func (j *JSendSuite) Test_FailResponse() {
	resp := cookbook.FailResponse(&cookbook.JSON{
		"key": "value",
	})

	assert.EqualValues(j.T(), cookbook.JSend{
		Status: "failed",
		Data: &cookbook.JSON{
			"key": "value",
		},
	}, resp)
}

func (j *JSendSuite) Test_ErrorResponse() {
	resp := cookbook.ErrorResponse("message", "1234567890")

	assert.EqualValues(j.T(), cookbook.JSend{
		Status:  "error",
		Message: "message",
		Code:    "1234567890",
	}, resp)
}

func (j *JSendSuite) Test_SuccessDataResponse() {
	resp := cookbook.SuccessDataResponse(&cookbook.JSON{
		"test": "test",
	}, cookbook.NewMeta(cookbook.NewPage(1, 25, 100)), cookbook.NewLinks("next", "prev", "first", "last"))

	assert.EqualValues(j.T(), cookbook.JSend{
		Status: "success",
		Data: &cookbook.JSON{
			"test": "test",
		},
		Meta: &cookbook.Meta{
			Page: &cookbook.Page{
				Current: 1,
				Size:    25,
				Total:   100,
			},
		},
		Link: &cookbook.Links{
			First: "first",
			Next:  "next",
			Prev:  "prev",
			Last:  "last",
		},
	}, resp)
}

func (j *JSendSuite) Test_Stringify_Success() {
	expected := `{"status":"success","data":{"test":"test"}}`

	result := cookbook.SuccessDataResponse(&cookbook.JSON{
		"test": "test",
	}, nil, nil).Stringify()

	assert.JSONEq(j.T(), expected, result)
}

func (j *JSendSuite) Test_Stringify_Failed_Marshal() {
	ch := make(chan string)
	expected := "json: unsupported type: chan string"

	result := cookbook.SuccessDataResponse(&cookbook.JSON{
		"key": ch,
	}, nil, nil).Stringify()

	assert.Equal(j.T(), expected, result)
}
