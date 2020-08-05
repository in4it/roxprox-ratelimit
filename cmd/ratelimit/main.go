package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	config "github.com/in4it/roxprox-ratelimit/pkg/config/ratelimit"
	"github.com/in4it/roxprox-ratelimit/pkg/service/ratelimit"
	storage "github.com/in4it/roxprox/pkg/storage"
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

	rules, err := config.GetRateLimitRules(s)
	if err != nil {
		fmt.Printf("Couldn't get rate limit rules: %s", err)
		os.Exit(1)
	}

	if len(rules) == 0 {
		fmt.Printf("No rules loaded. Make sure to use -storage-path or configure s3 storage.\n")
		os.Exit(1)
	}

	log.Printf("%d rules loaded\n", len(rules))

	// start server
	ratelimit.StartGrpcServer(rules)
}
