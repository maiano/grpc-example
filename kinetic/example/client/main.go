package main

import (
	"context"
	"fmt"
	"log"

	apikinetic "github.com/maiano/grpc-example/kinetic/gen/go/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect on %s: %w", connectTo, err)
	}
	log.Println("connected to", connectTo)

	e := apikinetic.NewKineticServiceClient(conn)
	if _, err := e.Kinetic(context.Background(), &apikinetic.KineticRequest{
		Mass:     1,
		Velocity: 3,
	}); err != nil {
		return fmt.Errorf("failed to produce: %w", err)
	}

	log.Println("ok")
	return nil
}
