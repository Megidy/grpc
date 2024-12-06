package main

import (
	"log"
	"net/http"

	handler "github.com/Megidy/grpc/services/orders/handlers/orders"
	service "github.com/Megidy/grpc/services/orders/services"
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
	OrdersService := service.NewOrderService()
	httpHandler := handler.NewHttpHandler(OrdersService)
	httpHandler.RegisterRoutes(router)
	log.Println("started http Server on port : ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
