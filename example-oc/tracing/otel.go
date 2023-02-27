package tracing

import (
	"context"
	"io"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func Init() (func(), error) {
	exporter, err := newExporter(os.Stdout)
	if err != nil {
		return nil, err
	}

	ssp := sdktrace.NewSimpleSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(ssp))
	shutdown := func() { _ = tp.Shutdown(context.Background()) }
	otel.SetTracerProvider(tp)

	// set global propagator to baggage (the default is no-op).
	otel.SetTextMapPropagator(propagation.Baggage{})
	return shutdown, nil
}

// newExporter returns a console exporter.
func newExporter(w io.Writer) (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

func TestInit() (func(), *tracetest.InMemoryExporter, error) {
	exporter := tracetest.NewInMemoryExporter()
	ssp := sdktrace.NewSimpleSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(ssp))
	shutdown := func() { _ = tp.Shutdown(context.Background()) }
	otel.SetTracerProvider(tp)

	// set global propagator to baggage (the default is no-op).
	otel.SetTextMapPropagator(propagation.Baggage{})
	return shutdown, exporter, nil
}
