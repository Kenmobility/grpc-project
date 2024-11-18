package gapi

import (
	"github.com/kenmobility/grpc-project/models"
	"github.com/kenmobility/grpc-project/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user models.User) *pb.User {
	return &pb.User{
		UserId:    user.PublicID,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
