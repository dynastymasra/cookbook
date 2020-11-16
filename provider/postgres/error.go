package postgres

import (
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

func IsDuplicate(err error) bool {
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
		return true
	}
	return false
}

func IsForeignNotFound(err error) bool {
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23503" {
		return true
	}
	return false
}

func IsInvalidInput(err error) bool {
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "22P02" {
		return true
	}
	return false
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
