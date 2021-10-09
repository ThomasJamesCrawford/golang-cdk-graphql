package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/pkg/db"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/pkg/gqlgen"
)

func main() {
	_, err := db.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.DB.Close()

	srv := gqlgen.GetServer()

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
