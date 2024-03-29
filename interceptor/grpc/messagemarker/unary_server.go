package messagemarker

import (
	"context"

	"google.golang.org/grpc"
)

type interceptorOpt struct {
	withoutRequestAttributes  bool
	withoutResponseAttributes bool
	requestOpts               requestOpts
	responseOpts              responseOpts
}

type requestOpts struct {
	name        string
	withoutName bool
}

type responseOpts struct {
	name        string
	withoutName bool
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

func WithRequestName(name string) InterceptorOpt {
	return func(opt *interceptorOpt) {
		opt.requestOpts.name = name
	}
}

type attributer interface {
	TraceAttributes(ctx context.Context)
}

type namedAttributer interface {
	TraceNamedAttributes(ctx context.Context, pfx string)
}

func UnaryServerInterceptor(opts ...InterceptorOpt) grpc.UnaryServerInterceptor {
	iOpts := &interceptorOpt{
		requestOpts:  requestOpts{name: "req"},
		responseOpts: responseOpts{name: "resp"},
	}

	for _, opt := range opts {
		opt(iOpts)
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		traceReq(ctx, iOpts, req)
		resp, err = handler(ctx, req)
		if err != nil {
			return resp, err
		}

		traceResp(ctx, iOpts, resp)

		return resp, err
	}
}

func UnaryClientInterceptor(opts ...InterceptorOpt) grpc.UnaryClientInterceptor {
	iOpts := &interceptorOpt{
		requestOpts:  requestOpts{name: "req"},
		responseOpts: responseOpts{name: "resp"},
	}

	for _, opt := range opts {
		opt(iOpts)
	}

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		traceReq(ctx, iOpts, req)
		if err := invoker(ctx, method, req, reply, cc, opts...); err != nil {
			return err
		}

		traceResp(ctx, iOpts, reply)

		return nil
	}

}
func traceReq(ctx context.Context, iopts *interceptorOpt, req interface{}) {
	if iopts.withoutRequestAttributes {
		return
	}

	if iopts.requestOpts.withoutName {
		if attributer, ok := req.(attributer); ok {
			attributer.TraceAttributes(ctx)
		}
		return
	}

	if namedAttributer, ok := req.(namedAttributer); ok {
		namedAttributer.TraceNamedAttributes(ctx, iopts.requestOpts.name)
	}
}

func traceResp(ctx context.Context, iopts *interceptorOpt, resp interface{}) {
	if iopts.withoutResponseAttributes {
		return
	}

	if iopts.responseOpts.withoutName {
		if attributer, ok := resp.(attributer); ok {
			attributer.TraceAttributes(ctx)
		}
		return
	}

	if namedAttributer, ok := resp.(namedAttributer); ok {
		namedAttributer.TraceNamedAttributes(ctx, iopts.responseOpts.name)
	}
}
