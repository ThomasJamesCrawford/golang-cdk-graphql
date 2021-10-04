package graphjin

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dosco/graphjin/core"
	"github.com/go-chi/render"
)

type GraphQLBody struct {
	Query     string
	Variables json.RawMessage
}

func Handler(gj *core.GraphJin) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body GraphQLBody
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		res, err := gj.GraphQL(r.Context(), body.Query, body.Variables, nil)

		if err != nil {
			log.Fatal(err)
		}

		render.JSON(w, r, res)
	}
}

func GetGraphJin(db *sql.DB) (*core.GraphJin, error) {
	return core.NewGraphJin(nil, db)
}
