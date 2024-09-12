package server

import (
	"context"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"project4/internal/user"
	desc "project4/pkg/service-component/pb"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*emptypb.Empty, error) {
	_, err := i.userService.Create(ctx, user.CreateUserRequest{
		Role:               user.RoleUser,
		Username:           req.Username,
		ProfileDescription: req.ProfileDescription,
		Avatar:             req.Avatar,
		Email:              req.Email,
		Password:           req.Password,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
