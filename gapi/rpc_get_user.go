package gapi

import (
	"context"
	"strings"

	"github.com/kenmobility/grpc-project/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if strings.TrimSpace(req.UserId) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user_id is required")
	}

	user, err := server.store.GetUserByPublicID(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %s", err)
	}

	rsp := &pb.GetUserResponse{
		User: convertUser(*user),
	}
	return rsp, nil
}
