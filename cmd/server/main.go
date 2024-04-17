package main

import (
	"fmt"
	"log"
)

func Run() error {
	// responsible for initializing and starting
	// our gRPC server
	return nil
}

func main() {
	fmt.Println("hello grpc-service")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
