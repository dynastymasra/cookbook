package mongo

import "go.mongodb.org/mongo-driver/mongo"

const (
	ErrDuplicateCode = 11000
)

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

func IsInvalidIndexValue(err error) bool {
	return err == mongo.ErrInvalidIndexValue
}
