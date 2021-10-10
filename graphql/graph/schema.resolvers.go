package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/generated"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/model"
)

func (r *companyResolver) ID(ctx context.Context, obj *model.Company) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Email(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Company returns generated.CompanyResolver implementation.
func (r *Resolver) Company() generated.CompanyResolver { return &companyResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type companyResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
