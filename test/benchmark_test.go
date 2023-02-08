package test

import (
	"context"
	"testing"

	examplepb "github.com/clly/proto-telemetry/example/gen/proto/go/echo/v1"
	"github.com/google/uuid"
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
