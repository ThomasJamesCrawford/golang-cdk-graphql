package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/dataloader"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/generated"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/model"
)

func (r *queryResolver) Company(ctx context.Context, id string) (*model.Company, error) {
	return dataloader.For(ctx).CompanyByID.Load(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
