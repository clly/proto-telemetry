package grpc_message_marker

import (
	"context"

	"google.golang.org/grpc"
)

type interceptorOpt struct {
	withoutRequestAttributes  bool
	withoutResponseAttributes bool
}

type InterceptorOpt func(opt *interceptorOpt)

func WithoutRequest() InterceptorOpt {
	return func(opt *interceptorOpt) {
		opt.withoutRequestAttributes = true
	}
}

func WithoutResponse() InterceptorOpt {
	return func(opt *interceptorOpt) {
		opt.withoutResponseAttributes = true
	}
}

type attributer interface {
	TraceAttributes(ctx context.Context)
}

type namedAttributer interface {
	NamedAttributes(ctx context.Context, pfx string)
}

type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error)

func UnaryInterceptor(opts ...InterceptorOpt) grpc.UnaryServerInterceptor {
	iOpts := &interceptorOpt{}

	for _, opt := range opts {
		opt(iOpts)
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		traceReq(ctx, iOpts, req)
		resp, err = handler(ctx, req)
		if err != nil {
			return resp, err
		}

		if !iOpts.withoutResponseAttributes {
			if attributer, ok := resp.(interface {
				TraceAttributes(context.Context)
			}); ok {
				attributer.TraceAttributes(ctx)
			}
		}

		return resp, err
	}
}

func traceReq(ctx context.Context, iopts *interceptorOpt, req interface{}) {
	if iopts.withoutRequestAttributes {
		return
	}
	if attributer, ok := req.(interface {
		TraceAttributes(context.Context)
	}); ok {
		attributer.TraceAttributes(ctx)
	}

}
