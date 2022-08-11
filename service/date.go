package service

import (
	"time"
)
///
func CheckDateValid(d1, d2 time.Time) bool {
	if d1.Before(d2) {
		return true
	}
	return false
}
