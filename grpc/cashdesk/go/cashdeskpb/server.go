package cashdeskpb

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryServerMetadata() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if err := checkMetadata(ctx); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func StreamServerMetadata() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := checkMetadata(ss.Context()); err != nil {
			return err
		}

		return handler(srv, ss)
	}
}
