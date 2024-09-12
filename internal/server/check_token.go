package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"project4/internal/user"
)

func (i *Implementation) CheckToken(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	_, err := i.getAuthData(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) getAuthData(ctx context.Context) (user.AuthData, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return user.AuthData{}, errors.New("unauthenticated")
	}

	var token string
	tokens := md.Get("authorization")
	if len(tokens) >= 1 {
		token = tokens[0]
	}

	if token == "" {
		return user.AuthData{}, errors.New("unauthenticated")
	}

	authData, err := i.userService.GetAuthData(token)
	if err != nil {
		return user.AuthData{}, errors.New("unauthenticated")
	}

	return authData, nil
}

func (i *Implementation) validUser(ctx context.Context) (user.AuthData, error) {
	authData, err := i.getAuthData(ctx)
	if err != nil {
		return user.AuthData{}, fmt.Errorf("getAuthData: %w", err)
	}

	isBanned, err := i.restrictionsService.IsUserBanned(ctx, authData.ID)
	if err != nil {
		return user.AuthData{}, status.Error(codes.Internal, err.Error())
	}
	if isBanned {
		return user.AuthData{}, status.Error(codes.PermissionDenied, "user banned")
	}

	return authData, nil
}
