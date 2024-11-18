package gapi

import (
	"github.com/kenmobility/grpc-project/models"
	"github.com/kenmobility/grpc-project/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user models.User) *pb.User {
	return &pb.User{
		UserId:    int32(user.ID),
		PublicId:  user.PublicID,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

func convertOrder(order models.Order) *pb.Order {
	return &pb.Order{
		OrderId:   int32(order.ID),
		PublicId:  order.PublicID,
		Status:    order.Status,
		CreatedAt: timestamppb.New(order.CreatedAt),
	}
}

func convertFromPbUser(pbUser *pb.User) *models.User {
	return &models.User{
		ID:        uint(pbUser.UserId),
		PublicID:  pbUser.PublicId,
		Email:     pbUser.Email,
		FullName:  pbUser.FullName,
		CreatedAt: pbUser.CreatedAt.AsTime().Local(),
	}
}
