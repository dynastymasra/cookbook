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
	resp := cookbook.FailResponse("message", nil)

	assert.EqualValues(j.T(), cookbook.JSend{
		Status:  "failed",
		Message: "message",
	}, resp)
}

func (j *JSendSuite) Test_SuccessDataResponse() {
	resp := cookbook.SuccessDataResponse(map[string]interface{}{
		"test": "test",
	}, cookbook.NewMeta(cookbook.NewLinks("next", "prev")))

	assert.EqualValues(j.T(), cookbook.JSend{
		Status: "success",
		Data: map[string]interface{}{
			"test": "test",
		},
		Meta: &cookbook.Meta{
			Links: &cookbook.Links{
				Next: "next",
				Prev: "prev",
			},
		},
	}, resp)
}

func (j *JSendSuite) Test_Stringify_Success() {
	expected := `{"status":"success","data":{"test":"test"}}`

	result := cookbook.SuccessDataResponse(map[string]interface{}{
		"test": "test",
	}, nil).Stringify()

	assert.JSONEq(j.T(), expected, result)
}

func (j *JSendSuite) Test_Stringify_Failed_Marshal() {
	ch := make(chan string)
	expected := "json: unsupported type: chan string"

	result := cookbook.SuccessDataResponse(ch, nil).Stringify()

	assert.Equal(j.T(), expected, result)
}
