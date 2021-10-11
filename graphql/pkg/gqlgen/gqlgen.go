package gqlgen

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/dataloader"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/generated"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetServer(conn *pgxpool.Pool) http.Handler {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return dataloader.Middleware(conn, srv)
}
