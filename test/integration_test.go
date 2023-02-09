package test

import (
	"context"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	echov1 "github.com/clly/proto-telemetry/example/gen/proto/go/echo/v1"
	"github.com/clly/proto-telemetry/example/tracing"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
)

func Test_IntegrationWithoutMap(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	cmd := exec.Cmd{
		Path: "run-dev.sh",
		Dir:  wd + "/../",
	}
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, string(output))

	req := &echov1.EchoRequest{
		Msg:    "msg",
		Num32:  1,
		Unum32: 2,
		Num64:  3,
		Details: &echov1.MessageDetails{
			Details: "signed-by: bob",
		},
		Meta: map[string]string{
			"reply-to": "hello@gmail.com",
		},
		Sender: "cindy",
	}

	shutdown, exporter, err := tracing.TestInit()
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

func mapIter(pfx string, m map[string]any, val reflect.Value) {
	iter := val.MapRange()
	for iter.Next() {
		m[pfx+"_"+iter.Key().String()] = iter.Value().Interface()
	}
}
