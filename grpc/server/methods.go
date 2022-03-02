package server

import (
	"context"

	foo "github.com/leonardogazio/external-call-test-mock-example/client/foo-service-api"
	pb "github.com/leonardogazio/external-call-test-mock-example/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*apiServer) GetFoo(ctx context.Context, _ *emptypb.Empty) (*pb.FooRes, error) {
	fooBar, err := foo.Client.GetFoo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.FooRes{Foo: fooBar.Foo}, nil
}
