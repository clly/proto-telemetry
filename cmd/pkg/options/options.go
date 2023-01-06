package options

import (
	optionsv1 "github.com/clly/proto-telemetry/proto/telemetry/options/v1"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/proto"
)

// TelemetryPackage returns the (gopherjs.gopherjs_package) option if
// specified, or an empty string if it was not.
func GetTelemetryPackage(file *descriptor.FileDescriptorProto) string {
	if file == nil || file.Options == nil {
		return ""
	}

	e := proto.GetExtension(file.Options, optionsv1.E_TelemetryPackage)

	if s, ok := e.(*string); ok {
		return *s
	}

	return ""
}

func GetTelemetryFieldExclude(field *descriptor.FieldDescriptorProto, defaultValue bool) bool {
	if field == nil || field.Options == nil {
		return defaultValue
	}

	e := proto.GetExtension(field.Options, optionsv1.E_Exclude)

	if s, ok := e.(*bool); ok {
		return *s
	}

	return defaultValue
}
