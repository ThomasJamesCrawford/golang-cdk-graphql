package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/pkg/db"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/pkg/gqlgen"
	"github.com/bradleyjkemp/cupaloy"
)

func Test_main(t *testing.T) {
	os.Setenv("DB_URL", "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	t.Run("Smoke test", func(t *testing.T) {
		postBody, _ := json.Marshal(map[string]string{
			"query": `
query {
	company(id: "aa0ed2d3-4bee-4471-9421-825b4b51d477") {
		id
	}
}
`,
		})

		req, err := http.NewRequest("POST", "/graphql", bytes.NewBuffer(postBody))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()

		db, _ := db.InitDB()

		db.Query(context.Background(), "INSERT INTO company (id) VALUES ('aa0ed2d3-4bee-4471-9421-825b4b51d477')")

		gqlgen.GetServer(db).ServeHTTP(res, req)

		cupaloy.SnapshotT(t, res.Body.String())
	})
}
