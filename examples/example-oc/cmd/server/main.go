package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	ocechov1 "github.com/clly/proto-telemetry/examples/example-oc/gen/proto/go/ocecho/v1"
	"github.com/clly/proto-telemetry/examples/example-oc/tracing"

	"go.opencensus.io/trace"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listen := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listen, err)
	}

	tracing.Init()

	server := grpc.NewServer()
	ocechov1.RegisterEchoServiceServer(server, &svr{})
	log.Println("listening on", listener.Addr())
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to server grpc server: %w", err)
	}

	return nil
}

type svr struct {
	ocechov1.UnimplementedEchoServiceServer
}

func (s *svr) Echo(ctx context.Context, req *ocechov1.EchoRequest) (*ocechov1.EchoResponse, error) {
	ctx, span := trace.StartSpan(ctx, "Echo")
	defer span.End()
	req.TraceAttributes(ctx)
	return &ocechov1.EchoResponse{
		Msg: req.Msg,
		Now: dtFromTime(time.Now()),
	}, nil
}

func dtFromTime(t time.Time) *datetime.DateTime {
	return &datetime.DateTime{
		Year:    int32(t.Year()),
		Month:   int32(t.Month()),
		Day:     int32(t.Day()),
		Hours:   int32(t.Hour()),
		Minutes: int32(t.Minute()),
		Seconds: int32(t.Second()),
		Nanos:   int32(t.Nanosecond()),
	}
}
