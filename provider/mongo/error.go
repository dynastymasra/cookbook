package mongo

import "go.mongodb.org/mongo-driver/mongo"

const (
	ErrDuplicateCode = 11000
)

// IsDuplicate check error from mongo if error because duplicated record
func IsDuplicate(err error) bool {
	mErr, ok := err.(mongo.WriteException)
	if !ok {
		return false
	}

	if mErr.WriteErrors[0].Code == ErrDuplicateCode {
		return true
	}

	return false
}

// IsInvalidIndexValue check error from mongo if error because invalid index value
func IsInvalidIndexValue(err error) bool {
	return err == mongo.ErrInvalidIndexValue
}
