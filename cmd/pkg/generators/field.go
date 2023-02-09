package generators

import (
	"fmt"
	"strings"

	"github.com/clly/proto-telemetry/cmd/pkg/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldAttribute struct {
	field     *protogen.Field
	goName    string
	attrName  string
	attrKind  string
	castCall  string
	isTrailer bool
	g         generator
}

func NewFieldGenerator(f *protogen.Field) FieldAttribute {
	return newField(f)
}

func (f *FieldAttribute) IsTrailer() bool {
	return f.isTrailer
}

func attributeFromKind(k protoreflect.Kind) (string, string) {
	switch k {
	case protoreflect.BoolKind:
		return "Bool", ""
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.DoubleKind, protoreflect.FloatKind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind:
		return "Int64", "int64"
	case protoreflect.StringKind:
		return "String", ""
	default:
		return "", ""
	}
}

func newField(field *protogen.Field) FieldAttribute {
	attrName := strings.ReplaceAll(field.GoIdent.GoName, "_", ".")
	attrName = strings.ToLower(attrName)
	attrKind, castCall := attributeFromKind(field.Desc.Kind())

	attrName = options.GetTelemetryFieldName(protodesc.ToFieldDescriptorProto(field.Desc), attrName)

	fa := FieldAttribute{
		attrName: attrName,
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
		s = fmt.Sprintf(`attribute.%s(%s, x.%s),`, f.attrKind, key, f.goName)
	} else {

		s = fmt.Sprintf(`attribute.%s(%s, %s(x.%s)),`, f.attrKind, key, f.castCall, f.goName)
	}

	g.P(s)
}
