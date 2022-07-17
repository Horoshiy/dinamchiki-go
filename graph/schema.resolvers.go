package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/dinamchiki/go-graphql/graph/generated"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserId)
}

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	return r.Domain.CreateMeetup(ctx, input)
}

// CreatePlace is the resolver for the createPlace field.
func (r *mutationResolver) CreatePlace(ctx context.Context, input models.PlaceInput) (*models.Place, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CreatePlace(input)
}

// UpdateMeetup is the resolver for the updateMeetup field.
func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input *models.UpdateMeetup) (*models.Meetup, error) {
	return r.Domain.UpdateMeetup(ctx, id, input)
}

// UpdatePlace is the resolver for the updatePlace field.
func (r *mutationResolver) UpdatePlace(ctx context.Context, id string, input *models.PlaceInput) (*models.Place, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteMeetup is the resolver for the deleteMeetup field.
func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return r.Domain.DeleteMeetup(ctx, id)
}

// DeletePlace is the resolver for the deletePlace field.
func (r *mutationResolver) DeletePlace(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}

	return r.Domain.Register(ctx, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}

	return r.Domain.Login(ctx, input)
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	return r.Domain.MeetupsRepo.GetMeetups(filter, limit, offset)
}

// Places is the resolver for the places field.
func (r *queryResolver) Places(ctx context.Context, filter *models.PlaceFilter, limit *int, offset *int) ([]*models.Place, error) {
	return r.Domain.PlacesRepo.GetPlaces(filter, limit, offset)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.Domain.UsersRepo.GetUserByID(id)
}

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return r.Domain.MeetupsRepo.GetMeetupsForUser(obj)
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	ErrInput = errors.New("inputs errors")
)
