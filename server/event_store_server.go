package server

import (
	"fmt"
	"log"
	"nats_example/eventstores/server/rpc"
	"net"

	mysqldao "nats_example/baselib/dao"
	"nats_example/baselib/mysql_client"
	eventstore "nats_example/baselib/proto"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	port = "50051"
)

func init() {

}

// cacheServer func
type eventstoreServer struct {
	server *grpc.Server
	impl   *rpc.EventStoreServiceImpl
}

// NewSyncServer func
func NewEventstoreServer() *eventstoreServer {
	return &eventstoreServer{}
}

// -----------------------------------------------------------------------------
// AppInstance interface
func (s *eventstoreServer) Initialize() error {
	var err error
	err = InitializeConfig()
	if err != nil {
		glog.V(1).Info(err)
		return err
	}
	mysql_client.InstallMysqlClientManager(Conf.MySQL)
	mysqldao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())

	// s.server = grpc_util.NewRPCServer(Conf.Server.Addr, &Conf.Server.RpcDiscovery)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()
	if err != nil {
		glog.V(1).Infof("Server init error: %v", err)
	}
	s.impl = rpc.NewEventStoreServiceImpl(mysqldao.GetOrderDAO("mysql"), Conf.NATS)
	eventstore.RegisterEventStoreServer(s.server, s.impl)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.server.Serve(lis); err != nil {

		log.Fatalf("failed to serve: %v", err)
	}

	glog.V(3).Info("Initialize done..!")
	return nil
}

// GetIdentification func;
func (s *eventstoreServer) GetIdentification() {
	return
}
func (s *eventstoreServer) RunLoop() {
	// go s.server.Serve(func(s2 *grpc.Server) {
	// 	s.impl = rpc.NewCacheServiceImpl(s.models, redis_client.GetRedisClient(dao.GLOBAL_CACHE))
	// 	mtproto.RegisterRPCCacheManagerServer(s2, s.impl)
	// })
	// s.client.Serve()
}
func (s *eventstoreServer) Destroy() {
	if s.impl != nil {
		s.impl.Destroy()
	}
	s.server.Stop()
}
