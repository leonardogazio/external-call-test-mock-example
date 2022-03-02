package main

import (
	foo "github.com/leonardogazio/external-call-test-mock-example/client/foo-service-api"
	grpc "github.com/leonardogazio/external-call-test-mock-example/grpc/server"
	"github.com/leonardogazio/external-call-test-mock-example/utils"
	"go.uber.org/zap"
)

func main() {
	if err := utils.NewLog(); err != nil {
		panic(err)
	}

	//foo-service-api client instance
	foo.Client = foo.New("http://testingmock.free.beeceptor.com/")

	if err := grpc.Start(); err != nil {
		utils.Log.Error(grpc.StrFailGRPCStart, zap.Error(err))
	}
}
