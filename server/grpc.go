package server

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	logger := logrus.New()
	// logs.SetFormatter(logger)
	logger.SetLevel(logrus.WarnLevel)

	grpc_logrus.ReplaceGrpcLogger(logrus.NewEntry(logger))
}

func RunGRPCServer(registerServer func(server *grpc.Server)) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_logrus.StreamServerInterceptor(logrusEntry)),
		grpc.UnaryInterceptor(grpc_logrus.UnaryServerInterceptor(logrusEntry)),
	)

	registerServer(grpcServer)

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	logrus.WithField("port", "8080").Info("Starting: gRPC Listener")
	logrus.Fatal(grpcServer.Serve(listen))
}
