package server

import (
	"net"

	pb "github.com/leonardogazio/external-call-test-mock-example/grpc/proto"
	"github.com/leonardogazio/external-call-test-mock-example/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type apiServer struct{ pb.ApiServiceServer }

func Start() (err error) {
	host := ":3000"
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterApiServiceServer(grpcServer, &apiServer{})
	reflection.Register(grpcServer)

	utils.Log.Info(strServerStarted, zap.String("host", host))

	return grpcServer.Serve(lis)
}
