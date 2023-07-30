package api

import (
	"fmt"
	"net"

	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/config"
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/product/pb"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.ProductServiceServer) (*Server, error) {

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterProductServiceServer(newServer, server)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listener)
}
