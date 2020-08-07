package main

import (
	"context"
	"testing"

	"github.com/in4it/roxprox-ratelimit/pkg/service/management"
	storage "github.com/in4it/roxprox/pkg/storage"
	"github.com/in4it/roxprox/proto/notification"
)

func TestReceiveConfigUpdates(t *testing.T) {
	// init storage
	s, err := storage.NewLocalStorage("../../resources/")
	if err != nil {
		t.Errorf("Couldn't initialize storage")
		return
	}
	// init rate limit service
	ratelimitService := initRateLimitService(s)

	notificationReceiver := management.NewNotificationReceiver(s, ratelimitService)

	notification := &notification.NotificationRequest{
		NotificationItem: []*notification.NotificationRequest_NotificationItem{
			{
				Filename:  "ratelimit.yaml",
				EventName: "ObjectCreated:Put",
			},
		},
	}

	notificationReceiver.SendNotification(context.Background(), notification)

	if ratelimitService.GetVersion() != 2 {
		t.Errorf("Test failed, unable to update config")
		return

	}

}
