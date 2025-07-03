package main

import "github.com/99designs/gqlgen/graphql"

type Server struct {
	// accountClient *account.client
	// catalogClient *catalog.client
	// orderClient   *order.client
}

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {

	// 	accountclient, err := account.NewClient(accountUrl)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	catalogclient, err := catalog.NewClient(catalogUrl)
	// 	if err != nil {
	// 		accountclient.Close()
	// 		return nil, err
	// 	}

	// 	orderclient, err := order.NewClient(orderUrl)
	// 	if err != nil {
	// 		accountclient.Close()
	// 		catalogclient.Close()
	// 		return nil, err
	// 	}

	return &Server{
		// 		accountclient,
		// 		catalogclient,
		// 		orderclient,
	}, nil

}

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}

// }

// func (s *Server) Account() AccountResolver {
// 	return &accountResolver{
// 		server: s,
// 	}
// }

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
