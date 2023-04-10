package grpc_message_marker

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/shoenig/test/must"
	"github.com/shoenig/test/portal"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/clly/proto-telemetry/examples/ping"
	pingv1 "github.com/clly/proto-telemetry/examples/ping/proto/gen/ping/v1"
)

func Test_UnaryInterceptor(t *testing.T) {
	testcases := map[string]struct {
		withoutRequest  bool
		withoutResponse bool
	}{
		"Normal": {},
		"WithoutRequest": {
			withoutRequest: true,
		},
		"WithoutResponse": {
			withoutResponse: true,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			grabber := portal.New(t)
			port := grabber.One()
			l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			must.NoError(t, err)

			ctx := context.Background()

			closer, exporter, err := tracer()
			must.NoError(t, err)

			opts := []InterceptorOpt{}
			if tc.withoutRequest {
				opts = append(opts, WithoutRequest())
			}
			if tc.withoutResponse {
				opts = append(opts, WithoutResponse())
			}

			s := grpc.NewServer(
				grpc.ChainUnaryInterceptor(
					otelgrpc.UnaryServerInterceptor(),
					UnaryInterceptor(opts...),
				),
			)
			pingsvr := &ping.PingServer{}
			pingv1.RegisterPingServiceServer(s, pingsvr)
			go func() {
				s.Serve(l)
			}()

			conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
			must.NoError(t, err)

			client := pingv1.NewPingServiceClient(conn)
			_, err = client.Ping(ctx, &pingv1.PingRequest{Name: "me"})
			must.NoError(t, err)

			spans := exporter.GetSpans()
			must.Eq(t, 1, len(spans.Snapshots()))
			for _, kv := range spans.Snapshots()[0].Attributes() {
				if kv.Key == "pingrequest.name" && !tc.withoutRequest {
					must.Eq(t, attribute.STRING, kv.Value.Type())
					must.Eq(t, "me", kv.Value.AsString())
				}
				if kv.Key == "pingrequest.name" && tc.withoutRequest {
					t.Fatalf("expected not to contain %s", kv.Key)
				}
				if kv.Key == "pingresponse.name" && !tc.withoutResponse {
					must.Eq(t, attribute.STRING, kv.Value.Type())
					must.Eq(t, "me", kv.Value.AsString())
				}
				if kv.Key == "pingresponse.name" && tc.withoutResponse {
					t.Fatalf("expected not to contain %s", kv.Key)
				}
			}

			closer()
			l.Close()
		})
	}
}

func tracer() (func(), *tracetest.InMemoryExporter, error) {
	exporter := tracetest.NewInMemoryExporter()
	ssp := sdktrace.NewSimpleSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(ssp))
	shutdown := func() { _ = tp.Shutdown(context.Background()) }
	otel.SetTracerProvider(tp)

	// set global propagator to baggage (the default is no-op).
	otel.SetTextMapPropagator(propagation.Baggage{})
	return shutdown, exporter, nil
}
