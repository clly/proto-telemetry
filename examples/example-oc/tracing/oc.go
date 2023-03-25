package tracing

import (
	"log"
	"sync"

	"go.opencensus.io/trace"
)

type consoleExporter struct{}
type InMemoryExporter struct {
	SpanData []*trace.SpanData
	m        sync.Mutex
}

func (i *InMemoryExporter) shutdown() {
	i.m.Lock()
	defer i.m.Unlock()
	i.SpanData = make([]*trace.SpanData, 0)
}

// Compile time assertion that the exporter implements trace.Exporter
var _ trace.Exporter = (*consoleExporter)(nil)
var _ trace.Exporter = (*InMemoryExporter)(nil)

func (cse *consoleExporter) ExportSpan(sd *trace.SpanData) {
	log.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n",
		sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}

func (ime *InMemoryExporter) ExportSpan(sd *trace.SpanData) {
	ime.m.Lock()
	defer ime.m.Unlock()

	ime.SpanData = append(ime.SpanData, sd)
}

func Init() {
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(new(consoleExporter))
}

func TestInit() (*InMemoryExporter, func()) {
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	exporter := &InMemoryExporter{
		SpanData: make([]*trace.SpanData, 0),
		m:        sync.Mutex{},
	}

	trace.RegisterExporter(exporter)
	return exporter, exporter.shutdown
}
