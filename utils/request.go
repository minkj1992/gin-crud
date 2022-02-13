package utils

import (
	"context"
	"time"
)

const (
	CONTEXT_TIMEOUT = 100*time.Second
)

func GetContextWithTimeOut() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
}