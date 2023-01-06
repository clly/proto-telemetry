package generators

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type generator interface {
	Generate(g *protogen.GeneratedFile)
}

type Generator[T generator] struct {
	typeGen T
}

type MapGenerator struct {
	m *protogen.Field
}

func MessageGenerator(m *protogen.Message) {

}

func NewFieldGenerator(f *protogen.Field) FieldAttribute {
	return newField(f)
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

func NewMapGenerator(m *protogen.Field) *MapGenerator {
	return &MapGenerator{
		m: m,
	}
}

func (m *MapGenerator) Generate(g *protogen.GeneratedFile) {
	g.P("for m, v := range x.Get", m.m.GoName, "() {")
	g.P("span.SetAttributes(")
	// g.P(`attribute.String("`, strings.ToLower(m.m.Parent.GoIdent.GoName), `.`, strings.ToLower(m.m.GoName), `_%s, x.Msg),`)
	g.P(`attribute.String(fmt.Sprintf("`, strings.ToLower(m.m.Parent.GoIdent.GoName), `.`, strings.ToLower(m.m.GoName), `_%s", m), v),`)
	// g.P("attribute.String(", strings.ToLower(m.m.Parent.GoIdent.GoName), ", v)")
	g.P(")")
	g.P("}")
}
