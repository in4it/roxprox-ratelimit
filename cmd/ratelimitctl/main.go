package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/in4it/roxprox-ratelimit/proto/cache"
	"google.golang.org/grpc"
)

const timeout = 60

func main() {
	var (
		host string
		port string
		cmd  string
	)
	flag.StringVar(&host, "host", "127.0.0.1", "host of management server")
	flag.StringVar(&port, "port", "50051", "port of management server")
	flag.StringVar(&cmd, "cmd", "", "cmd to send to management server")

	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if strings.ToLower(cmd) == "getcache" {
		getCache(conn)
	} else {
		fmt.Printf("Command not found. Try: -cmd getCache\n")
		return
	}
}

func getCache(conn *grpc.ClientConn) {
	client := cache.NewCacheClient(conn)

	req := cache.GetCacheRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	res, err := client.GetCache(ctx, &req)
	if err != nil {
		fmt.Printf("GetCache error: %s", err)
		return
	}

	fmt.Println("Server response (key:value):")
	fmt.Println("----------------------------")

	for _, v := range res.CacheItems {
		fmt.Printf("%s: %s\n", v.Key, v.Value)
	}
}
