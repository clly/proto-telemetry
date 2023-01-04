// Code generated by protoco-telemetry-go. DO NOT EDIT.

package echov1

import (
	context "context"
	fmt "fmt"
	attribute "go.opentelemetry.io/otel/attribute"
	trace "go.opentelemetry.io/otel/trace"
)

func (x *MessageDetails) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("messagedetails.details", x.Details),
	)
}

func (x *EchoRequest_Envelope) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("echorequest.envelope.name", x.Name),
	)
}

func (x *EchoResponse) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("echoresponse.msg", x.Msg),
	)
}

func (x *Foo) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes()
}

func (x *EchoRequest) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("echorequest.msg", x.Msg),
		attribute.Int64("echorequest.num32", int64(x.Num32)),
		attribute.Int64("echorequest.unum32", int64(x.Unum32)),
		attribute.Int64("echorequest.num64", int64(x.Num64)),
		attribute.String("echorequest.sender", x.Sender),
	)
	for m, v := range x.GetMeta() {
		span.SetAttributes(
			attribute.String(fmt.Sprintf("echorequest.meta_%s", m), v),
		)
	}
}
