package rpc

import (
	"context"
	"fmt"
	"log"
	eventstore "nats_example/baselib/proto"
)

// EventStoreServiceImpl type
type EventStoreServiceImpl struct {
}

// NewEventStoreServiceImpl func
func NewEventStoreServiceImpl() *EventStoreServiceImpl {
	impl := &EventStoreServiceImpl{}

	return impl
}

func (e *EventStoreServiceImpl) CreateEvent(ctx context.Context, request *eventstore.CreateEventRequest) (*eventstore.CreateEventResponse, error) {
	log.Printf("CreateEvent with: %v", request)
	return &eventstore.CreateEventResponse{
		IsSuccess: true,
	}, nil
}

func (e *EventStoreServiceImpl) GetEvents(ctx context.Context, request *eventstore.GetEventsRequest) (*eventstore.GetEventsResponse, error) {
	fmt.Printf("GetEvents with: %s", request.EventId)
	return &eventstore.GetEventsResponse{}, nil
}

func (e *EventStoreServiceImpl) GetEventsStream(*eventstore.GetEventsRequest, eventstore.EventStore_GetEventsStreamServer) error {
	return nil
}

// Destroy func
func (s *EventStoreServiceImpl) Destroy() {
	// s.closeChan <- 1
}
