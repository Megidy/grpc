package main

import (
	"log"
	"net"

	handler "github.com/Megidy/grpc/services/orders/handlers/orders"
	service "github.com/Megidy/grpc/services/orders/services"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	address string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		address: addr,
	}
}

func (s *gRPCServer) Run() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("failed to listen tcp protocol : %v", err)
	}
	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGrpcHandler(grpcServer, orderService)
	log.Printf("Started gRPCserver on port : %v", s.address)
	return grpcServer.Serve(listener)
}
