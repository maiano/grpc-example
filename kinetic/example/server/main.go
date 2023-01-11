package main

import (
	"log"
	"net"

	"github.com/maiano/grpc-example/kinetic/example/pkg/produce"
	apikinetic "github.com/maiano/grpc-example/kinetic/gen/go/api/v1"

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
