package interceptors

import (
	"beyond/pkg/codex"
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ClientErrorInterceptor 把gRPC的错误码转为自定义的业务错误码
func ClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			// 把err转化为grpcStatus
			grpcStatus, _ := status.FromError(err)
			// 把grpcStatus转为Codex
			xc := codex.GrpcStatusToCodeX(grpcStatus)
			//
			err = errors.WithMessage(xc, grpcStatus.Message())
		}

		return err
	}
}
