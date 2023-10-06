package cashdeskpb

import "time"

type recvRetriever struct {
	Cashdesk_StreamEventsClient
	attempts   int
	retryDelay time.Duration
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
