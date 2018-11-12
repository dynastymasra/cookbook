package cookbook_test

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/dynastymasra/cookbook"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type EnvironmentSuite struct {
	suite.Suite
}

func Test_EnvironmentSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentSuite))
}

func (e *EnvironmentSuite) Test_StringEnv_Success() {
	viper.Set("STRING_KEY_SUCCESS", "string")
	result, err := cookbook.StringEnv("STRING_KEY_SUCCESS")

	assert.NotEmpty(e.T(), result)
	assert.Equal(e.T(), "string", result)
	assert.Nil(e.T(), err)
}

func (e *EnvironmentSuite) Test_StringEnv_Failed() {
	result, err := cookbook.StringEnv("STRING_KEY_FAILED")

	assert.Empty(e.T(), result)
	assert.NotNil(e.T(), err)
}

func (e *EnvironmentSuite) Test_BoolEnv_Success() {
	viper.Set("BOOL_KEY_SUCCESS", true)
	result, err := cookbook.BoolEnv("BOOL_KEY_SUCCESS")

	assert.NotEmpty(e.T(), result)
	assert.Equal(e.T(), true, result)
	assert.Nil(e.T(), err)
}

func (e *EnvironmentSuite) Test_BoolEnv_Failed() {
	result, err := cookbook.BoolEnv("BOOL_KEY_FAILED")

	assert.Empty(e.T(), result)
	assert.Equal(e.T(), false, result)
	assert.NotNil(e.T(), err)
}

func (e *EnvironmentSuite) Test_IntEnv_Success() {
	viper.Set("INT_KEY_SUCCESS", 10)
	result, err := cookbook.IntEnv("INT_KEY_SUCCESS")

	assert.NotEmpty(e.T(), result)
	assert.Equal(e.T(), 10, result)
	assert.Nil(e.T(), err)
}

func (e *EnvironmentSuite) Test_IntEnv_Failed() {
	result, err := cookbook.IntEnv("INT_KEY_FAILED")

	assert.Empty(e.T(), result)
	assert.NotNil(e.T(), err)
}
