package management

import (
	"fmt"
	"net"

	"github.com/in4it/roxprox-ratelimit/proto/cache"
	n "github.com/in4it/roxprox/proto/notification"
	"github.com/juju/loggo"
	"google.golang.org/grpc"
)

var logger = loggo.GetLogger("management")

func NewServer(port string, notificationReceiver n.NotificationServer, cacheReceiver cache.CacheServer) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	logger.Infof("Starting grpc management interface")
	grpcServer := grpc.NewServer()

	// register servers
	n.RegisterNotificationServer(grpcServer, notificationReceiver)
	cache.RegisterCacheServer(grpcServer, cacheReceiver)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Errorf("failed to serve: %v", err)
		}
	}()

	return nil
}
