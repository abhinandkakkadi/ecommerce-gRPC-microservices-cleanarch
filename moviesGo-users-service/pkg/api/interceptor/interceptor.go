package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func WithServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

// general unary interceptor function to handle auth per RPC call as well as logging
func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	// TODO: add JWT validation logic
	fmt.Println("Before calling server handler")
	h, err := handler(ctx, req)
	fmt.Println("After calling server handler")

	return h, err
}
