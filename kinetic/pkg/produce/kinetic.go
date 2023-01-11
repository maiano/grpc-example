package produce

import (
	"context"
	"log"

	apikinetic "github.com/maiano/grpc-example/kinetic/gen/go/api/v1"
)

type KineticEnergyServer struct {
	apikinetic.UnimplementedKineticServiceServer
}

func (s *KineticEnergyServer) Kinetic(ctx context.Context, r *apikinetic.KineticRequest) (*apikinetic.KineticResponse, error) {
	result := r.Mass * r.Velocity * r.Velocity / 2
	log.Printf("got a request: m = %.02f, v = %.02f - kinetic energy is %f", r.Mass, r.Velocity, result)
	return &apikinetic.KineticResponse{Result: r.Mass * r.Velocity * r.Velocity / 2}, nil
}
