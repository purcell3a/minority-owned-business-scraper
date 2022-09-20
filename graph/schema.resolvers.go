package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/purcell3a/minority-biz-api/graph/generated"
	"github.com/purcell3a/minority-biz-api/graph/model"
)

// AddAll is the resolver for the addAll field.
func (r *mutationResolver) AddAll(ctx context.Context, input model.AddAllInput) (*model.AddAll, error) {
	panic(fmt.Errorf("not implemented: AddAll - addAll"))
}

// FetchAllProducts is the resolver for the FetchAllProducts field.
func (r *queryResolver) FetchAllProducts(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: FetchAllProducts - FetchAllProducts"))
}

// FetchProduct is the resolver for the FetchProduct field.
func (r *queryResolver) FetchProduct(ctx context.Context, input string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: FetchProduct - FetchProduct"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: Products - Products"))
}
func (r *queryResolver) Product(ctx context.Context, input string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: Product - Product"))
}
