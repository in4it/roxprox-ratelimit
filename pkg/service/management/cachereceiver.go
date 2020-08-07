package management

import (
	"context"

	config "github.com/in4it/roxprox-ratelimit/pkg/config/ratelimit"
	"github.com/in4it/roxprox-ratelimit/pkg/service/ratelimit"
	"github.com/in4it/roxprox-ratelimit/proto/cache"
)

type CacheReceiver struct {
	storage          config.Storage
	ratelimitService *ratelimit.Service
}

//SendNotification is triggered when a grpc SendNotification is send to the management server
func (n *CacheReceiver) GetCache(ctx context.Context, in *cache.GetCacheRequest) (*cache.GetCacheReply, error) {
	cacheEntries := n.ratelimitService.GetCache()
	items := []*cache.GetCacheReply_CacheItem{}
	for k, v := range cacheEntries {
		items = append(items, &cache.GetCacheReply_CacheItem{
			Key:   k,
			Value: v,
		})
	}
	return &cache.GetCacheReply{
		CacheItems: items,
	}, nil
}

func NewCacheReceiver(storage config.Storage, ratelimitService *ratelimit.Service) *CacheReceiver {
	return &CacheReceiver{
		storage:          storage,
		ratelimitService: ratelimitService,
	}
}
