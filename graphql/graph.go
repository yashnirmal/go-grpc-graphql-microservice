package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/yashnirmal/grpc-graphql-microservice/graphql/graph/resolvers"
)

type Server struct {
	// 	AccountClient
	// 	ProductClient
	//  OrderClient
}

func NewGraphQLServer(accountUrl string, catalogUrl string, orderUrl string) (*Server, error) {
	// accountClient, err := account.NewClient(accountUrl)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogClient, err := catalog.NewClient(catalogUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	return nil, err
	// }
	// orderClient, err := order.NewClient(orderUrl)
	// if err != nil {
	// 	accountClient.Close()
	// 	catalogClient.Close()
	// 	return nil, err
	// }

	// return &Server{
	// 	accountClient,
	// 	catalogClient,
	// 	orderClient,
	// }, nil
}

// func (s *Server) Mutation() resolvers.MutationResolver {
// 	return &resolvers.MutationResolver{Server: s}
// }

// func (s *Server) Query() resolvers.QueryResolver {
// 	return &resolvers.QueryResolver{server: s}
// }

// func (s *Server) Account() resolvers.AccountResolver {
// 	return &resolvers.AccountResolver{server: s}
// }

func (srv *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}