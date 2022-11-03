package main

import (
	"fmt"
	"log"
	"net"

	echov1 "github.com/clly/protoc-telemetry-go/example/gen/proto/go/v1"
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

	server := grpc.NewServer()
	echov1.RegisterEchoServiceServer(server, &svr{})
	log.Println("listening on", listener.Addr())
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to server grpc server: %w", err)
	}

	return nil
}

type svr struct {
	echov1.UnimplementedEchoServiceServer
}

// func (s *svr) Echo(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
// 	return nil, nil
// }
