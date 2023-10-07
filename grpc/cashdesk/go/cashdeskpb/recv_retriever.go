package cashdeskpb

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type cashdeskStreamEventsClientWithRetry struct {
	Cashdesk_StreamEventsClient
	cnf  Config
	opts []grpc.DialOption
}

type Config struct {
	Addr string

	//reconnect&retry
	Attempts   int
	RetryDelay time.Duration
}

func NewCashdeskClientWithRetry(cnf *Config) (Cashdesk_StreamEventsClient, error) {

}

func NewRecvRetriever(cashDeskStreamClient Cashdesk_StreamEventsClient, attempts int, retryDelay time.Duration) Cashdesk_StreamEventsClient {

	return &recvRetriever{
		Cashdesk_StreamEventsClient: cashDeskStreamClient,

		attempts:   attempts,
		retryDelay: retryDelay,
	}
}

func (r *recvRetriever) Recv() (*StreamWorkstationEventsResponse, error) {
	response, err := r.Cashdesk_StreamEventsClient.Recv()
	if err != nil {
		for a := r.attempts; a > 0 && err != nil; a-- {
			time.Sleep(r.retryDelay)
			response, err = r.Cashdesk_StreamEventsClient.Recv()
		}
	}
	return response, err
}

func retrier(ctx context.Context, attempts int, delay time.Duration, f func(ctx context.Context) error) (err error) {
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for ; attempts > 0; attempts-- {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err = f(ctx); err == nil {
				return nil
			}
		}
	}
	return err
}

func connect(ctx context.Context, serverAddr string, opts ...grpc.DialOption) error {
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return fmt.Errorf("dealing conn: %w", err)
	}
	client := NewCashdeskClient(conn)
	stream, err := client.StreamEvents(ctx)

}
