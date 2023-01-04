package infra

import (
	"context"
	"time"
)

func CreateContext(timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	ctxParent := context.Background()
	ctx, cancel = context.WithTimeout(ctxParent, timeout)
	return
}
