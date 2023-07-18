package rpc

import (
	"context"
	"fmt"
	"log"
	"nats_example/baselib/core/order"
	mysqldao "nats_example/baselib/dao/mysql_dao"
	natsclient "nats_example/baselib/nats_client"
	eventstore "nats_example/baselib/proto"

	"github.com/nats-io/nats.go"
)

// EventStoreServiceImpl type
type EventStoreServiceImpl struct {
	orderModel *order.OrderDAO
	natsConn   *nats.Conn
}

// NewEventStoreServiceImpl func
func NewEventStoreServiceImpl(orderDAO *mysqldao.OrderDAO, natsConf natsclient.NatsConf) *EventStoreServiceImpl {
	impl := &EventStoreServiceImpl{
		orderModel: order.NewOrderDAO(orderDAO),
		natsConn:   natsclient.NewNatsClient(natsConf),
	}

	return impl
}

func (e *EventStoreServiceImpl) CreateEvent(ctx context.Context, request *eventstore.CreateEventRequest) (*eventstore.CreateEventResponse, error) {
	// store to event_store using mysqldb
	e.orderModel.OrderMysqlDAO.Insert()
	log.Println("Event is created")
	// concurent with publish event to NATS
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
