package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	graph "assignment/internal/graph/generate"
	"assignment/internal/model"
	"context"
	"fmt"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UserRepo.FindById(id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, filter model.UserFilterInput) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
