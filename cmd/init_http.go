package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"project4/pkg/logger"
	"project4/pkg/service-component/pb"

	_ "project4/pkg"
)

const httpPort = ":6001"

func initHttp(ctx context.Context) {
	logger.Info("init http storage_service")
	serveMux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterRTServiceHandlerFromEndpoint(ctx, serveMux, grpcPort, opt)
	if err != nil {
		logger.Fatal("can't create http storage_service from grpc endpoint",
			zap.String("error", err.Error()))
	}

	mainMux := http.NewServeMux()

	mainMux.Handle("/", serveMux)

	mainMux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./pkg/apidocs.swagger.json")
	})

	mainMux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./swagger-ui"))))

	go func() {
		if err = http.ListenAndServe(":6001", mainMux); err != nil {
			logger.Fatal("can't start http storage_service",
				zap.String("error", err.Error()), zap.String("port", httpPort))
		}
	}()
}
