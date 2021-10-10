package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/generated"
	"github.com/ThomasJamesCrawford/golang-cdk-graphql/graphql/graph/model"
)

func (r *serviceResolver) ID(ctx context.Context, obj *model.Service) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *serviceResolver) Name(ctx context.Context, obj *model.Service) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *serviceResolver) Description(ctx context.Context, obj *model.Service) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *serviceResolver) Length(ctx context.Context, obj *model.Service) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *serviceResolver) Price(ctx context.Context, obj *model.Service) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Service returns generated.ServiceResolver implementation.
func (r *Resolver) Service() generated.ServiceResolver { return &serviceResolver{r} }

type serviceResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *serviceResolver) Technicians(ctx context.Context, obj *model.Service) ([]*model.Technician, error) {
	panic(fmt.Errorf("not implemented"))
}
