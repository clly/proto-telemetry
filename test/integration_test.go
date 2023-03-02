package test

import (
	"context"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.opencensus.io/trace"
	"go.opentelemetry.io/otel"

	ocechov1 "github.com/clly/proto-telemetry/example-oc/gen/proto/go/ocecho/v1"
	octracing "github.com/clly/proto-telemetry/example-oc/tracing"
	otelechov1 "github.com/clly/proto-telemetry/example-otel/gen/proto/go/otecho/v1"
	oteltracing "github.com/clly/proto-telemetry/example-otel/tracing"
)

type IntegrationSuite struct {
	suite.Suite
}

func (s *IntegrationSuite) SetupTest() {
	wd, err := os.Getwd()
	require.NoError(s.T(), err)

	cmd := exec.Cmd{
		Path: "run-dev.sh",
		Dir:  wd + "/../",
	}

	output, err := cmd.CombinedOutput()
	require.NoError(s.T(), err, string(output))
}

func echoRequest() *otelechov1.EchoRequest {
	return &otelechov1.EchoRequest{
		Msg:    "msg",
		Num32:  1,
		Unum32: 2,
		Num64:  3,
		Details: &otelechov1.MessageDetails{
			Details: "signed-by: bob",
		},
		Meta: map[string]string{
			"reply-to": "hello@gmail.com",
		},
		Sender: "cindy",
	}
}

func ocEchoRequest() *ocechov1.EchoRequest {
	return &ocechov1.EchoRequest{
		Msg:    "msg",
		Num32:  1,
		Unum32: 2,
		Num64:  3,
		Details: &ocechov1.MessageDetails{
			Details: "signed-by: bob",
		},
		Meta: map[string]string{
			"reply-to": "hello@gmail.com",
		},
		Sender: "cindy",
	}
}

func (s *IntegrationSuite) Test_IntegrationOpenTelemetryWithMap() {
	t := s.T()
	req := ocEchoRequest()

	shutdown, exporter, err := oteltracing.TestInit()
	defer shutdown()
	require.NoError(t, err)

	ctx := context.Background()
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/server").Start(ctx, "Echo")
	req.TraceAttributes(ctx)
	span.End()

	snapshots := exporter.GetSpans().Snapshots()
	require.Len(t, snapshots, 1)
	testspan := snapshots[0]
	val := reflect.ValueOf(req).Elem()

	pfx := strings.ToLower(val.Type().Name())
	m := map[string]any{}
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Type().Field(i).IsExported() {
			name := strings.ToLower(val.Type().Field(i).Name)
			if val.Field(i).Type().Kind() == reflect.Map {
				mapIter(pfx+"."+name, m, val.Field(i))
				continue
			}
			// num32 is testing custom field_name
			if name == "num32" {
				m["number"] = val.Field(i).Interface()
				continue
			}
			m[pfx+"."+name] = val.Field(i).Interface()
		}
	}

	for _, kv := range testspan.Attributes() {
		val, ok := m[string(kv.Key)]
		if !ok {
			require.Fail(t, "unknown key traced", kv.Key)
		}
		require.EqualValues(t, val, kv.Value.AsInterface())
	}
}

func (s *IntegrationSuite) Test_IntegrationOpenCensus() {
	t := s.T()
	req := echoRequest()

	exporter := octracing.TestInit()

	ctx := context.Background()
	ctx, span := trace.StartSpan(ctx, "Echo")
	req.TraceAttributes(ctx)
	span.End()

	spanData := exporter.SpanData
	require.Len(t, spanData, 1)
	testspan := spanData[0]
	val := reflect.ValueOf(req).Elem()

	pfx := strings.ToLower(val.Type().Name())
	m := map[string]any{}
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Type().Field(i).IsExported() {
			name := strings.ToLower(val.Type().Field(i).Name)
			if val.Field(i).Type().Kind() == reflect.Map {
				mapIter(pfx+"."+name, m, val.Field(i))
				continue
			}
			// num32 is testing custom field_name
			if name == "num32" {
				m["number"] = val.Field(i).Interface()
				continue
			}
			m[pfx+"."+name] = val.Field(i).Interface()
		}
	}

	for key, attrVal := range testspan.Attributes {
		val, ok := m[string(key)]
		if !ok {
			require.Fail(t, "unknown key traced", key)
		}
		require.EqualValues(t, val, attrVal)
	}
}

func mapIter(pfx string, m map[string]any, val reflect.Value) {
	iter := val.MapRange()
	for iter.Next() {
		m[pfx+"_"+iter.Key().String()] = iter.Value().Interface()
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func Test_TelemetryIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (t *IntegrationSuite) TestTrue() {
	t.Suite.True(true)
}
