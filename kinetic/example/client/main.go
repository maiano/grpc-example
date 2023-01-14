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
	connectTo := "localhost:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect on %s: %w", connectTo, err)
	}
	defer conn.Close()

	log.Println("connected to", connectTo)

	client := apikinetic.NewKineticServiceClient(conn)

	res, err := client.Kinetic(context.Background(), &apikinetic.KineticRequest{
		Data: &apikinetic.Data{
			Mass:     1,
			Velocity: 3,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to calculate energy value: %w", err)
	}

	fmt.Printf("energy value is %v\n", res.Result)
	return nil
}
