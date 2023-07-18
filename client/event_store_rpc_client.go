package event_store_client

import (
	"context"
	"log"
	grpc_util "nats_example/baselib/grpc_utils"
	eventstore "nats_example/baselib/proto"
	"time"
)

type EventStoreClient struct {
	Client eventstore.EventStoreClient
}

var (
	EventStoreInstance = &EventStoreClient{}
)

type contextNoCancel struct {
	ctx context.Context
}

func (c contextNoCancel) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (c contextNoCancel) Done() <-chan struct{}             { return nil }
func (c contextNoCancel) Err() error                        { return nil }
func (c contextNoCancel) Value(key interface{}) interface{} { return c.ctx.Value(key) }

// WithoutCancel returns a context that is never canceled.
func WithoutCancel(ctx context.Context) context.Context {
	return contextNoCancel{ctx: ctx}
}

// GetCacheClient func
func GetStoreClient() *EventStoreClient {
	return EventStoreInstance
}

// InstallSyncClient func
func InstallEventStoreClient() {
	rpcClient, err := grpc_util.NewRPCClientByServiceDiscovery(grpc_util.ServiceDiscoveryServerConfig{Target: "0.0.0.0:50051"})
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}

	EventStoreInstance.Client = eventstore.NewEventStoreClient(rpcClient)
}
