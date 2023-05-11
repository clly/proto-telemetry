package main

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"

	"github.com/clly/proto-telemetry/examples/example-otel/gen/proto/go/otecho/v1"
	"github.com/clly/proto-telemetry/examples/example-otel/tracing"
	messagemarker "github.com/clly/proto-telemetry/interceptor/grpc/messagemarker"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(otelgrpc.UnaryClientInterceptor(), messagemarker.UnaryClientInterceptor()))
	if err != nil {
		return fmt.Errorf("failed to connect to ocecho service at %s: %w", connectTo, err)
	}
	log.Println("connected to", connectTo)

	shutdown, err := tracing.Init()
	if err != nil {
		return err
	}
	defer shutdown()

	echo := otechov1.NewEchoServiceClient(conn)

	ctx := context.Background()
	ctx, span := otel.Tracer("protoc-gen-go-telemetry/example/client").Start(ctx, "Echo Client")
	defer span.End()

	req := &otechov1.EchoRequest{
		Msg: "Hello World!",
	}
	req.TraceAttributes(ctx)
	if _, err := echo.Echo(context.Background(), &otechov1.EchoRequest{
		Msg: "Hello World!",
	}); err != nil {
		return fmt.Errorf("failed to Echo: %w", err)
	}

	log.Println("Successfully echoed")
	return nil
}
