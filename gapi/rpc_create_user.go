package gapi

import (
	"context"
	"strings"

	"github.com/kenmobility/grpc-project/helpers"
	"github.com/kenmobility/grpc-project/models/dto"
	"github.com/kenmobility/grpc-project/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if strings.TrimSpace(req.Email) == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if strings.TrimSpace(req.FullName) == "" {
		return nil, status.Error(codes.InvalidArgument, "full_name is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	hashedPassword, err := helpers.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := dto.CreateUserParams{
		Password: hashedPassword,
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok { // try converting the error to pq.Error type
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "email already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(*user),
	}
	return rsp, nil
}
