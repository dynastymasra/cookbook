package cookbook_test

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/nyaruka/phonenumbers"

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

func (p *ParseSuite) Test_ParseStringToPhone_Success() {

	res1, err1 := cookbook.ParseStringToPhone("+64 3 345 6789", "", phonenumbers.E164)
	res2, err2 := cookbook.ParseStringToPhone("+64 21 345 687", "", phonenumbers.E164)
	res3, err3 := cookbook.ParseStringToPhone("+64 021 345 687", "", phonenumbers.E164)
	res4, err4 := cookbook.ParseStringToPhone("+64 021-345-687", "", phonenumbers.E164)
	res5, err5 := cookbook.ParseStringToPhone("+14155552671", "", phonenumbers.E164)
	res6, err6 := cookbook.ParseStringToPhone("+442071838750", "", phonenumbers.E164)

	assert.Equal(p.T(), "+6433456789", res1)
	assert.Equal(p.T(), "+6421345687", res2)
	assert.Equal(p.T(), "+6421345687", res3)
	assert.Equal(p.T(), "+6421345687", res4)
	assert.Equal(p.T(), "+14155552671", res5)
	assert.Equal(p.T(), "+442071838750", res6)
	assert.NoError(p.T(), err1)
	assert.NoError(p.T(), err2)
	assert.NoError(p.T(), err3)
	assert.NoError(p.T(), err4)
	assert.NoError(p.T(), err5)
	assert.NoError(p.T(), err6)
}

func (p *ParseSuite) Test_ParseStringToPhone_Failed() {
	res1, err1 := cookbook.ParseStringToPhone("03 345 6789", "", phonenumbers.E164)
	res2, err2 := cookbook.ParseStringToPhone("021 345 678", "", phonenumbers.E164)
	res3, err3 := cookbook.ParseStringToPhone("64 03 345 6789", "", phonenumbers.E164)
	res4, err4 := cookbook.ParseStringToPhone("64 21 345 678", "", phonenumbers.E164)
	res5, err5 := cookbook.ParseStringToPhone("+644155552671", "", phonenumbers.E164)
	res6, err6 := cookbook.ParseStringToPhone("+622071838750", "", phonenumbers.E164)

	assert.Equal(p.T(), "", res1)
	assert.Equal(p.T(), "", res2)
	assert.Equal(p.T(), "", res3)
	assert.Equal(p.T(), "", res4)
	assert.Equal(p.T(), "", res5)
	assert.Equal(p.T(), "", res6)
	assert.Error(p.T(), err1)
	assert.Error(p.T(), err2)
	assert.Error(p.T(), err3)
	assert.Error(p.T(), err4)
	assert.Error(p.T(), err5)
	assert.Error(p.T(), err6)
}

func (p *ParseSuite) Test_ParsePtrGRPCTimeToPtrTime_Nil() {
	result := cookbook.ParsePtrGRPCTimeToPtrTime(nil)

	assert.Nil(p.T(), result)
}

func (p *ParseSuite) Test_ParsePtrGRPCTimeToPtrTime_Zero() {
	t := &timestamp.Timestamp{Seconds: -62135596800, Nanos: 0}
	result := cookbook.ParsePtrGRPCTimeToPtrTime(t)

	assert.Nil(p.T(), result)
}

func (p *ParseSuite) Test_ParsePtrGRPCTimeToPtrTime() {
	t := ptypes.TimestampNow()
	result := cookbook.ParsePtrGRPCTimeToPtrTime(t)

	assert.NotNil(p.T(), result)
}

func (p *ParseSuite) Test_ParsePtrGRPCTimeToTime_Nil() {
	ts := time.Now().UTC()
	result := cookbook.ParsePtrGRPCTimeToTime(nil)

	assert.NotNil(p.T(), result)
	assert.Equal(p.T(), ts.Year(), result.Year())
	assert.Equal(p.T(), ts.Month(), result.Month())
	assert.Equal(p.T(), ts.Day(), result.Day())
	assert.Equal(p.T(), ts.Hour(), result.Hour())
	assert.Equal(p.T(), ts.Minute(), result.Minute())
}

func (p *ParseSuite) Test_ParsePtrGRPCTimeToTime_Zero() {
	ts := time.Now().UTC()
	t := &timestamp.Timestamp{Seconds: -62135596800, Nanos: 0}

	result := cookbook.ParsePtrGRPCTimeToTime(t)

	assert.NotNil(p.T(), result)
	assert.Equal(p.T(), ts.Year(), result.Year())
	assert.Equal(p.T(), ts.Month(), result.Month())
	assert.Equal(p.T(), ts.Day(), result.Day())
	assert.Equal(p.T(), ts.Hour(), result.Hour())
	assert.Equal(p.T(), ts.Minute(), result.Minute())
}

func (p *ParseSuite) Test_ParsePtrGRPCTimeToTime() {
	t := ptypes.TimestampNow()
	ts, _ := ptypes.Timestamp(t)
	result := cookbook.ParsePtrGRPCTimeToTime(t)

	assert.NotNil(p.T(), result)
	assert.Equal(p.T(), ts.Year(), result.Year())
	assert.Equal(p.T(), ts.Month(), result.Month())
	assert.Equal(p.T(), ts.Day(), result.Day())
	assert.Equal(p.T(), ts.Hour(), result.Hour())
	assert.Equal(p.T(), ts.Minute(), result.Minute())
	assert.Equal(p.T(), ts.Second(), result.Second())
}
