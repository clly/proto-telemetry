package main

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"

	echov1 "github.com/clly/proto-telemetry/example-otel/gen/proto/go/otecho/v1"
	"github.com/clly/proto-telemetry/example-otel/tracing"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to ocecho service at %s: %w", connectTo, err)
	}
	log.Println("connected to", connectTo)

	shutdown, err := tracing.Init()
	if err != nil {
		return err
	}
	defer shutdown()

	echo := echov1.NewEchoServiceClient(conn)

	ctx := context.Background()
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/client").Start(ctx, "Echo Client")
	defer span.End()

	req := &echov1.EchoRequest{
		Msg: "Hello World!",
	}
	req.TraceAttributes(ctx)
	if _, err := echo.Echo(context.Background(), &echov1.EchoRequest{
		Msg: "Hello World!",
	}); err != nil {
		return fmt.Errorf("failed to Echo: %w", err)
	}

	log.Println("Successfully echoed")
	return nil
}
