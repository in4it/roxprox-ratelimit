package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	config "github.com/in4it/roxprox-ratelimit/pkg/config/ratelimit"
	"github.com/in4it/roxprox-ratelimit/pkg/service/management"
	"github.com/in4it/roxprox-ratelimit/pkg/service/ratelimit"
	storage "github.com/in4it/roxprox/pkg/storage"
	"github.com/juju/loggo"
)

func parseFlags() (storage.Storage, string) {
	var (
		err            error
		loglevel       string
		storageType    string
		storagePath    string
		storageBucket  string
		awsRegion      string
		managementPort string
		s              storage.Storage
	)
	flag.StringVar(&loglevel, "loglevel", "INFO", "log level")
	flag.StringVar(&storageType, "storage-type", "local", "storage type")
	flag.StringVar(&storagePath, "storage-path", "", "storage path")
	flag.StringVar(&storageBucket, "storage-bucket", "", "s3 storage bucket")
	flag.StringVar(&awsRegion, "aws-region", "", "AWS region")
	flag.StringVar(&managementPort, "management-port", "50051", "Port for the management server")

	flag.Parse()

	if storageType == "local" {
		s, err = storage.NewLocalStorage(storagePath)
		if err != nil {
			fmt.Printf("Couldn't inialize storage: %s", err)
			os.Exit(1)
		}
	} else if storageType == "s3" {
		s, err = storage.NewS3Storage(storageBucket, storagePath, awsRegion, "", false)
		if err != nil {
			fmt.Printf("Couldn't inialize storage: %s", err)
			os.Exit(1)
		}
	} else {
		panic("unknown storage")
	}

	// init logging
	loglevel = strings.ToUpper(loglevel)

	if loglevel == "DEBUG" || loglevel == "INFO" || loglevel == "TRACE" || loglevel == "ERROR" {
		loggo.ConfigureLoggers(`<root>=` + loglevel)
	} else {
		loggo.ConfigureLoggers(`<root>=INFO`)
	}
	return s, managementPort
}

func main() {
	// init storage, parse flags
	s, managementPort := parseFlags()

	// init rate limit service
	ratelimitService := initRateLimitService(s)

	// start management server
	notificationReceiver := management.NewNotificationReceiver(s, ratelimitService)
	cacheReceiver := management.NewCacheReceiver(s, ratelimitService)
	err := management.NewServer(managementPort, notificationReceiver, cacheReceiver)
	if err != nil {
		fmt.Printf("Couldn't start management interface: %s", err)
		os.Exit(1)
	}

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
