package dataloader

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type key int
type Table int

const (
	keyDataloaders key = iota
)

type Loaders struct {
	CompanyByID CompanyLoader
}

const (
	Company Table = iota
)

func (t Table) String() string {
	switch t {
	case Company:
		return "company"
	}

	panic("Table name needs to be in Table constant")
}

func generateArguments(ids []string) (string, []interface{}) {
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))

	for i := 0; i < len(ids); i++ {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = ids[i]
	}

	return strings.Join(placeholders, ","), args
}

func Middleware(conn *pgxpool.Pool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), keyDataloaders, &Loaders{
			CompanyByID: CompanyLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*model.Company, []error) {
					in, args := generateArguments(ids)
					res, err := conn.Query(context.Background(), `SELECT id FROM company WHERE id IN (`+in+`)`, args...)

					if err != nil {
						log.Panic(err)
					}

					defer res.Close()

					companyByID := map[string]*model.Company{}
					for res.Next() {
						company := model.Company{}
						err := res.Scan(&company.ID)

						if err != nil {
							log.Panic(err)
						}

						companyByID[company.ID] = &company
					}

					companies := make([]*model.Company, len(ids))
					for i, id := range ids {
						companies[i] = companyByID[id]
					}

					return companies, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(keyDataloaders).(*Loaders)
}
