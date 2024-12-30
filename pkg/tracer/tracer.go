package tracer

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"vehicle-registration-manager/pkg/logger"
)

const tracerIdKey = "tracer_id"

type Tracer struct {
	ctx    context.Context
	Logger *logger.Logger
	ID     string
}

func NewTracer(request *http.Request) *Tracer {
	id := getTracerIDFromRequest(request)
	logg := logger.NewLoggerWithTrace(id)
	return &Tracer{ctx: request.Context(), ID: id, Logger: logg}
}

func getTracerIDFromRequest(request *http.Request) string {
	tracerId := request.Header.Get(tracerIdKey)
	if tracerId != "" {
		return tracerId
	}
	ctx := context.Background()
	if request.Context() != nil {
		ctx = request.Context()
	}
	return getTracerID(ctx)
}

func getTracerID(ctx context.Context) string {
	tracerID, exists := ctx.Value(tracerIdKey).(string)
	if exists {
		return tracerID
	}

	tracerID = uuid.New().String()
	ctx = context.WithValue(ctx, tracerIdKey, tracerID)
	return tracerID
}

// FakeTracer
func NewFakeTracer() *Tracer {
	id := "fake-tracer-id"
	logg := logger.NewLoggerWithTrace(id)
	return &Tracer{ctx: context.Background(), ID: id, Logger: logg}
}
