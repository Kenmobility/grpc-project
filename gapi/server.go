package gapi

import (
	"github.com/kenmobility/grpc-project/config"
	"github.com/kenmobility/grpc-project/pb"
	"github.com/kenmobility/grpc-project/repo"
)

// Server serves gRPC requests.
type Server struct {
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedUserServiceServer
	config config.Config
	store  repo.StoreRepository
}

// NewServer creates a new gRPC server.
func NewServer(config config.Config, store repo.StoreRepository) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
