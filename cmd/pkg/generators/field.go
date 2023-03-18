package generators

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/clly/proto-telemetry/cmd/pkg/options"
)

type FieldAttribute struct {
	field              *protogen.Field
	goName             string
	attrName           string
	attrKind           string
	castCall           string
	isTrailer          bool
	g                  generator
	telemetryGenerator TelemetryBackend
}

func NewFieldGenerator(f *protogen.Field, t TelemetryBackend) FieldAttribute {
	return newField(f, t)
}

func (f *FieldAttribute) IsTrailer() bool {
	return f.isTrailer
}

func attributeFromKind(t TelemetryBackend, k protoreflect.Kind) (string, string) {
	attribute := t.AttributeType(k)
	switch k {
	case protoreflect.BoolKind:
		return attribute, ""
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.DoubleKind, protoreflect.FloatKind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind:
		return attribute, "int64"
	case protoreflect.StringKind:
		return attribute, ""
	default:
		return "", ""
	}
}

type attrName struct {
	parent    string
	fieldName string
	separator string
}

func (a attrName) String() string {
	b := strings.Builder{}
	b.WriteString(a.parent)
	b.WriteString(a.separator)
	b.WriteString(a.fieldName)
	// replace _ with .
	str := strings.ToLower(b.String())
	str = strings.ReplaceAll(str, "_", ".")
	// replace . with whatever the user wants the separator to be
	return strings.ReplaceAll(str, ".", a.separator)
}

func newField(field *protogen.Field, t TelemetryBackend) FieldAttribute {
	name := attrName{
		parent:    field.Parent.GoIdent.GoName,
		fieldName: field.GoName,
		separator: ".",
	}
	attrKind, castCall := attributeFromKind(t, field.Desc.Kind())

	name.parent = options.GetTelemetryMessageName(protodesc.ToDescriptorProto(field.Parent.Desc), name.parent)
	name.fieldName = options.GetTelemetryFieldName(protodesc.ToFieldDescriptorProto(field.Desc), name.fieldName)

	fa := FieldAttribute{
		attrName: name.String(),
		attrKind: attrKind,
		castCall: castCall,
		goName:   field.GoName,
		field:    field,
	}

	if field.Desc.IsMap() {
		fa.isTrailer = true
		mg := NewMapGenerator(field)
		fa.g = mg
	}

	return fa
}

func (f *FieldAttribute) Generate(g *protogen.GeneratedFile, named bool) {
	if f.attrKind == "" {
		// fmt.Fprintln(os.Stderr, "Kind", f.field.Desc.Kind().GoString(), "of type", f.field.GoIdent.GoName, "in", f.field.Parent.GoIdent.GoName, "is unsupported")
		return
	}

	if options.GetTelemetryFieldExclude(protodesc.ToFieldDescriptorProto(f.field.Desc), false) {
		return
	}

	var key = fmt.Sprintf(`"%s"`, f.attrName)
	if named {
		key = fmt.Sprintf(`pfx+".%s"`, f.attrName)
	}
	var s string
	if f.castCall == "" {
		s = fmt.Sprintf(`%s(%s, x.%s),`, f.attrKind, key, f.goName)
	} else {

		s = fmt.Sprintf(`%s(%s, %s(x.%s)),`, f.attrKind, key, f.castCall, f.goName)
	}

	g.P(s)
}
