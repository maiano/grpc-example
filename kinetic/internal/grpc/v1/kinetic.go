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

	log.Printf("got a request: %+v\n", req.Data)
	result := req.Data.Mass * req.Data.Velocity * req.Data.Velocity / 2
	return &apikinetic.KineticResponse{Result: result}, nil
}
