package postgres

import (
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

// IsDuplicate check error from postgres if error is because duplicated record
func IsDuplicate(err error) bool {
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
		return true
	}
	return false
}

// IsForeignNotFound check error from postgres if error is because foreign key not found
func IsForeignNotFound(err error) bool {
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23503" {
		return true
	}
	return false
}

// IsInvalidInput check error from postgres if error is because invalid input enumeration
func IsInvalidInput(err error) bool {
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "22P02" {
		return true
	} else if ok && err.Code == "23502" {
		return true
	}
	return false
}

// IsNotFound check error from postgres if error is because record not found
func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
