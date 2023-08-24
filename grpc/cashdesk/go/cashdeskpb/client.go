package cashdeskpb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ClientMetadata(workstationID string) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithUnaryInterceptor(metadataUnaryClient(workstationID)),
		grpc.WithStreamInterceptor(metadataStreamClient(workstationID)),
	}
}

func metadataUnaryClient(workstationID string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(
			metadata.AppendToOutgoingContext(ctx, metadataWorkstationID, workstationID),
			method,
			req,
			reply,
			cc,
			opts...,
		)
	}
}

func metadataStreamClient(workstationID string) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		return streamer(
			metadata.AppendToOutgoingContext(ctx, metadataWorkstationID, workstationID),
			desc,
			cc,
			method,
			opts...,
		)
	}
}
