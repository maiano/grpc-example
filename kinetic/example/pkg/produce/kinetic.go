package produce

import (
	"context"

	apikinetic "github.com/maiano/grpc-example/kinetic/gen/go/api/v1"
)

type KineticEnergyServer struct {
	apikinetic.UnimplementedKineticServiceServer
}

func (s *KineticEnergyServer) Kinetic(ctx context.Context, r *apikinetic.KineticRequest) (*apikinetic.KineticResponse, error) {
	return &apikinetic.KineticResponse{Result: r.Mass * r.Velocity * r.Velocity / 2}, nil
}
