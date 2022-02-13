package utils

import (
	"time"
)

func CurrentDateTime() (time.Time, error) {
	return time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
}

