package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor() grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(clientInterceptor)
}

func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	// TODO: authorization at interceptor
	ctx = metadata.AppendToOutgoingContext(ctx, "name", "abhinand K R")
	// Invoke the original method call
	err := invoker(ctx, method, req, reply, cc, opts...)

	// TODO: execute logic after invoking the rpc endpoint
	log.Printf("after interceptor")
	return err

}
