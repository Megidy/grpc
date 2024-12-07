package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	orders "github.com/Megidy/grpc/services/common/genproto/orders/protobuf"
	"google.golang.org/grpc"
)

type httpHandler struct {
}

var conn grpc.ClientConnInterface

func NewHttpHandler() *httpHandler {
	return &httpHandler{}
}
func (h *httpHandler) RegisterRoutes(router *http.ServeMux, c grpc.ClientConnInterface) {
	conn = c
	router.HandleFunc("/", h.GetOrders)
}
func (h *httpHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	client := orders.NewOrderServiceClient(conn)
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()
	// _, err := client.CreateOrder(ctx, &orders.CreateOrderRequest{
	// 	CustomerID: 1,
	// 	ProductID:  2,
	// 	Quantity:   3333,
	// })
	// if err != nil {
	// 	log.Fatalf("client error : %v", err)
	// }
	res, err := client.GetOrders(ctx, &orders.GetOrderRequest{
		CustomeID: 42,
	})
	if err != nil {
		log.Fatalf("client error: %v", err)
	}
	t := template.Must(template.New("orders").Parse(ordersTemplate))

	if err := t.Execute(w, res.GetOrders()); err != nil {
		log.Fatalf("template error: %v", err)
	}

}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderId}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
