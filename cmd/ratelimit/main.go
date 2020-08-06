package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	config "github.com/in4it/roxprox-ratelimit/pkg/config/ratelimit"
	"github.com/in4it/roxprox-ratelimit/pkg/service/ratelimit"
	"github.com/in4it/roxprox/pkg/management"
	storage "github.com/in4it/roxprox/pkg/storage"
	"github.com/in4it/roxprox/proto/notification"
)

func initStorage() storage.Storage {
	var (
		err           error
		loglevel      string
		storageType   string
		storagePath   string
		storageBucket string
		awsRegion     string
		s             storage.Storage
	)
	flag.StringVar(&loglevel, "loglevel", "INFO", "log level")
	flag.StringVar(&storageType, "storage-type", "local", "storage type")
	flag.StringVar(&storagePath, "storage-path", "", "storage path")
	flag.StringVar(&storageBucket, "storage-bucket", "", "s3 storage bucket")
	flag.StringVar(&awsRegion, "aws-region", "", "AWS region")

	flag.Parse()

	if storageType == "local" {
		s, err = storage.NewLocalStorage(storagePath)
		if err != nil {
			fmt.Printf("Couldn't inialize storage: %s", err)
			os.Exit(1)
		}
	} else if storageType == "s3" {
		s, err = storage.NewS3Storage(storageBucket, storagePath, awsRegion)
		if err != nil {
			fmt.Printf("Couldn't inialize storage: %s", err)
			os.Exit(1)
		}
	} else {
		panic("unknown storage")
	}
	return s
}

func main() {
	// init storage
	s := initStorage()

	// start management server
	notificationQueue, err := management.NewServer()
	if err != nil {
		fmt.Printf("Couldn't start management interface: %s\n", err)
		os.Exit(1)
	}

	// init rate limit service
	ratelimitService := initRateLimitService(s)

	// listen for rule updates
	go receiveConfigUpdates(s, ratelimitService, notificationQueue.GetQueue())

	// start grpc server with ratelimitService
	ratelimit.StartGrpcServer(ratelimitService)

}

func initRateLimitService(s config.Storage) *ratelimit.Service {
	rules, err := config.GetRateLimitRules(s)
	if err != nil {
		fmt.Printf("Couldn't get rate limit rules: %s", err)
		os.Exit(1)
	}

	if len(rules) == 0 {
		fmt.Printf("Warning: No rules loaded. Make sure to use -storage-path or configure s3 storage.\n")
	} else {
		log.Printf("%d rules loaded\n", len(rules))
	}
	// initialize ratelimit service
	cacheSizeMb := ratelimit.CacheMbSizeDefault
	if os.Getenv("CACHE_SIZE_MB") != "" {
		cacheSizeMb, err = strconv.Atoi(os.Getenv("CACHE_SIZE_MB"))
		if err != nil {
			fmt.Printf("Couldn't parse CACHE_SIZE_MB. Unknown integer value: %s. Using default", os.Getenv("CACHE_SIZE_MB"))
			cacheSizeMb = ratelimit.CacheMbSizeDefault
		}
	}
	ratelimitService := ratelimit.NewRateLimitService(cacheSizeMb)

	// put initial rules
	ratelimitService.PutRules(rules)

	return ratelimitService
}

func receiveConfigUpdates(s config.Storage, ratelimitService *ratelimit.Service, queue chan []*notification.NotificationRequest_NotificationItem) {
	for {
		notifications := <-queue

		for _, v := range notifications {
			if v.EventName == "ObjectCreated:Put" {
				// handle put object
				rules, err := config.GetRateLimitRule(s, v.Filename)
				if err != nil {
					fmt.Printf("Error fetching new object: %s", v.Filename)
				}
				ratelimitService.PutRules(rules)
			}
		}
	}
}
