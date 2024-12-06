package main

import (
	"log"
	"net/http"

	"github.com/Megidy/grpc/services/kitchen/handlers"
)

type httpServer struct {
	addr string
}

func NewHTTPServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	conn, err := NewGrpcClient(":9000")
	if err != nil {

		return err
	}
	defer conn.Close()

	h := handlers.NewHttpHandler()
	h.RegisterRoutes(router, conn)
	log.Println("started http Server on port : ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
