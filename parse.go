package cookbook

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nyaruka/phonenumbers"
)

// ParseStringToInt Parse string to int, error return default value
func ParseStringToInt(str string, def int) int {
	parse, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return parse
}

// ParseStringToFloat64 Parse string to float64, error return default value
func ParseStringToFloat64(str string, def float64) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return def
	}

	return f
}

// Stringify Convert interface to string, error will return string error message
func Stringify(str interface{}) string {
	out, err := json.Marshal(str)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

// ParseStringNil function to check string if empty return nil
func ParseStringNil(s string) *string {
	if len(s) > 0 {
		return &s
	}
	return nil
}

// ParsePtrString function check pointer string, if nil return empty string
func ParsePtrString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ParseStringToPhone Function to parse string to formatted phone number
// See details documentation in https://github.com/nyaruka/phonenumbers
// defaultRegionCode using ISO 3166-Alpha 2
func ParseStringToPhone(phone, defaultRegionCode string, phoneFormat phonenumbers.PhoneNumberFormat) (string, error) {
	num, err := phonenumbers.Parse(phone, defaultRegionCode)
	if err != nil {
		return "", err
	}

	if !phonenumbers.IsValidNumber(num) {
		return "", fmt.Errorf("phone number %s is invalid phone for region %s %s", phone, defaultRegionCode, phonenumbers.GetRegionCodeForNumber(num))
	}

	formattedNum := phonenumbers.Format(num, phoneFormat)

	return formattedNum, nil
}

// ParsePtrGRPCTimeToPtrTime Parse pointer timestamp from GRPC struct and return nil if empty or invalid timestamp
func ParsePtrGRPCTimeToPtrTime(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}

	ts := t.AsTime()

	if ts.IsZero() {
		return nil
	}

	return &ts
}

// ParsePtrGRPCTimeToTime parse pointer timestamp from GRPC and return new timestamp if nil or error
func ParsePtrGRPCTimeToTime(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.Now().UTC()
	}

	ts := t.AsTime()

	if ts.IsZero() {
		return time.Now().UTC()
	}

	return ts
}
