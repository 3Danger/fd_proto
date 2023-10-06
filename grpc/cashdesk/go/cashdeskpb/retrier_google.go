package cashdeskpb

import (
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

//grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
//grpcOpenTracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

func RetryOptions(maxRetries uint, retryDelay time.Duration) []grpc.DialOption {
	callOptions := []grpcRetry.CallOption{
		grpcRetry.WithMax(maxRetries),
		grpcRetry.WithBackoff(grpcRetry.BackoffLinear(retryDelay)),
		grpcRetry.WithCodes(codes.NotFound, codes.Aborted, codes.Unavailable),
	}
	return []grpc.DialOption{
		grpc.WithStreamInterceptor(grpcRetry.StreamClientInterceptor(callOptions...)),
		grpc.WithUnaryInterceptor(grpcRetry.UnaryClientInterceptor(callOptions...)),
	}
}
