package main

import (
	"fmt"
	"go-course/grpc-service/internal/db"
	"go-course/grpc-service/internal/rocket"
	"go-course/grpc-service/internal/transport/grpc"
	"log"
)

func Run() error {
	// responsible for initializing and starting
	// our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to migrate rocket store")
		return err
	}
	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("hello grpc-service")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
