package grpc_message_marker

import (
	"context"

	"google.golang.org/grpc"
)

type interceptorOpt struct{}

type InterceptorOpt func(opt interceptorOpt)

type attributer interface {
	TraceAttributes(ctx context.Context)
}

type namedAttributer interface {
	NamedAttributes(ctx context.Context, pfx string)
}

type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error)

func UnaryInterceptor(opts ...InterceptorOpt) grpc.UnaryServerInterceptor {
	iOpts := interceptorOpt{}

	for _, opt := range opts {
		opt(iOpts)
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if attributer, ok := req.(interface {
			TraceAttributes(context.Context)
		}); ok {
			attributer.TraceAttributes(ctx)
		}

		resp, err = handler(ctx, req)
		return resp, err
	}
}
