package main

import (
	"dev.murali.go-microservice/account"
	"dev.murali.go-microservice/catalog"
	"dev.murali.go-microservice/order"
	"github.com/99designs/gqlgen/graphql"
)

type Server struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {

	accountclient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	catalogclient, err := catalog.NewClient(catalogUrl)
	if err != nil {
		accountclient.Close()
		return nil, err
	}

	orderclient, err := order.NewClient(orderUrl)
	if err != nil {
		accountclient.Close()
		catalogclient.Close()
		return nil, err
	}

	return &Server{
		accountclient,
		catalogclient,
		orderclient,
	}, nil

}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}

}

func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
