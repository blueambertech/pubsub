package pubsub

import (
	"context"
	"time"
)

type Handler interface {
	Subscribe(cancCtx context.Context, subID string, refreshRate time.Duration, msgProc func(c context.Context, msgData []byte))
	Push(ctx context.Context, topicID, msg string) error
}
