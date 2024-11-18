package gapi

import (
	"context"
	"log"
	"strings"

	"github.com/kenmobility/grpc-project/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	if strings.TrimSpace(req.UserId) == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	user, err := server.store.GetUserByPublicID(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user_id: %s", err)
	}

	//log.Println("pbUser: ", pbUser)

	log.Println("pbUser.User.UserId: ", user.ID)

	order, err := server.store.CreateOrder(ctx, int(user.ID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order: %s", err)
	}

	order.User = user

	rsp := &pb.CreateOrderResponse{
		Order: convertOrder(*order),
	}

	return rsp, nil
}
