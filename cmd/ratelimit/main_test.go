package main

import (
	"fmt"
	"testing"
	"time"

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

	queue := make(chan []*notification.NotificationRequest_NotificationItem)
	go receiveConfigUpdates(s, ratelimitService, queue)

	notification := []*notification.NotificationRequest_NotificationItem{
		{
			Filename:  "ratelimit.yaml",
			EventName: "ObjectCreated:Put",
		},
	}
	queue <- notification
	testOk := false
	for i := 0; i < 10; i++ {
		fmt.Printf("Waiting until version updates: version: %d\n", ratelimitService.GetVersion())
		if ratelimitService.GetVersion() == 2 {
			testOk = true
			break
		}
		i++
		time.Sleep(1 * time.Second)
	}
	if !testOk {
		t.Errorf("Test failed, unable to update config")
		return
	}
}
