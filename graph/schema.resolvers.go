package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/allgurgel/FC-GO-EXPERT-DESAFIO-CLEAN-ARCH/internal/infra/graph"
	"github.com/allgurgel/FC-GO-EXPERT-DESAFIO-CLEAN-ARCH/internal/infra/graph/model"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// ListOrders is the resolver for the listOrders field.
func (r *queryResolver) ListOrders(ctx context.Context, input *model.ListOrderInput) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: ListOrders - listOrders"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
