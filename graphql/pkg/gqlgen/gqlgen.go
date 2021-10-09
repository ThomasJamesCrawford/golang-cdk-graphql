package gqlgen

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/generated"
)

func GetServer() *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
}
