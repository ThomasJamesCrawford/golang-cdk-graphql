package main

import (
	"log"
	"net/http"

	"github.com/ThomasJamesCrawford/graphql-api/internal/db"
	"github.com/ThomasJamesCrawford/graphql-api/internal/graphjin"
	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func init() {
	db.InitDB()
}

func main() {
	gj, err := graphjin.GetGraphJin(db.DB)

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Post("/graphql", graphjin.Handler(gj))

	http.ListenAndServe(":3000", r)
}
