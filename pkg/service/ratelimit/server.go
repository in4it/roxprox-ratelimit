package ratelimit

import (
	"log"
	"net"

	ratelimitv3 "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v3"
	"google.golang.org/grpc"
)

//StartGrpcServer starts the grpc server for the RateLimitService
func StartGrpcServer(service *Service) {
	grpcServer := grpc.NewServer()
	ratelimitv3.RegisterRateLimitServiceServer(grpcServer, service)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://:8081")
	grpcServer.Serve(l)
}
