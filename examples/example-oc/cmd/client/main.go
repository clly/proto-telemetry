package main

import (
	"context"
	"fmt"
	"log"

	ocechov1 "github.com/clly/proto-telemetry/examples/example-oc/gen/proto/go/ocecho/v1"
	"github.com/clly/proto-telemetry/examples/example-oc/tracing"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to ocecho service at %s: %w", connectTo, err)
	}
	log.Println("connected to", connectTo)

	tracing.Init()

	echo := ocechov1.NewEchoServiceClient(conn)

	ctx := context.Background()
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/client").Start(ctx, "Echo Client")
	defer span.End()

	req := &ocechov1.EchoRequest{
		Msg: "Hello World!",
	}
	req.TraceAttributes(ctx)
	if _, err := echo.Echo(context.Background(), &ocechov1.EchoRequest{
		Msg: "Hello World!",
	}); err != nil {
		return fmt.Errorf("failed to Echo: %w", err)
	}

	log.Println("Successfully echoed")
	return nil
}
