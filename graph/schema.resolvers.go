package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gitlab.com/dinamchiki/go-graphql/graph/generated"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

// Author is the resolver for the author field.
func (r *articleResolver) Author(ctx context.Context, obj *models.Article) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.AuthorID)
}

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserId)
}

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	return r.Domain.CreateMeetup(ctx, input)
}

// CreatePlace is the resolver for the createPlace field.
func (r *mutationResolver) CreatePlace(ctx context.Context, input models.PlaceInput) (*models.PlacePayload, error) {
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

// ArticlesDelete is the resolver for the articlesDelete field.
func (r *mutationResolver) ArticlesDelete(ctx context.Context, ids []string) ([]*models.ArticlePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ArticlesPublishUpdate is the resolver for the articlesPublishUpdate field.
func (r *mutationResolver) ArticlesPublishUpdate(ctx context.Context, ids []string) ([]*models.ArticlePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ArticlesRestore is the resolver for the articlesRestore field.
func (r *mutationResolver) ArticlesRestore(ctx context.Context, ids []string) ([]*models.ArticlePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ArticlesSave is the resolver for the articlesSave field.
func (r *mutationResolver) ArticlesSave(ctx context.Context, articleInput []*models.ArticleInput) ([]*models.ArticlePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ArticlesUpdate is the resolver for the articlesUpdate field.
func (r *mutationResolver) ArticlesUpdate(ctx context.Context, articleInput []*models.ArticleInputWithID) ([]*models.ArticlePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CartsDelete is the resolver for the cartsDelete field.
func (r *mutationResolver) CartsDelete(ctx context.Context, ids []string) ([]*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CartsPublishUpdate is the resolver for the cartsPublishUpdate field.
func (r *mutationResolver) CartsPublishUpdate(ctx context.Context, ids []string) ([]*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CartsRestore is the resolver for the cartsRestore field.
func (r *mutationResolver) CartsRestore(ctx context.Context, ids []string) ([]*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CartsSave is the resolver for the cartsSave field.
func (r *mutationResolver) CartsSave(ctx context.Context, cartInput []*models.CartInput) ([]*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CartsUpdate is the resolver for the cartsUpdate field.
func (r *mutationResolver) CartsUpdate(ctx context.Context, cartInput []*models.CartInputWithID) ([]*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalancesDelete is the resolver for the clubBalancesDelete field.
func (r *mutationResolver) ClubBalancesDelete(ctx context.Context, ids []string) ([]*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalancesPublishUpdate is the resolver for the clubBalancesPublishUpdate field.
func (r *mutationResolver) ClubBalancesPublishUpdate(ctx context.Context, ids []string) ([]*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalancesRestore is the resolver for the clubBalancesRestore field.
func (r *mutationResolver) ClubBalancesRestore(ctx context.Context, ids []string) ([]*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalancesSave is the resolver for the clubBalancesSave field.
func (r *mutationResolver) ClubBalancesSave(ctx context.Context, clubBalanceInput []*models.ClubBalanceInput) ([]*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalancesUpdate is the resolver for the clubBalancesUpdate field.
func (r *mutationResolver) ClubBalancesUpdate(ctx context.Context, clubBalanceInput []*models.ClubBalanceInputWithID) ([]*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByMonthDelete is the resolver for the coachPaymentsByMonthDelete field.
func (r *mutationResolver) CoachPaymentsByMonthDelete(ctx context.Context, ids []string) ([]*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByMonthPublishUpdate is the resolver for the coachPaymentsByMonthPublishUpdate field.
func (r *mutationResolver) CoachPaymentsByMonthPublishUpdate(ctx context.Context, ids []string) ([]*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByMonthRestore is the resolver for the coachPaymentsByMonthRestore field.
func (r *mutationResolver) CoachPaymentsByMonthRestore(ctx context.Context, ids []string) ([]*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByMonthSave is the resolver for the coachPaymentsByMonthSave field.
func (r *mutationResolver) CoachPaymentsByMonthSave(ctx context.Context, coachPaymentByMonthInput []*models.CoachPaymentByMonthInput) ([]*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByMonthUpdate is the resolver for the coachPaymentsByMonthUpdate field.
func (r *mutationResolver) CoachPaymentsByMonthUpdate(ctx context.Context, coachPaymentByMonthInput []*models.CoachPaymentByMonthInputWithID) ([]*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTeamDelete is the resolver for the coachPaymentsByTeamDelete field.
func (r *mutationResolver) CoachPaymentsByTeamDelete(ctx context.Context, ids []string) ([]*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTeamPublishUpdate is the resolver for the coachPaymentsByTeamPublishUpdate field.
func (r *mutationResolver) CoachPaymentsByTeamPublishUpdate(ctx context.Context, ids []string) ([]*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTeamRestore is the resolver for the coachPaymentsByTeamRestore field.
func (r *mutationResolver) CoachPaymentsByTeamRestore(ctx context.Context, ids []string) ([]*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTeamSave is the resolver for the coachPaymentsByTeamSave field.
func (r *mutationResolver) CoachPaymentsByTeamSave(ctx context.Context, coachPaymentByTeamInput []*models.CoachPaymentByTeamInput) ([]*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTeamUpdate is the resolver for the coachPaymentsByTeamUpdate field.
func (r *mutationResolver) CoachPaymentsByTeamUpdate(ctx context.Context, coachPaymentByTeamInput []*models.CoachPaymentByTeamInputWithID) ([]*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTrainingDelete is the resolver for the coachPaymentsByTrainingDelete field.
func (r *mutationResolver) CoachPaymentsByTrainingDelete(ctx context.Context, ids []string) ([]*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTrainingPublishUpdate is the resolver for the coachPaymentsByTrainingPublishUpdate field.
func (r *mutationResolver) CoachPaymentsByTrainingPublishUpdate(ctx context.Context, ids []string) ([]*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTrainingRestore is the resolver for the coachPaymentsByTrainingRestore field.
func (r *mutationResolver) CoachPaymentsByTrainingRestore(ctx context.Context, ids []string) ([]*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTrainingSave is the resolver for the coachPaymentsByTrainingSave field.
func (r *mutationResolver) CoachPaymentsByTrainingSave(ctx context.Context, coachPaymentByTrainingInput []*models.CoachPaymentByTrainingInput) ([]*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTrainingUpdate is the resolver for the coachPaymentsByTrainingUpdate field.
func (r *mutationResolver) CoachPaymentsByTrainingUpdate(ctx context.Context, coachPaymentByTrainingInput []*models.CoachPaymentByTrainingInputWithID) ([]*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorsDelete is the resolver for the creatorsDelete field.
func (r *mutationResolver) CreatorsDelete(ctx context.Context, ids []string) ([]*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorsPublishUpdate is the resolver for the creatorsPublishUpdate field.
func (r *mutationResolver) CreatorsPublishUpdate(ctx context.Context, ids []string) ([]*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorsRestore is the resolver for the creatorsRestore field.
func (r *mutationResolver) CreatorsRestore(ctx context.Context, ids []string) ([]*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorsSave is the resolver for the creatorsSave field.
func (r *mutationResolver) CreatorsSave(ctx context.Context, creatorInput []*models.CreatorInput) ([]*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorsUpdate is the resolver for the creatorsUpdate field.
func (r *mutationResolver) CreatorsUpdate(ctx context.Context, creatorInput []*models.CreatorInputWithID) ([]*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitsDelete is the resolver for the kitsDelete field.
func (r *mutationResolver) KitsDelete(ctx context.Context, ids []string) ([]*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitsPublishUpdate is the resolver for the kitsPublishUpdate field.
func (r *mutationResolver) KitsPublishUpdate(ctx context.Context, ids []string) ([]*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitsRestore is the resolver for the kitsRestore field.
func (r *mutationResolver) KitsRestore(ctx context.Context, ids []string) ([]*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitsSave is the resolver for the kitsSave field.
func (r *mutationResolver) KitsSave(ctx context.Context, kitInput []*models.KitInput) ([]*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitsUpdate is the resolver for the kitsUpdate field.
func (r *mutationResolver) KitsUpdate(ctx context.Context, kitInput []*models.KitInputWithID) ([]*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadsDelete is the resolver for the leadsDelete field.
func (r *mutationResolver) LeadsDelete(ctx context.Context, ids []string) ([]*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadsPublishUpdate is the resolver for the leadsPublishUpdate field.
func (r *mutationResolver) LeadsPublishUpdate(ctx context.Context, ids []string) ([]*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadsRestore is the resolver for the leadsRestore field.
func (r *mutationResolver) LeadsRestore(ctx context.Context, ids []string) ([]*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadsSave is the resolver for the leadsSave field.
func (r *mutationResolver) LeadsSave(ctx context.Context, leadInput []*models.LeadInput) ([]*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadsUpdate is the resolver for the leadsUpdate field.
func (r *mutationResolver) LeadsUpdate(ctx context.Context, leadInput []*models.LeadInputWithID) ([]*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostsDelete is the resolver for the moneyCostsDelete field.
func (r *mutationResolver) MoneyCostsDelete(ctx context.Context, ids []string) ([]*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostsPublishUpdate is the resolver for the moneyCostsPublishUpdate field.
func (r *mutationResolver) MoneyCostsPublishUpdate(ctx context.Context, ids []string) ([]*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostsRestore is the resolver for the moneyCostsRestore field.
func (r *mutationResolver) MoneyCostsRestore(ctx context.Context, ids []string) ([]*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostsSave is the resolver for the moneyCostsSave field.
func (r *mutationResolver) MoneyCostsSave(ctx context.Context, moneyCostInput []*models.MoneyCostInput) ([]*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostsUpdate is the resolver for the moneyCostsUpdate field.
func (r *mutationResolver) MoneyCostsUpdate(ctx context.Context, moneyCostInput []*models.MoneyCostInputWithID) ([]*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMovesDelete is the resolver for the moneyMovesDelete field.
func (r *mutationResolver) MoneyMovesDelete(ctx context.Context, ids []string) ([]*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMovesPublishUpdate is the resolver for the moneyMovesPublishUpdate field.
func (r *mutationResolver) MoneyMovesPublishUpdate(ctx context.Context, ids []string) ([]*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMovesRestore is the resolver for the moneyMovesRestore field.
func (r *mutationResolver) MoneyMovesRestore(ctx context.Context, ids []string) ([]*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMovesSave is the resolver for the moneyMovesSave field.
func (r *mutationResolver) MoneyMovesSave(ctx context.Context, moneyMoveInput []*models.MoneyMoveInput) ([]*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMovesUpdate is the resolver for the moneyMovesUpdate field.
func (r *mutationResolver) MoneyMovesUpdate(ctx context.Context, moneyMoveInput []*models.MoneyMoveInputWithID) ([]*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrdersDelete is the resolver for the ordersDelete field.
func (r *mutationResolver) OrdersDelete(ctx context.Context, ids []string) ([]*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrdersPublishUpdate is the resolver for the ordersPublishUpdate field.
func (r *mutationResolver) OrdersPublishUpdate(ctx context.Context, ids []string) ([]*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrdersRestore is the resolver for the ordersRestore field.
func (r *mutationResolver) OrdersRestore(ctx context.Context, ids []string) ([]*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrdersSave is the resolver for the ordersSave field.
func (r *mutationResolver) OrdersSave(ctx context.Context, orderInput []*models.OrderInput) ([]*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrdersUpdate is the resolver for the ordersUpdate field.
func (r *mutationResolver) OrdersUpdate(ctx context.Context, orderInput []*models.OrderInputWithID) ([]*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// PlacesDelete is the resolver for the placesDelete field.
func (r *mutationResolver) PlacesDelete(ctx context.Context, ids []string) ([]*models.PlacePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// PlacesPublishUpdate is the resolver for the placesPublishUpdate field.
func (r *mutationResolver) PlacesPublishUpdate(ctx context.Context, ids []string) ([]*models.PlacePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// PlacesRestore is the resolver for the placesRestore field.
func (r *mutationResolver) PlacesRestore(ctx context.Context, ids []string) ([]*models.PlacePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// PlacesSave is the resolver for the placesSave field.
func (r *mutationResolver) PlacesSave(ctx context.Context, placeInput []*models.PlaceInput) ([]*models.PlacePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// PlacesUpdate is the resolver for the placesUpdate field.
func (r *mutationResolver) PlacesUpdate(ctx context.Context, placeInput []*models.PlaceInputWithID) ([]*models.PlacePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Refresh is the resolver for the refresh field.
func (r *mutationResolver) Refresh(ctx context.Context, phone string, token string) (*models.Token, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthDelete is the resolver for the rentPaymentByMonthDelete field.
func (r *mutationResolver) RentPaymentByMonthDelete(ctx context.Context, ids []string) ([]*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthPublishUpdate is the resolver for the rentPaymentByMonthPublishUpdate field.
func (r *mutationResolver) RentPaymentByMonthPublishUpdate(ctx context.Context, ids []string) ([]*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthRestore is the resolver for the rentPaymentByMonthRestore field.
func (r *mutationResolver) RentPaymentByMonthRestore(ctx context.Context, ids []string) ([]*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthSave is the resolver for the rentPaymentByMonthSave field.
func (r *mutationResolver) RentPaymentByMonthSave(ctx context.Context, rentPaymentInput []*models.RentPaymentByMonthInput) ([]*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthUpdate is the resolver for the rentPaymentByMonthUpdate field.
func (r *mutationResolver) RentPaymentByMonthUpdate(ctx context.Context, rentPaymentInput []*models.RentPaymentByMonthInputWithID) ([]*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingDelete is the resolver for the rentPaymentByTrainingDelete field.
func (r *mutationResolver) RentPaymentByTrainingDelete(ctx context.Context, ids []string) ([]*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingPublishUpdate is the resolver for the rentPaymentByTrainingPublishUpdate field.
func (r *mutationResolver) RentPaymentByTrainingPublishUpdate(ctx context.Context, ids []string) ([]*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingRestore is the resolver for the rentPaymentByTrainingRestore field.
func (r *mutationResolver) RentPaymentByTrainingRestore(ctx context.Context, ids []string) ([]*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingSave is the resolver for the rentPaymentByTrainingSave field.
func (r *mutationResolver) RentPaymentByTrainingSave(ctx context.Context, rentPaymentInput []*models.RentPaymentByTrainingInput) ([]*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingUpdate is the resolver for the rentPaymentByTrainingUpdate field.
func (r *mutationResolver) RentPaymentByTrainingUpdate(ctx context.Context, rentPaymentInput []*models.RentPaymentByTrainingInputWithID) ([]*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StadiumsDelete is the resolver for the stadiumsDelete field.
func (r *mutationResolver) StadiumsDelete(ctx context.Context, ids []string) ([]*models.StadiumPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StadiumsPublishUpdate is the resolver for the stadiumsPublishUpdate field.
func (r *mutationResolver) StadiumsPublishUpdate(ctx context.Context, ids []string) ([]*models.StadiumPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StadiumsRestore is the resolver for the stadiumsRestore field.
func (r *mutationResolver) StadiumsRestore(ctx context.Context, ids []string) ([]*models.StadiumPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StadiumsSave is the resolver for the stadiumsSave field.
func (r *mutationResolver) StadiumsSave(ctx context.Context, stadiumInput []*models.StadiumInput) ([]*models.StadiumPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StadiumsUpdate is the resolver for the stadiumsUpdate field.
func (r *mutationResolver) StadiumsUpdate(ctx context.Context, stadiumInput []*models.StadiumInputWithID) ([]*models.StadiumPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffDelete is the resolver for the staffDelete field.
func (r *mutationResolver) StaffDelete(ctx context.Context, ids []string) ([]*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffPublishUpdate is the resolver for the staffPublishUpdate field.
func (r *mutationResolver) StaffPublishUpdate(ctx context.Context, ids []string) ([]*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffRestore is the resolver for the staffRestore field.
func (r *mutationResolver) StaffRestore(ctx context.Context, ids []string) ([]*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffSave is the resolver for the staffSave field.
func (r *mutationResolver) StaffSave(ctx context.Context, staffInput []*models.StaffInput) ([]*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffUpdate is the resolver for the staffUpdate field.
func (r *mutationResolver) StaffUpdate(ctx context.Context, staffInput []*models.StaffInputWithID) ([]*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitsDelete is the resolver for the studentVisitsDelete field.
func (r *mutationResolver) StudentVisitsDelete(ctx context.Context, ids []string) ([]*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitsPublishUpdate is the resolver for the studentVisitsPublishUpdate field.
func (r *mutationResolver) StudentVisitsPublishUpdate(ctx context.Context, ids []string) ([]*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitsRestore is the resolver for the studentVisitsRestore field.
func (r *mutationResolver) StudentVisitsRestore(ctx context.Context, ids []string) ([]*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitsSave is the resolver for the studentVisitsSave field.
func (r *mutationResolver) StudentVisitsSave(ctx context.Context, studentVisitInput []*models.StudentVisitInput) ([]*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitsUpdate is the resolver for the studentVisitsUpdate field.
func (r *mutationResolver) StudentVisitsUpdate(ctx context.Context, studentVisitInput []*models.StudentVisitInputWithID) ([]*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentsDelete is the resolver for the studentsDelete field.
func (r *mutationResolver) StudentsDelete(ctx context.Context, ids []string) ([]*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentsPublishUpdate is the resolver for the studentsPublishUpdate field.
func (r *mutationResolver) StudentsPublishUpdate(ctx context.Context, ids []string) ([]*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentsRestore is the resolver for the studentsRestore field.
func (r *mutationResolver) StudentsRestore(ctx context.Context, ids []string) ([]*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentsSave is the resolver for the studentsSave field.
func (r *mutationResolver) StudentsSave(ctx context.Context, studentInput []*models.StudentInput) ([]*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentsUpdate is the resolver for the studentsUpdate field.
func (r *mutationResolver) StudentsUpdate(ctx context.Context, studentInput []*models.StudentInputWithID) ([]*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TasksDelete is the resolver for the tasksDelete field.
func (r *mutationResolver) TasksDelete(ctx context.Context, ids []string) ([]*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TasksPublishUpdate is the resolver for the tasksPublishUpdate field.
func (r *mutationResolver) TasksPublishUpdate(ctx context.Context, ids []string) ([]*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TasksRestore is the resolver for the tasksRestore field.
func (r *mutationResolver) TasksRestore(ctx context.Context, ids []string) ([]*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TasksSave is the resolver for the tasksSave field.
func (r *mutationResolver) TasksSave(ctx context.Context, taskInput []*models.TaskInput) ([]*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TasksUpdate is the resolver for the tasksUpdate field.
func (r *mutationResolver) TasksUpdate(ctx context.Context, taskInput []*models.TaskInputWithID) ([]*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalancesDelete is the resolver for the teamBalancesDelete field.
func (r *mutationResolver) TeamBalancesDelete(ctx context.Context, ids []string) ([]*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalancesPublishUpdate is the resolver for the teamBalancesPublishUpdate field.
func (r *mutationResolver) TeamBalancesPublishUpdate(ctx context.Context, ids []string) ([]*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalancesRestore is the resolver for the teamBalancesRestore field.
func (r *mutationResolver) TeamBalancesRestore(ctx context.Context, ids []string) ([]*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalancesSave is the resolver for the teamBalancesSave field.
func (r *mutationResolver) TeamBalancesSave(ctx context.Context, teamBalanceInput []*models.TeamBalanceInput) ([]*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalancesUpdate is the resolver for the teamBalancesUpdate field.
func (r *mutationResolver) TeamBalancesUpdate(ctx context.Context, teamBalanceInput []*models.TeamBalanceInputWithID) ([]*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamsDelete is the resolver for the teamsDelete field.
func (r *mutationResolver) TeamsDelete(ctx context.Context, ids []string) ([]*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamsPublishUpdate is the resolver for the teamsPublishUpdate field.
func (r *mutationResolver) TeamsPublishUpdate(ctx context.Context, ids []string) ([]*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamsRestore is the resolver for the teamsRestore field.
func (r *mutationResolver) TeamsRestore(ctx context.Context, ids []string) ([]*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamsSave is the resolver for the teamsSave field.
func (r *mutationResolver) TeamsSave(ctx context.Context, teamInput []*models.TeamInput) ([]*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamsUpdate is the resolver for the teamsUpdate field.
func (r *mutationResolver) TeamsUpdate(ctx context.Context, teamInput []*models.TeamInputWithID) ([]*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDaysDelete is the resolver for the trainingDaysDelete field.
func (r *mutationResolver) TrainingDaysDelete(ctx context.Context, ids []string) ([]*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDaysPublishUpdate is the resolver for the trainingDaysPublishUpdate field.
func (r *mutationResolver) TrainingDaysPublishUpdate(ctx context.Context, ids []string) ([]*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDaysRestore is the resolver for the trainingDaysRestore field.
func (r *mutationResolver) TrainingDaysRestore(ctx context.Context, ids []string) ([]*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDaysSave is the resolver for the trainingDaysSave field.
func (r *mutationResolver) TrainingDaysSave(ctx context.Context, trainingDayInput []*models.TrainingDayInput) ([]*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDaysUpdate is the resolver for the trainingDaysUpdate field.
func (r *mutationResolver) TrainingDaysUpdate(ctx context.Context, trainingDayInput []*models.TrainingDayInputWithID) ([]*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingsDelete is the resolver for the trainingsDelete field.
func (r *mutationResolver) TrainingsDelete(ctx context.Context, ids []string) ([]*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingsPublishUpdate is the resolver for the trainingsPublishUpdate field.
func (r *mutationResolver) TrainingsPublishUpdate(ctx context.Context, ids []string) ([]*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingsRestore is the resolver for the trainingsRestore field.
func (r *mutationResolver) TrainingsRestore(ctx context.Context, ids []string) ([]*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingsSave is the resolver for the trainingsSave field.
func (r *mutationResolver) TrainingsSave(ctx context.Context, trainingInput []*models.TrainingInput) ([]*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingsUpdate is the resolver for the trainingsUpdate field.
func (r *mutationResolver) TrainingsUpdate(ctx context.Context, trainingInput []*models.TrainingInputWithID) ([]*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UsersDelete is the resolver for the usersDelete field.
func (r *mutationResolver) UsersDelete(ctx context.Context, ids []string) ([]*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UsersPublishUpdate is the resolver for the usersPublishUpdate field.
func (r *mutationResolver) UsersPublishUpdate(ctx context.Context, ids []string) ([]*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UsersRestore is the resolver for the usersRestore field.
func (r *mutationResolver) UsersRestore(ctx context.Context, ids []string) ([]*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UsersSave is the resolver for the usersSave field.
func (r *mutationResolver) UsersSave(ctx context.Context, userInput []*models.UserInput) ([]*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UsersUpdate is the resolver for the usersUpdate field.
func (r *mutationResolver) UsersUpdate(ctx context.Context, userInput []*models.UserInputWithID) ([]*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {
	return r.Domain.MeetupsRepo.GetMeetups(filter, limit, offset)
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, id string) (*models.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context, after *string, before *string, first *int, last *int) (*models.ArticleConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context, id string) (*models.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

// Carts is the resolver for the carts field.
func (r *queryResolver) Carts(ctx context.Context, after *string, before *string, first *int, last *int) (*models.CartConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalance is the resolver for the clubBalance field.
func (r *queryResolver) ClubBalance(ctx context.Context, id string) (*models.ClubBalance, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalances is the resolver for the clubBalances field.
func (r *queryResolver) ClubBalances(ctx context.Context, after *string, before *string, first *int, last *int) (*models.ClubBalanceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByMonth is the resolver for the coachPaymentByMonth field.
func (r *queryResolver) CoachPaymentByMonth(ctx context.Context, id string) (*models.CoachPaymentByMonth, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTeam is the resolver for the coachPaymentByTeam field.
func (r *queryResolver) CoachPaymentByTeam(ctx context.Context, id string) (*models.CoachPaymentByTeam, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTraining is the resolver for the coachPaymentByTraining field.
func (r *queryResolver) CoachPaymentByTraining(ctx context.Context, id string) (*models.CoachPaymentByTraining, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByMonth is the resolver for the coachPaymentsByMonth field.
func (r *queryResolver) CoachPaymentsByMonth(ctx context.Context, after *string, before *string, date time.Time, first *int, last *int) (*models.CoachPaymentByMonthConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTeam is the resolver for the coachPaymentsByTeam field.
func (r *queryResolver) CoachPaymentsByTeam(ctx context.Context, after *string, before *string, first *int, last *int, team *models.TeamDto) (*models.CoachPaymentByTeamConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTraining is the resolver for the coachPaymentsByTraining field.
func (r *queryResolver) CoachPaymentsByTraining(ctx context.Context, after *string, before *string, date time.Time, first *int, last *int) (*models.CoachPaymentByTrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Creator is the resolver for the creator field.
func (r *queryResolver) Creator(ctx context.Context, id *string) (*models.Creator, error) {
	panic(fmt.Errorf("not implemented"))
}

// Creators is the resolver for the creators field.
func (r *queryResolver) Creators(ctx context.Context, after *string, before *string, first *int, last *int) (*models.CreatorConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Kit is the resolver for the kit field.
func (r *queryResolver) Kit(ctx context.Context, id string) (*models.Kit, error) {
	panic(fmt.Errorf("not implemented"))
}

// Kits is the resolver for the kits field.
func (r *queryResolver) Kits(ctx context.Context, after *string, before *string, first *int, last *int) (*models.KitConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Lead is the resolver for the lead field.
func (r *queryResolver) Lead(ctx context.Context, id string) (*models.Lead, error) {
	panic(fmt.Errorf("not implemented"))
}

// Leads is the resolver for the leads field.
func (r *queryResolver) Leads(ctx context.Context, after *string, before *string, first *int, last *int) (*models.LeadConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCost is the resolver for the moneyCost field.
func (r *queryResolver) MoneyCost(ctx context.Context, id string) (*models.MoneyCost, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCosts is the resolver for the moneyCosts field.
func (r *queryResolver) MoneyCosts(ctx context.Context, after *string, before *string, first *int, last *int) (*models.MoneyCostConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMove is the resolver for the moneyMove field.
func (r *queryResolver) MoneyMove(ctx context.Context, id string) (*models.MoneyMove, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMoves is the resolver for the moneyMoves field.
func (r *queryResolver) MoneyMoves(ctx context.Context, after *string, before *string, first *int, last *int) (*models.MoneyMoveConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, after *string, before *string, first *int, last *int) (*models.OrderConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Place is the resolver for the place field.
func (r *queryResolver) Place(ctx context.Context, id string) (*models.Place, error) {
	panic(fmt.Errorf("not implemented"))
}

// Places is the resolver for the places field.
func (r *queryResolver) Places(ctx context.Context, filter *models.PlaceFilter, first *int, after *string, last *int, before *string) (*models.PlaceConnection, error) {
	return r.Domain.PlacesRepo.GetPlaces(filter, first, last, after, before)
}

// RentPaymentByMonth is the resolver for the rentPaymentByMonth field.
func (r *queryResolver) RentPaymentByMonth(ctx context.Context, id string) (*models.RentPaymentByMonth, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTraining is the resolver for the rentPaymentByTraining field.
func (r *queryResolver) RentPaymentByTraining(ctx context.Context, id string) (*models.RentPaymentByTraining, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentsByMonth is the resolver for the rentPaymentsByMonth field.
func (r *queryResolver) RentPaymentsByMonth(ctx context.Context, after *string, before *string, first *int, last *int) (*models.RentPaymentByMonthConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentsByTraining is the resolver for the rentPaymentsByTraining field.
func (r *queryResolver) RentPaymentsByTraining(ctx context.Context, after *string, before *string, first *int, last *int) (*models.RentPaymentByTrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Stadium is the resolver for the stadium field.
func (r *queryResolver) Stadium(ctx context.Context, id string) (*models.Stadium, error) {
	panic(fmt.Errorf("not implemented"))
}

// Stadiums is the resolver for the stadiums field.
func (r *queryResolver) Stadiums(ctx context.Context, after *string, before *string, first *int, last *int) (*models.StadiumConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Staff is the resolver for the staff field.
func (r *queryResolver) Staff(ctx context.Context, after *string, before *string, first *int, last *int) (*models.StaffConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffPerson is the resolver for the staffPerson field.
func (r *queryResolver) StaffPerson(ctx context.Context, id string) (*models.Staff, error) {
	panic(fmt.Errorf("not implemented"))
}

// Student is the resolver for the student field.
func (r *queryResolver) Student(ctx context.Context, id string) (*models.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisit is the resolver for the studentVisit field.
func (r *queryResolver) StudentVisit(ctx context.Context, id string) (*models.StudentVisit, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisits is the resolver for the studentVisits field.
func (r *queryResolver) StudentVisits(ctx context.Context, after *string, before *string, first *int, last *int) (*models.StudentVisitConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Students is the resolver for the students field.
func (r *queryResolver) Students(ctx context.Context, after *string, before *string, first *int, last *int) (*models.StudentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*models.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// NearestStudentBirthdays is the resolver for the nearestStudentBirthdays field.
func (r *queryResolver) NearestStudentBirthdays(ctx context.Context, offset int, first *int, last *int, after *string, before *string) (*models.StudentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// NearestStaffBirthdays is the resolver for the nearestStaffBirthdays field.
func (r *queryResolver) NearestStaffBirthdays(ctx context.Context, offset int, first *int, last *int, after *string, before *string) (*models.StaffConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// UnPayedStudents is the resolver for the unPayedStudents field.
func (r *queryResolver) UnPayedStudents(ctx context.Context, first *int, last *int, after *string, before *string) (*models.StudentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// CurrentTasks is the resolver for the currentTasks field.
func (r *queryResolver) CurrentTasks(ctx context.Context, first *int, last *int, after *string, before *string) (*models.TaskConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// TimeTable is the resolver for the timeTable field.
func (r *queryResolver) TimeTable(ctx context.Context, startDay *time.Time, first *int, last *int, after *string, before *string) (*models.TrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingsByDay is the resolver for the trainingsByDay field.
func (r *queryResolver) TrainingsByDay(ctx context.Context, date *time.Time, first *int, last *int, after *string, before *string) (*models.TrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, after *string, before *string, first *int, last *int) (*models.TaskConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Team is the resolver for the team field.
func (r *queryResolver) Team(ctx context.Context, id string) (*models.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalance is the resolver for the teamBalance field.
func (r *queryResolver) TeamBalance(ctx context.Context, id string) (*models.TeamBalance, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalances is the resolver for the teamBalances field.
func (r *queryResolver) TeamBalances(ctx context.Context, after *string, before *string, first *int, last *int) (*models.TeamBalanceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Teams is the resolver for the teams field.
func (r *queryResolver) Teams(ctx context.Context, after *string, before *string, first *int, last *int) (*models.TeamConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Training is the resolver for the training field.
func (r *queryResolver) Training(ctx context.Context, id string) (*models.Training, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDay is the resolver for the trainingDay field.
func (r *queryResolver) TrainingDay(ctx context.Context, id string) (*models.TrainingDay, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDays is the resolver for the trainingDays field.
func (r *queryResolver) TrainingDays(ctx context.Context, after *string, before *string, first *int, last *int) (*models.TrainingDayConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Trainings is the resolver for the trainings field.
func (r *queryResolver) Trainings(ctx context.Context, after *string, before *string, first *int, last *int) (*models.TrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.Domain.UsersRepo.GetUserByID(id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *string, before *string, first *int, last *int) (*models.UserConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return r.Domain.MeetupsRepo.GetMeetupsForUser(obj)
}

// Article returns generated.ArticleResolver implementation.
func (r *Resolver) Article() generated.ArticleResolver { return &articleResolver{r} }

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type articleResolver struct{ *Resolver }
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
