package ping

import (
	"context"

	pingv1pb "github.com/clly/proto-telemetry/examples/ping/proto/gen/ping/v1"
)

// PingServer implements the PingServerService grpc server.
type PingServer struct{}

func (s *PingServer) Ping(ctx context.Context, req *pingv1pb.PingRequest) (*pingv1pb.PingResponse, error) {
	return &pingv1pb.PingResponse{
		Response: req.Name,
	}, nil
}
