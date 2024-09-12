package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"project4/internal/server"
	"project4/pkg/logger"
	"project4/pkg/service-component/pb"
)

const grpcPort = ":6002"

func initGrpc(service *server.Implementation) {
	logger.Info("init grpc storage_service")
	grpcServer := grpc.NewServer()
	pb.RegisterRTServiceServer(grpcServer, service)
	reflection.Register(grpcServer)
	lsn, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Fatal("listening port ended with error",
			zap.String("error", err.Error()), zap.String("port", grpcPort))
	}

	go func() {
		if err = grpcServer.Serve(lsn); err != nil {
			logger.Fatal("grpc server ended with error",
				zap.String("error", err.Error()), zap.String("port", grpcPort))
		}
	}()
}
