package generators

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type generator interface {
	Generate(f *FileGenerator, named bool)
}

type Generator[T generator] struct {
	typeGen T
}

type MapGenerator struct {
	m *protogen.Field
}

func NewMapGenerator(m *protogen.Field) generator {
	return &MapGenerator{
		m: m,
	}
}

func (m *MapGenerator) Generate(f *FileGenerator, named bool) {
	g := f.g
	var key = fmt.Sprintf("%s.%s", strings.ToLower(m.m.Parent.GoIdent.GoName), strings.ToLower(m.m.GoName))
	if named {
		key = fmt.Sprintf("pfx.%s", key)
	}

	g.P("for m, v := range x.Get", m.m.GoName, "() {")
	g.P(f.Telemetry.Attribute(), "(")
	// g.P(`attribute.String("`, strings.ToLower(m.m.Parent.GoIdent.GoName), `.`, strings.ToLower(m.m.GoName), `_%s, x.Msg),`)
	g.P(f.Telemetry.AttributeType(protoreflect.StringKind), `(fmt.Sprintf("`, key, `_%s", m), v),`)
	// g.P("attribute.String(", strings.ToLower(m.m.Parent.GoIdent.GoName), ", v)")
	g.P(")")
	g.P("}")
}
