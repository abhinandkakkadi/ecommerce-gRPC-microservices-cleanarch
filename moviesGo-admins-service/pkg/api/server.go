package api

import (
	"fmt"
	"net"

	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/admin/pb"
	"github.com/abhinandkakkadi/moviesgo-admin-service/pkg/config"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.AdminAuthServiceServer) (*Server, error) {

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterAdminAuthServiceServer(newServer, server)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50054")
	return c.server.Serve(c.listener)
}
