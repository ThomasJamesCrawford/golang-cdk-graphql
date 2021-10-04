package main

import (
	"log"

	"github.com/ThomasJamesCrawford/graphql-api/internal/db"
	"github.com/ThomasJamesCrawford/graphql-api/internal/graphjin"
	"github.com/aws/aws-lambda-go/lambda"
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

	lambda.Start(graphjin.Handler(gj))
}
