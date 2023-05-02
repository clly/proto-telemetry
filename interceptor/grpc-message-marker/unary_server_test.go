package grpc_message_marker

import (
	"context"
	"fmt"
	"net"
	"strings"
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
		withoutReqName  bool
		reqName         string
	}{
		"Normal": {},
		"WithoutRequest": {
			withoutRequest: true,
		},
		"WithoutResponse": {
			withoutResponse: true,
		},
		"WithRequestName": {
			reqName: "reqmsg",
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
			reqName := "req"
			if tc.reqName != "" {
				reqName = tc.reqName
				opts = append(opts, WithRequestName(tc.reqName))
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
			reqAttr := &attribute.KeyValue{
				Key:   attribute.Key(reqName + ".pingrequest.name"),
				Value: attribute.StringValue("me"),
			}
			respAttr := &attribute.KeyValue{
				Key:   "resp.pingresponse.response",
				Value: attribute.StringValue("me"),
			}
			if tc.withoutRequest {
				must.NotContains[*attribute.KeyValue](t, reqAttr, containsAttr(spans.Snapshots()[0].Attributes()),
					must.Sprintf("expected no element %v=%v, all elements: %s", reqAttr.Key, reqAttr.Value.AsString(),
						allKVs(spans.Snapshots()[0].Attributes())),
				)
			} else {
				must.Contains[*attribute.KeyValue](t, reqAttr, containsAttr(spans.Snapshots()[0].Attributes()),
					must.Sprintf("expected element %v=%v, all elements: %s", reqAttr.Key, reqAttr.Value.AsString(),
						allKVs(spans.Snapshots()[0].Attributes())),
				)
			}
			if tc.withoutResponse {
				must.NotContains[*attribute.KeyValue](t, respAttr, containsAttr(spans.Snapshots()[0].Attributes()),
					must.Sprintf("expected no element %v=%v, all elements: %s", respAttr.Key, respAttr.Value.AsString(),
						allKVs(spans.Snapshots()[0].Attributes())),
				)
			} else {
				must.Contains[*attribute.KeyValue](t, respAttr, containsAttr(spans.Snapshots()[0].Attributes()),
					must.Sprintf("expected element %v=%v, all elements: %s", respAttr.Key, respAttr.Value.AsString(),
						allKVs(spans.Snapshots()[0].Attributes())),
				)
			}

			closer()
			l.Close()
		})
	}
}

type ContainsFunc[T any] func(val T) bool

func (c ContainsFunc[T]) Contains(val T) bool {
	return c(val)
}

func containsAttr(vals []attribute.KeyValue) ContainsFunc[*attribute.KeyValue] {
	return func(val *attribute.KeyValue) bool {
		for _, v := range vals {
			if v.Key == val.Key && v.Value == val.Value {
				return true
			}
		}
		return false
	}
}

func allKVs(kvs []attribute.KeyValue) string {
	sb := strings.Builder{}
	for _, kv := range kvs {
		sb.WriteString(fmt.Sprintf("%s=%v\n", kv.Key, kv.Value.AsString()))
	}
	return sb.String()
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
