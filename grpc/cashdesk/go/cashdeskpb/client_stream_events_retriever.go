package cashdeskpb

import "time"

type recvStreamEventsRetriever struct {
	Cashdesk_StreamEventsClient
	attempts   int
	retryDelay time.Duration
}

func NewRecvRetriever(stream Cashdesk_StreamEventsClient, attempts int, retryDelay time.Duration) Cashdesk_StreamEventsClient {
	return &recvStreamEventsRetriever{
		Cashdesk_StreamEventsClient: stream,

		attempts:   attempts,
		retryDelay: retryDelay,
	}
}

func (r *recvStreamEventsRetriever) Recv() (*StreamWorkstationEventsResponse, error) {
	response, err := r.Cashdesk_StreamEventsClient.Recv()
	if err != nil {
		for i := r.attempts; i > 0 && err != nil; i-- {
			time.Sleep(r.retryDelay)
			response, err = r.Cashdesk_StreamEventsClient.Recv()
		}
	}
	return response, err
}
