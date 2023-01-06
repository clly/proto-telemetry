package generators

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type FileGenerator[T telemetryGenerator] struct {
	telemetry T
}

type telemetryGenerator interface {
	*opencensusGenerator | *opentelemetryGenerator
	generator
}

func NewFileGenerator(g *protogen.GeneratedFile) *FileGenerator[*opentelemetryGenerator] {
	return &FileGenerator[*opentelemetryGenerator]{
		telemetry: &opentelemetryGenerator{},
	}
}

func (f *FileGenerator[T]) Generate(g *protogen.GeneratedFile) {
	f.telemetry.Generate(g)
}

type opentelemetryGenerator struct {
	attributeIdent string
	traceIdent     string
	ctxIdent       string
}

func (o *opentelemetryGenerator) Generate(g *protogen.GeneratedFile) {
	o.attributeIdent = g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "attribute",
		GoImportPath: "go.opentelemetry.io/otel/attribute",
	})

	o.traceIdent = g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "trace",
		GoImportPath: "go.opentelemetry.io/otel/trace",
	})

	o.ctxIdent = g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "context",
	})
}

type opencensusGenerator struct{}

func (o *opencensusGenerator) Generate(g *protogen.GeneratedFile) {

}
