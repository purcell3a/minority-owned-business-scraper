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
func (r *mutationResolver) AddAll(ctx context.Context, input model.AddAllInput) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: AddAll - addAll"))
}

// Products is the resolver for the Products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: Products - Products"))
}

// Product is the resolver for the Product field.
func (r *queryResolver) Product(ctx context.Context, input string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: Product - Product"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
