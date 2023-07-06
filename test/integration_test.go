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
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/clly/proto-telemetry/cmd/pkg/options"
	octracing "github.com/clly/proto-telemetry/examples/example-oc/tracing"
	oteltracing "github.com/clly/proto-telemetry/examples/example-otel/tracing"
	optionsv1 "github.com/clly/proto-telemetry/proto/telemetry/options/v1"
	ottestv1 "github.com/clly/proto-telemetry/test/open-telemetry/proto/gen/test/v1"
	octestv1 "github.com/clly/proto-telemetry/test/opencensus/proto/gen/test/v1"
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

type TraceNamedAttributer interface {
	TraceNamedAttributes(ctx context.Context, pfx string)
}

type TestTraceAttributer interface {
	ProtoReflect() protoreflect.Message
	TraceAttributer
	TraceNamedAttributer
}

func otTrace(t *testing.T, ta TraceAttributer) {
	ctx := context.Background()
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/server").Start(ctx, t.Name())
	ta.TraceAttributes(ctx)
	span.End()
}

// func otNamedTrace(t *testing.T, na TraceNamedAttributer, pfx string) {
// 	ctx := context.Background()
// 	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/server").Start(ctx, t.Name())
// 	na.TraceNamedAttributes(ctx, pfx)
// 	span.End()
// }

func ocTrace(t *testing.T, ta TraceAttributer) {
	ctx := context.Background()
	ctx, span := trace.StartSpan(ctx, t.Name())
	ta.TraceAttributes(ctx)
	span.End()
}

func (s *IntegrationSuite) SetupTest() {

	shutdown, exporter, err := oteltracing.TestInit()
	require.NoError(s.T(), err)

	s.openTelemetryShutdown = shutdown
	s.openTelemetryExporter = exporter

	ocexporter, ocshutdown := octracing.TestInit()

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(ocexporter)

	s.ocexporter = ocexporter
	s.ocexporterShutdown = ocshutdown

}

func (s *IntegrationSuite) AfterTest() {
	s.ocexporterShutdown()
	s.NoError(s.openTelemetryExporter.Shutdown(context.Background()))
}

func (s *IntegrationSuite) AfterSuite() {
	s.openTelemetryShutdown()
	s.ocexporterShutdown()
}

func (s *IntegrationSuite) Test_ExcludedFile() {
	msg := &ottestv1.TestMessage{
		TestField: uuid.New().String(),
	}

	_, taOK := any(msg).(TraceAttributer)
	_, naOK := any(msg).(TraceNamedAttributer)
	require.False(s.T(), taOK, "message should not implement TraceAttributer")
	require.False(s.T(), naOK, "message should not implement TraceNamedAttributer")
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
		"RenameMessage": {
			msg: &ottestv1.RenameMessagePrefix{
				Msg: uuid.New().String(),
			},
		},
		"NameField": {
			msg: &ottestv1.NameField{
				Msg: uuid.New().String(),
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

			// attributes turns opentelemetry key/values to a map of strings to any
			attributes := func() map[string]any {
				attrs := make(map[string]any)
				for _, attr := range testspan.Attributes() {
					attrs[string(attr.Key)] = attr.Value.AsInterface()
				}
				return attrs
			}
			verify(t, msg, attributes())
		})
	}
}

func verify(t *testing.T, msg TestTraceAttributer, attributes map[string]any) {
	type value struct {
		excluded bool
		value    interface{}
		typ      protoreflect.Kind
		isMap    bool
		realMap  any
	}

	fields := make(map[string]value)

	msg.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		name := options.GetTelemetryFieldName(protodesc.ToFieldDescriptorProto(fd), string(fd.Name()))
		val := value{
			excluded: options.GetTelemetryFieldExclude(protodesc.ToFieldDescriptorProto(fd), false),
			value:    v.Interface(),
			typ:      fd.Kind(),
			isMap:    fd.IsMap(),
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
	pfx = options.GetTelemetryMessageName(protodesc.ToDescriptorProto(msg.ProtoReflect().Descriptor()), pfx)

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

	for key, val := range attributes {
		expVal, ok := expectedAttributes[key]
		if !ok {
			spew.Dump(expectedAttributes)
			require.Fail(t, "unknown key traced", key)
		}
		require.EqualValues(t, expVal, val, "key %s does not match %v!=%v", key, expVal, val)
	}

	for name, value := range expectedAttributes {
		m := attributes
		val, ok := m[name]
		if !ok {
			spew.Dump(m)
			require.Fail(t, "missing key traced", name)
		}
		require.EqualValues(t, value, val, "key %s does not match %v!=%v", name, value, val)
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

func (s *IntegrationSuite) Test_IntegrationOpenCensus() {
	// t := s.T()

	testcases := map[string]struct {
		msg       TestTraceAttributer
		reference map[string]string
	}{
		"StringMessage": {
			msg: &octestv1.StringMessage{Msg: uuid.New().String()},
		},
		"Int32Message": {
			msg: &octestv1.Int32Message{Num32: 5},
		},
		"Uint32Message": {
			msg: &octestv1.Uint32Message{Unum32: 10},
		},
		"Int64Message": {
			msg: &octestv1.Int64Message{Num64: 15},
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
			msg: &octestv1.MapMessage{
				Meta: map[string]string{
					"a": uuid.New().String(),
					"b": uuid.New().String(),
					"c": uuid.New().String(),
				},
			},
		},
		"ExcludeField": {
			msg: &octestv1.ExcludeField{
				Name:      uuid.New().String(),
				NonMasked: uuid.New().String(),
			},
		},
		"ExcludeMessage": {
			msg: &octestv1.ExcludeMessage{
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
		"RenameMessage": {
			msg: &octestv1.RenameMessagePrefix{
				Msg: uuid.New().String(),
			},
		},
		"NameField": {
			msg: &octestv1.NameField{
				Msg: uuid.New().String(),
			},
		},
	}

	for name, tc := range testcases {
		s.Run(name, func() {
			t := s.T()
			defer s.ocexporterShutdown()
			msg := tc.msg
			ocTrace(t, msg)

			snapshots := s.ocexporter.SpanData
			require.Len(t, snapshots, 1)
			testspan := snapshots[0]

			var excludeMessage bool
			e := proto.GetExtension(msg.ProtoReflect().Descriptor().Options(), optionsv1.E_ExcludeMessage)
			if s, ok := e.(bool); ok {
				excludeMessage = s
			}
			if excludeMessage {
				require.Len(t, testspan.Attributes, 0)
				return
			}

			verify(t, msg, testspan.Attributes)
		})
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
