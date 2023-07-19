package rpc

import (
	"context"
	"fmt"
	"log"
	"nats_example/baselib/core/event"
	mysqldao "nats_example/baselib/dao/mysql_dao"
	natsclient "nats_example/baselib/nats_client"
	eventstore "nats_example/baselib/proto"

	"github.com/nats-io/nats.go"
)

// EventStoreServiceImpl type
type EventStoreServiceImpl struct {
	eventModel *event.EventDAO
	natsConn   *nats.Conn
}

// NewEventStoreServiceImpl func
func NewEventStoreServiceImpl(eventDAO *mysqldao.EventDAO, natsConf natsclient.NatsConf) *EventStoreServiceImpl {
	impl := &EventStoreServiceImpl{
		eventModel: event.NewEventDAO(eventDAO),
		natsConn:   natsclient.NewNatsClient(natsConf),
	}

	return impl
}

func (e *EventStoreServiceImpl) CreateEvent(ctx context.Context, request *eventstore.CreateEventRequest) (*eventstore.CreateEventResponse, error) {
	// store to event_store using mysqldb
	e.eventModel.OrderMysqlDAO.Insert(ctx, request.Event)
	log.Println("Event is created")
	// concurent with publish event to NATS
	eventMsg := []byte(request.Event.EventData)
	e.natsConn.Publish(request.Event.EventType, eventMsg)
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

// TODO: Create snapshot for each order_id

// Destroy func
func (s *EventStoreServiceImpl) Destroy() {
	// s.closeChan <- 1
}
