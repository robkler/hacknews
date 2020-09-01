package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "github.com/robkler/hacknews/graph/generated"
	model1 "github.com/robkler/hacknews/graph/model"
	"github.com/robkler/hacknews/pkg/links"
	"strconv"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model1.NewLink) (*model1.Link, error) {
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	linkID := link.Save()
	return &model1.Link{ID: strconv.FormatInt(linkID, 10), Title:link.Title, Address:link.Address}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model1.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model1.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model1.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model1.Link, error) {
	var resultLinks []*model1.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks{
		resultLinks = append(resultLinks, &model1.Link{ID:link.ID, Title:link.Title, Address:link.Address})
	}
	return resultLinks, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
