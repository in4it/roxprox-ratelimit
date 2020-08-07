package management

import (
	"context"

	config "github.com/in4it/roxprox-ratelimit/pkg/config/ratelimit"
	"github.com/in4it/roxprox-ratelimit/pkg/service/ratelimit"
	notification "github.com/in4it/roxprox/proto/notification"
)

type NotificationReceiver struct {
	storage          config.Storage
	ratelimitService *ratelimit.Service
}

//SendNotification is triggered when a grpc SendNotification is send to the management server
func (n *NotificationReceiver) SendNotification(ctx context.Context, in *notification.NotificationRequest) (*notification.NotificationReply, error) {
	logger.Debugf("Received %d events", len(in.GetNotificationItem()))
	for _, v := range in.GetNotificationItem() {
		if v.EventName == "ObjectCreated:Put" {
			// handle put object
			rules, err := config.GetRateLimitRule(n.storage, v.Filename)
			if err != nil {
				logger.Errorf("Error fetching new object: %s", v.Filename)
			}
			logger.Infof("%d new rules loaded\n", len(rules))
			n.ratelimitService.PutRules(rules)
		}
	}
	return &notification.NotificationReply{Result: true}, nil
}

func NewNotificationReceiver(storage config.Storage, ratelimitService *ratelimit.Service) *NotificationReceiver {
	return &NotificationReceiver{
		storage:          storage,
		ratelimitService: ratelimitService,
	}
}
