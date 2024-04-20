package main

import (
	"fmt"
	"go-course/grpc-service/internal/db"
	"go-course/grpc-service/internal/rocket"
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
	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	fmt.Println("hello grpc-service")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
