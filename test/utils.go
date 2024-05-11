//+acceptance

package test

import (
	rocket "go-course/grpc-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetClient() rocket.RocketServiceClient {
	log.Println("testing grpc client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	rocketClient := rocket.NewRocketServiceClient(conn)
	return rocketClient
}
