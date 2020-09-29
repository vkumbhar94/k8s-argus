package worker

import (
	"context"
	"errors"
	"time"

	lmjaeger "github.com/logicmonitor/k8s-argus/pkg/jaeger"
)

// RLTokenizer tokenizer which contains regulated token generation
type RLTokenizer struct {
	ch     <-chan interface{}
	ctx    context.Context
	cancel context.CancelFunc
}

// NewRLTokenizer creates new tokenizer with mentioned limit and starts it implicitly
func NewRLTokenizer(limit int) *RLTokenizer {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan interface{}, limit)
	go func(wch chan<- interface{}) {
		ticker := time.NewTicker(time.Duration(60000/limit) * time.Millisecond)
		for {
			select {
			case <-ctx.Done():
				close(wch)
				return
			case <-ticker.C:
				wch <- 1
			}
		}
	}(ch)
	return &RLTokenizer{
		ch:     ch,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (rlt *RLTokenizer) popToken(span lmjaeger.LMSpan) error {
	span.LogKV("event", "poping token")
	if rlt.ctx.Err() != nil {
		span.LogKV("event", "context is already canceled")
		span.SetTag("parent-context", "cancelled")
		return rlt.ctx.Err()
	}
	select {
	case <-rlt.ch:
		span.LogKV("event", "token poped")
		return nil
	case <-time.After(1 * time.Minute):
		span.LogKV("event", "token pop timout")
		span.SetTag("timeout", "1m")
		span.SetTag("error", true)
		span.LogKV("event", "timed out")
		return errors.New("new token did not received in 1 minute, reconfigure tokenizer")
	}
}

func (rlt *RLTokenizer) stop() {
	rlt.cancel()
}
