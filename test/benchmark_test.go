package test

import (
	"context"
	"testing"

	"github.com/google/uuid"

	examplepb "github.com/clly/proto-telemetry/examples/example-otel/gen/proto/go/otecho/v1"
)

func BenchmarkTraceAttributes(b *testing.B) {
	var msg = generateMessage()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		msg.TraceAttributes(ctx)
	}
}

func BenchmarkNamedAttributes(b *testing.B) {
	var msg = generateMessage()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		msg.NamedAttributes(ctx, "resource")
	}
}

func generateMessage() *examplepb.EchoRequest {
	var msg = &examplepb.EchoRequest{
		Msg:    "",
		Num32:  0,
		Unum32: 0,
		Num64:  0,
		Details: &examplepb.MessageDetails{
			Details: "",
		},
		Meta: map[string]string{
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
			uuid.New().String(): uuid.New().String(),
		},
		Sender: "",
	}

	return msg
}
