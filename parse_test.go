package cookbook_test

import (
	"testing"

	"github.com/dynastymasra/cookbook"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParseSuite struct {
	suite.Suite
}

func Test_ParseSuit(t *testing.T) {
	suite.Run(t, new(ParseSuite))
}

func (p *ParseSuite) Test_ParseStringToInt_Success() {
	result := cookbook.ParseStringToInt("10", 0)

	assert.NotZero(p.T(), result)
	assert.Equal(p.T(), 10, result)
}

func (p *ParseSuite) Test_ParseStringToInt_Failed_ReturnDefaultValue() {
	result := cookbook.ParseStringToInt("", 10)

	assert.NotZero(p.T(), result)
	assert.Equal(p.T(), 10, result)
}

func (p *ParseSuite) Test_ParseStringToFloat64_Success() {
	result := cookbook.ParseStringToFloat64("5.25", 10.5)

	assert.NotZero(p.T(), result)
	assert.Equal(p.T(), 5.25, result)
}

func (p *ParseSuite) Test_ParseStringToFloat64_Failed_ReturnDefaultValue() {
	result := cookbook.ParseStringToFloat64("", 10.5)

	assert.NotZero(p.T(), result)
	assert.Equal(p.T(), 10.5, result)
}

func (p *ParseSuite) Test_Stringify_Success() {
	str := "Test parse interface"

	result := cookbook.Stringify(str)

	assert.NotEmpty(p.T(), result)
}

func (p *ParseSuite) Test_Stringify_Failed_Marshal() {
	ch := make(chan string)

	result := cookbook.Stringify(ch)

	assert.NotEmpty(p.T(), result)
}

func (p *ParseSuite) Test_ParseStringNil() {
	s := "test"

	result := cookbook.ParseStringNil(s)

	assert.NotNil(p.T(), result)
}

func (p *ParseSuite) Test_ParseStringNil_Empty() {
	s := ""

	result := cookbook.ParseStringNil(s)

	assert.Nil(p.T(), result)
}

func (p *ParseSuite) Test_ParsePtrString() {
	s := "test"

	result := cookbook.ParsePtrString(&s)

	assert.NotEmpty(p.T(), result)
	assert.Equal(p.T(), s, result)
}

func (p *ParseSuite) Test_ParsePtrString_Nil() {
	result := cookbook.ParsePtrString(nil)

	assert.Empty(p.T(), result)
}
