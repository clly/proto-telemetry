package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc"

	messagemarker "github.com/clly/proto-telemetry/interceptor/grpc/messagemarker"

	otechov1 "github.com/clly/proto-telemetry/examples/example-otel/gen/proto/go/echo/v1"
	"github.com/clly/proto-telemetry/examples/example-otel/tracing"
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

	shutdown, err := tracing.Init()
	if err != nil {
		return err
	}
	defer shutdown()

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
			messagemarker.UnaryServerInterceptor(),
		),
	)
	otechov1.RegisterEchoServiceServer(server, &svr{})
	log.Println("listening on", listener.Addr())
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to server grpc server: %w", err)
	}

	return nil
}

type svr struct {
	otechov1.UnimplementedEchoServiceServer
}

func (s *svr) Echo(ctx context.Context, req *otechov1.EchoRequest) (*otechov1.EchoResponse, error) {
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/server").Start(ctx, "Echo")
	defer span.End()
	req.TraceAttributes(ctx)
	return &otechov1.EchoResponse{
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
