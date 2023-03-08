package generators

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FileGenerator struct {
	Telemetry TelemetryBackend
	g         *protogen.GeneratedFile
}

type TelemetryBackend interface {
	Imports(g *protogen.GeneratedFile)
	Span() string
	AttributeType(k protoreflect.Kind) string
	Attribute() string
}

func NewFileGenerator(g *protogen.GeneratedFile, gen TelemetryBackend) *FileGenerator {
	return &FileGenerator{
		Telemetry: gen,
		g:         g,
	}
}

func (f *FileGenerator) Generate() {
	f.Telemetry.Imports(f.g)
}

type OpentelemetryGenerator struct {
	attributeIdent string
	traceIdent     string
	ctxIdent       string
}

func (o *OpentelemetryGenerator) Imports(g *protogen.GeneratedFile) {
	g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "attribute",
		GoImportPath: "go.opentelemetry.io/otel/attribute",
	})
	o.attributeIdent = "attribute"

	g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "trace",
		GoImportPath: "go.opentelemetry.io/otel/trace",
	})
	o.traceIdent = "trace"

	o.ctxIdent = g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "context",
	})
}

func (o *OpentelemetryGenerator) Span() string {
	return fmt.Sprintf("span := %s.SpanFromContext(ctx)", o.traceIdent)
}

func (o *OpentelemetryGenerator) Attribute() string {
	return "span.SetAttributes"
}

func (o *OpentelemetryGenerator) AttributeType(k protoreflect.Kind) string {
	switch k {
	case protoreflect.StringKind:
		return fmt.Sprintf("%s.String", o.attributeIdent)
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return fmt.Sprintf("%s.Int64", o.attributeIdent)
	case protoreflect.BoolKind:
		return fmt.Sprintf("%s.Bool", o.attributeIdent)
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return fmt.Sprintf("%s.Float64", o.attributeIdent)
	default:
		return ""
	}
}

type OpencensusGenerator struct {
	traceIdent     string
	attributeIdent string
	ctxIdent       string
}

func (o *OpencensusGenerator) Imports(g *protogen.GeneratedFile) {
	o.traceIdent = g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "trace",
		GoImportPath: "go.opencensus.io/trace",
	})
	o.ctxIdent = g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "context",
	})

	o.attributeIdent = "trace"

}

func (o *OpencensusGenerator) Span() string {
	return "span := trace.FromContext(ctx)"
}

func (o *OpencensusGenerator) AttributeType(k protoreflect.Kind) string {
	switch k {
	case protoreflect.StringKind:
		return fmt.Sprintf("%s.StringAttribute", o.attributeIdent)
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return fmt.Sprintf("%s.Int64Attribute", o.attributeIdent)
	case protoreflect.BoolKind:
		return fmt.Sprintf("%s.BoolAttribute", o.attributeIdent)
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return fmt.Sprintf("%s.Float64Attribute", o.attributeIdent)
	default:
		return ""
	}
}

func (o *OpencensusGenerator) Attribute() string {
	return "span.AddAttributes"
}
