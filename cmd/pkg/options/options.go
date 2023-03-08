package options

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/proto"

	optionsv1 "github.com/clly/proto-telemetry/proto/telemetry/options/v1"
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

	if s, ok := e.(bool); ok {
		return s
	}

	return defaultValue
}

func GetTelemetryFieldName(field *descriptor.FieldDescriptorProto, defaultValue string) string {
	if field == nil || field.Options == nil {
		return defaultValue
	}

	e := proto.GetExtension(field.Options, optionsv1.E_FieldName)

	if s, ok := e.(string); ok {
		return s
	}

	return defaultValue
}

func GetTelemetryMessageExclude(message *descriptor.DescriptorProto, defaultValue bool) bool {
	if message == nil || message.Options == nil {
		return defaultValue
	}

	e := proto.GetExtension(message.Options, optionsv1.E_ExcludeMessage)

	if s, ok := e.(bool); ok {
		return s
	}

	return defaultValue
}
