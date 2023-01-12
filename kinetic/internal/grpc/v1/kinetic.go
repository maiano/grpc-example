package produce

import (
	"context"
	"log"

	apikinetic "github.com/maiano/grpc-example/kinetic/gen/go/api/v1"
)

type KineticEnergyServer struct {
	apikinetic.UnimplementedKineticServiceServer
}

func (s *KineticEnergyServer) Kinetic(_ context.Context, req *apikinetic.KineticRequest) (*apikinetic.KineticResponse, error) {

	log.Printf("got a request: m = %.02f, v = %.02f\n", req.Mass, req.Velocity)
	result := req.Mass * req.Velocity * req.Velocity / 2
	return &apikinetic.KineticResponse{Result: result}, nil
}
