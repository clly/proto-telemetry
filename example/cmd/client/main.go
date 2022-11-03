package main

import (
	"context"
	"fmt"
	"log"

	echov1 "github.com/clly/protoc-telemetry-go/example/gen/proto/go/v1"
	"google.golang.org/grpc"
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
		return fmt.Errorf("failed to connect to echo service at %s: %w", connectTo, err)
	}
	log.Println("connected to", connectTo)

	echo := echov1.NewEchoServiceClient(conn)

	if _, err := echo.Echo(context.Background(), &echov1.EchoRequest{
		Msg: "Hello World!",
	}); err != nil {
		return fmt.Errorf("failed to Echo: %w", err)
	}

	log.Println("Successfully echoed")
	return nil
}
