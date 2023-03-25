package test

import (
	"context"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.opencensus.io/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/clly/proto-telemetry/examples/example-oc/gen/proto/go/ocecho/v1"
	octracing "github.com/clly/proto-telemetry/examples/example-oc/tracing"
	otelechov1 "github.com/clly/proto-telemetry/examples/example-otel/gen/proto/go/otecho/v1"
	oteltracing "github.com/clly/proto-telemetry/examples/example-otel/tracing"
	optionsv1 "github.com/clly/proto-telemetry/proto/telemetry/options/v1"
	ottestv1 "github.com/clly/proto-telemetry/test/open-telemetry/proto/gen/test/v1"
)

type IntegrationSuite struct {
	suite.Suite
	openTelemetryExporter *tracetest.InMemoryExporter
	openTelemetryShutdown func()

	ocexporter         *octracing.InMemoryExporter
	ocexporterShutdown func()
}

func (s *IntegrationSuite) SetupSuite() {
	wd, err := os.Getwd()
	require.NoError(s.T(), err)

	cmd := exec.Cmd{
		Path: "./generate.bash",
		Dir:  wd,
	}

	output, err := cmd.CombinedOutput()
	require.NoError(s.T(), err, string(output))
}

type TraceAttributer interface {
	TraceAttributes(ctx context.Context)
}

type NamedAttributer interface {
	NamedAttributes(ctx context.Context, pfx string)
}

type TestTraceAttributer interface {
	ProtoReflect() protoreflect.Message
	TraceAttributer
	NamedAttributer
}

func otTrace(t *testing.T, ta TraceAttributer) {
	ctx := context.Background()
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/server").Start(ctx, t.Name())
	ta.TraceAttributes(ctx)
	span.End()
}

func otNamedTrace(t *testing.T, na NamedAttributer, pfx string) {
	ctx := context.Background()
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/server").Start(ctx, t.Name())
	na.NamedAttributes(ctx, pfx)
	span.End()
}

func (s *IntegrationSuite) SetupTest() {

	shutdown, exporter, err := oteltracing.TestInit()
	require.NoError(s.T(), err)

	s.openTelemetryShutdown = shutdown
	s.openTelemetryExporter = exporter

	ocexporter, ocshutdown := octracing.TestInit()
	s.ocexporter = ocexporter
	s.ocexporterShutdown = ocshutdown
}

func (s *IntegrationSuite) AfterTest() {
	s.ocexporterShutdown()
	s.openTelemetryExporter.Shutdown(context.Background())
}

func (s *IntegrationSuite) AfterSuite() {
	s.openTelemetryShutdown()
	s.ocexporterShutdown()
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

func (s *IntegrationSuite) Test_OpenTelemetry() {

	testcases := map[string]struct {
		msg       TestTraceAttributer
		reference map[string]string
	}{
		"StringMessage": {
			msg: &ottestv1.StringMessage{Msg: uuid.New().String()},
		},
		"Int32Message": {
			msg: &ottestv1.Int32Message{Num32: 5},
		},
		"Uint32Message": {
			msg: &ottestv1.Uint32Message{Unum32: 10},
		},
		"Int64Message": {
			msg: &ottestv1.Int64Message{Num64: 15},
		},
		// submessages aren't supported and we can't test them yet
		// "SubMessage": {
		// 	msg: &ottestv1.SubMessage{
		// 		Details: &ottestv1.MessageDetails{
		// 			Details: uuid.New().String(),
		// 		},
		// 		Envelope: &ottestv1.SubMessage_Envelope{
		// 			Name: uuid.New().String(),
		// 		},
		// 	},
		// },
		"MapMessage": {
			msg: &ottestv1.MapMessage{
				Meta: map[string]string{
					"a": uuid.New().String(),
					"b": uuid.New().String(),
					"c": uuid.New().String(),
				},
			},
		},
		"ExcludeField": {
			msg: &ottestv1.ExcludeField{
				Name:      uuid.New().String(),
				NonMasked: uuid.New().String(),
			},
		},
		"ExcludeMessage": {
			msg: &ottestv1.ExcludeMessage{
				Msg: uuid.New().String(),
				Now: &datetime.DateTime{
					Year:    int32(time.Now().Year()),
					Month:   int32(time.Now().Month()),
					Day:     int32(time.Now().Day()),
					Hours:   int32(time.Now().Hour()),
					Minutes: int32(time.Now().Minute()),
					Seconds: int32(time.Now().Second()),
					Nanos:   int32(time.Now().Nanosecond()),
				},
			},
		},
	}

	for name, tc := range testcases {
		s.Run(name, func() {
			t := s.T()
			defer s.openTelemetryExporter.Reset()
			msg := tc.msg
			otTrace(t, msg)

			snapshots := s.openTelemetryExporter.GetSpans().Snapshots()
			require.Len(t, snapshots, 1)
			testspan := snapshots[0]

			var excludeMessage bool
			e := proto.GetExtension(msg.ProtoReflect().Descriptor().Options(), optionsv1.E_ExcludeMessage)
			if s, ok := e.(bool); ok {
				excludeMessage = s
			}
			if excludeMessage {
				require.Len(t, testspan.Attributes(), 0)
				return
			}

			verify(t, msg, testspan)
		})
	}
}

func verify(t *testing.T, msg TestTraceAttributer, testspan sdktrace.ReadOnlySpan) {
	type value struct {
		excluded  bool
		fieldName string
		value     interface{}
		typ       protoreflect.Kind
		isMap     bool
		realMap   any
	}

	fields := make(map[string]value)

	msg.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		e := proto.GetExtension(fd.Options(), optionsv1.E_Exclude)

		val := value{
			excluded:  false,
			fieldName: string(fd.Name()),
			value:     v.Interface(),
			typ:       fd.Kind(),
			isMap:     fd.IsMap(),
		}
		if s, ok := e.(bool); ok {
			val.excluded = s
		} else {
			val.excluded = false
		}

		name := val.fieldName

		if s, ok := e.(bool); ok {
			val.excluded = s
		} else {
			val.excluded = false
		}

		if val.isMap {
			switch fd.MapKey().Kind() {
			case protoreflect.StringKind:
				realMap := make(map[string]interface{})
				v.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
					realMap[key.String()] = value.Interface()
					return true
				})
				val.realMap = realMap
			case protoreflect.BoolKind:
				realMap := make(map[bool]interface{})
				v.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
					realMap[key.Bool()] = value.Interface()
					return true
				})
			case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
				protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
				realMap := make(map[int64]interface{})
				v.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
					realMap[key.Int()] = value.Interface()
					return true
				})
				val.realMap = realMap
			case protoreflect.Uint64Kind, protoreflect.Fixed64Kind, protoreflect.Uint32Kind,
				protoreflect.Fixed32Kind:
				realMap := make(map[uint64]interface{})
				v.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
					realMap[key.Uint()] = value.Interface()
					return true
				})
				val.realMap = realMap
			default:
				panic("unknown map key type")
			}
		}

		fields[name] = val
		return true
	})

	val := reflect.ValueOf(msg).Elem()
	pfx := strings.ToLower(val.Type().Name())

	expectedAttributes := make(map[string]any)
	for name, val := range fields {
		if val.isMap {
			mapIter(pfx+"."+name, expectedAttributes, reflect.ValueOf(val.realMap))
			continue
		}
		if !val.excluded {
			expectedAttributes[pfx+"."+name] = val.value
		}
	}

	for _, kv := range testspan.Attributes() {
		val, ok := expectedAttributes[string(kv.Key)]
		if !ok {
			spew.Dump(expectedAttributes)
			require.Fail(t, "unknown key traced", kv.Key)
		}
		require.EqualValues(t, val, kv.Value.AsInterface())
	}

	for name, value := range expectedAttributes {
		contains := Contains(testspan.Attributes(), func(keyvalue attribute.KeyValue) bool {
			if string(keyvalue.Key) == name {
				require.EqualValues(t, value, keyvalue.Value.AsInterface())
				return true
			}
			return false
		})
		require.True(t, contains, "attributes do not contain key %s \nattributes: %s", name, testspan.Attributes())
	}
}

func Contains[T any](slice []T, f func(T) bool) bool {
	for _, val := range slice {
		contains := f(val)
		if contains {
			return true
		}
	}
	return false
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
				m[pfx+".number"] = val.Field(i).Interface()
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

func (s *IntegrationSuite) Test_ExcludeMessage() {
	t := s.T()
	req := &otelechov1.MessageDetails{}

	shutdown, exporter, err := oteltracing.TestInit()
	defer shutdown()
	require.NoError(t, err)

	ctx := context.Background()
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/server").Start(ctx, "message details")
	req.TraceAttributes(ctx)
	span.End()

	snapshots := exporter.GetSpans().Snapshots()
	require.Len(t, snapshots, 1)
	testspan := snapshots[0]
	require.Len(t, testspan.Attributes(), 0)

	ctx = context.Background()
	ctx, span = otel.Tracer("protoc-gen-go-telemetry/example/server").Start(ctx, "named message details")
	req.NamedAttributes(ctx, "prefix")
	span.End()

	snapshots = exporter.GetSpans().Snapshots()
	require.Len(t, snapshots, 2)
	testspan = snapshots[1]
	require.Len(t, testspan.Attributes(), 0)
}

func (s *IntegrationSuite) Test_IntegrationOpenCensus() {
	t := s.T()
	req := echoRequest()

	exporter, shutdown := octracing.TestInit()
	defer shutdown()

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
				m[pfx+".number"] = val.Field(i).Interface()
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
