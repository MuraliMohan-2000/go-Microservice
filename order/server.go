package order

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"dev.murali.go-microservice/account"
	"dev.murali.go-microservice/catalog"
	"dev.murali.go-microservice/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service       Service
	accountClient *account.Client
	catalogClient *catalog.Client
}

func ListenGRPC(s Service, accountURL, catalogURL string, port int) error {
	accountClient, err := account.NewCLient(accountURL)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.NewClient(catalogURL)
	if err != nil {
		accountClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
	}
	serv := grpc.NewServer()
	pb.RegisterOrdersServiceServer(serv, &grpcServer{
		service:       s,
		accountClient: accountClient,
		catalogClient: catalogClient,
	})

	reflection.Register(serv)
	return serv.Serve(lis)

}

func (s *grpcServer) PostOrder(ctx context.Context, r *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	_, err := s.accountClient.GetAccount(ctx, r.AccountId)
	if err != nil {
		log.Println("Error getting account:", err)
		return nil, errors.New("account not found")
	}

	productIDs := []string{}
	orderedProducts, err := s.catalogClient.GetProducts(ctx, productIDs, 0, 0, "")
	if err != nil {
		log.Println("Error getting products: ", err)
		return nil, errors.New("product not found")
	}

	products := []OrderedProduct{}
	for _, p := range orderedProducts {
		product := OrderedProduct{
			ID:          p.ID,
			Quantity:    0,
			Price:       p.Price,
			Name:        p.Name,
			Description: p.Description,
		}
		for _, rp := range r.Products {

			if rp.ProductId == p.ID {
				product.Quantity = rp.Quantity
				break
			}
		}
		if product.Quantity != 0 {
			products = append(products, product)
		}
	}

	order, err := s.PostOrder(ctx, r.AccountId, products)
	if err != nil {
		log.Println("Error posting order :", err)
		return nil, errors.New("could not post order")
	}

	orderProto := &pb.Order{}

}
