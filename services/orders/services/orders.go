package service

import (
	"context"

	orders "github.com/Megidy/grpc/services/common/genproto/orders/protobuf"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
func (s *OrderService) CreateOrder(ctx context.Context, orderRequest *orders.Order) error {
	ordersDb = append(ordersDb, orderRequest)
	return nil
}
func (s *OrderService) GetOrders(context.Context) []*orders.Order {
	return ordersDb
}
