package generators

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
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
