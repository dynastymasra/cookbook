package postgres

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func IsDuplicate(err error) bool {
	if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
		return true
	}
	return false
}

func IsForeignNotFound(err error) bool {
	if err, ok := err.(*pq.Error); ok && err.Code == "23503" {
		return true
	}
	return false
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
