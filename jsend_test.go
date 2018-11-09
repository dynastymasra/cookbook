package cookbook_test

import (
	"testing"

	"github.com/dynastymasra/cookbook"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type JsendSuite struct {
	suite.Suite
}

func Test_JsendSuite(t *testing.T) {
	suite.Run(t, new(JsendSuite))
}

func (j *JsendSuite) Test_SuccessResponse() {
	resp := cookbook.SuccessResponse(cookbook.NewMeta("request_id", cookbook.NewLinks("next", "prev")))

	assert.EqualValues(j.T(), cookbook.Jsend{
		Status: "success",
		Meta: &cookbook.Meta{
			RequestID: "request_id",
			Links: &cookbook.Links{
				Next: "next",
				Prev: "prev",
			},
		},
	}, resp)
}

func (j *JsendSuite) Test_FailResponse() {
	resp := cookbook.FailResponse("message", nil, cookbook.NewMeta("request_id", nil))

	assert.EqualValues(j.T(), cookbook.Jsend{
		Status:  "failed",
		Message: "message",
		Meta: &cookbook.Meta{
			RequestID: "request_id",
		},
	}, resp)
}

func (j *JsendSuite) Test_SuccessDataResponse() {
	resp := cookbook.SuccessDataResponse(map[string]interface{}{
		"test": "test",
	}, cookbook.NewMeta("request_id", cookbook.NewLinks("next", "prev")))

	assert.EqualValues(j.T(), cookbook.Jsend{
		Status: "success",
		Data: map[string]interface{}{
			"test": "test",
		},
		Meta: &cookbook.Meta{
			RequestID: "request_id",
			Links: &cookbook.Links{
				Next: "next",
				Prev: "prev",
			},
		},
	}, resp)
}

func (j *JsendSuite) Test_Stringify_Success() {
	expected := `{"status":"success","data":{"test":"test"}}`

	result := cookbook.SuccessDataResponse(map[string]interface{}{
		"test": "test",
	}, nil).Stringify()

	assert.JSONEq(j.T(), expected, result)
}

func (j *JsendSuite) Test_Stringify_Failed_Marshal() {
	ch := make(chan string)
	expected := "json: unsupported type: chan string"

	result := cookbook.SuccessDataResponse(ch, nil).Stringify()

	assert.Equal(j.T(), expected, result)
}
