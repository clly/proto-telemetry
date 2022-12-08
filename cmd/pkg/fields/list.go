package fields

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

type ListGenerator struct {
	f *protogen.Field
}

func NewListGenerator(f *protogen.Field) *ListGenerator {
	if f.Desc.IsList() {
		return nil
	}
	return &ListGenerator{
		f: f,
	}
}

func (l *ListGenerator) Generate(g *protogen.GeneratedFile) {
	if l == nil {
		return
	}
	n := fmt.Sprintf("%s.%s_length", strings.ToLower(l.f.Parent.GoIdent.GoName), strings.ToLower(l.f.GoName))
	g.P(`attribute.Int64(`, n, `, len(x.Get`, l.f.GoName, `(),`)
}
