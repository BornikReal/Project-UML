package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	desc "project4/pkg/service-component/pb"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Empty request")
	}

	token, err := i.userService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.LoginResponse{
		Token: token,
	}, nil
}
