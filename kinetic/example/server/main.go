package main

import (
	"example/microservices/example/pkg/produce"
	apikinetic "example/microservices/gen/go/api/v1"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}

	server := grpc.NewServer()

	e := &produce.KineticEnergyServer{}

	apikinetic.RegisterKineticServiceServer(server, e)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}
}
