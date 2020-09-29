package utilities

import (
	"context"

	"github.com/logicmonitor/k8s-argus/pkg/connection"
	lmjaeger "github.com/logicmonitor/k8s-argus/pkg/jaeger"
	"github.com/logicmonitor/k8s-argus/pkg/lmctx"
	"github.com/logicmonitor/k8s-collectorset-controller/api"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
)

// GetCollectorID - get collectorID from csc
func GetCollectorID(lctx *lmctx.LMContext) int32 {
	span := lmjaeger.Span(lctx)
	ctx := context.Background()
	if span != nil {
		log.Debugf("bundled span in context")
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	log.Debugf("parent grpc span context %v", opentracing.SpanFromContext(ctx))
	reply, err := connection.GetCSCClient().CollectorID(ctx, &api.CollectorIDRequest{})
	if err != nil || reply == nil {
		log.Errorf("Failed to get collector ID: %v", err)
		return 0
	}

	return reply.Id
}
