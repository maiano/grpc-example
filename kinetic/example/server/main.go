package main

import (
	"log"
	"net"

	apikinetic "github.com/maiano/grpc-example/kinetic/gen/go/api/v1"
	v1 "github.com/maiano/grpc-example/kinetic/internal/grpc/v1"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}

	server := grpc.NewServer()

	s := &v1.KineticEnergyServer{}

	apikinetic.RegisterKineticServiceServer(server, s)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}
}
