package generators

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type RepeatedGenerator struct {
	r *protogen.Field
}

func NewRepeatedGenerator(r *protogen.Field) generator {
	return &RepeatedGenerator{
		r: r,
	}
}

func (r *RepeatedGenerator) Generate(f *FileGenerator, named bool) {
	g := f.g
	var key = fmt.Sprintf("%s.%s", strings.ToLower(r.r.Parent.GoIdent.GoName), strings.ToLower(r.r.GoName))
	if named {
		key = fmt.Sprintf("pfx.%s", key)
	}

	// g.P(`attribute.String("`, strings.ToLower(m.m.Parent.GoIdent.GoName), `.`, strings.ToLower(m.m.GoName), `_%s, x.Msg),`)
	g.P(f.Telemetry.AttributeType(protoreflect.StringKind), `("`, key, `", x.Get`, r.r.GoName, `()[0]),`)
	// g.P("attribute.String(", strings.ToLower(m.m.Parent.GoIdent.GoName), ", v)")
	// g.P(")")
}
