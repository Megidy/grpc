package handler

import (
	"net/http"

	orders "github.com/Megidy/grpc/services/common/genproto/orders/protobuf"
	"github.com/Megidy/grpc/services/common/utils"
	"github.com/Megidy/grpc/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpHandler(orderService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{
		ordersService: orderService,
	}

}

func (h *OrdersHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var payload orders.CreateOrderRequest
	err := utils.ParseJson(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	order := &orders.Order{
		OrderId:    41,
		CustomerID: payload.GetCustomerID(),
		ProductID:  payload.GetProductID(),
		Quantity:   payload.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	res := &orders.CreateOrderResponse{
		Status: "Success",
	}
	utils.WriteJson(w, http.StatusOK, res)
}
