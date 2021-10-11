package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/pkg/db"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/pkg/gqlgen"
)

func main() {
	db, err := db.InitDB()

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	srv := gqlgen.GetServer(db)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Panic(http.ListenAndServe(":8080", nil))
}
