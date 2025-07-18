package order

import (
	"context"
	"log"
	"time"

	"dev.murali.go-microservice/order/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.OrdersServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := pb.NewOrdersServiceClient(conn)
	return &Client{
		conn:    conn,
		service: c,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostOrder(ctx context.Context, accountId string, products []OrderedProduct) (*Order, error) {
	protoProducts := []*pb.PostOrderRequest_OrderProduct{}

	for _, p := range products {
		protoProducts = append(protoProducts, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ID,
			Quantity:  p.Quantity,
		})
	}
	r, err := c.service.PostOrder(ctx, &pb.PostOrderRequest{
		AccountId: accountId,
		Products:  protoProducts,
	},
	)
	if err != nil {
		return nil, err
	}

	newOrder := r.Order
	newOrderCreatedAt := time.Time{}
	newOrderCreatedAt.UnmarshalBinary(newOrder.CreatedAt)
	return &Order{
		ID:         newOrder.Id,
		CreatedAt:  newOrderCreatedAt,
		TotalPrice: newOrder.TotalPrice,
		AccountID:  newOrder.AccountId,
		Products:   products,
	}, nil

}

func (c *Client) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	r, err := c.service.GetOrdersForAccount(ctx, &pb.GetOrdersForAccountRequest{
		AccountId: accountID,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	orders := []Order{}
	for _, OrderProto := range r.Orders {
		newOrder := Order{
			ID:         OrderProto.Id,
			TotalPrice: OrderProto.TotalPrice,
			AccountID:  OrderProto.AccountId,
		}

		newOrder.CreatedAt = time.Time{}
		newOrder.CreatedAt.UnmarshalBinary(OrderProto.CreatedAt)
		products := []OrderedProduct{}

		for _, p := range OrderProto.Products {
			products = append(products, OrderedProduct{
				ID:          p.Id,
				Quantity:    p.Quantity,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			})
		}
		newOrder.Products = products
		orders = append(orders, newOrder)
	}

	return orders, nil
}
