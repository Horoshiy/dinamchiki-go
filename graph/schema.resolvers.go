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

// Author is the resolver for the author field.
func (r *articleResolver) Author(ctx context.Context, obj *models.Article) (*models.User, error) {
	return For(ctx).UserLoader.Load(obj.AuthorID)
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

// ArticleDelete is the resolver for the articleDelete field.
func (r *mutationResolver) ArticleDelete(ctx context.Context, id string) (*models.ArticlePayload, error) {
	_, err := r.Domain.ArticleDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.ArticlePayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// ArticlePublishUpdate is the resolver for the articlePublishUpdate field.
func (r *mutationResolver) ArticlePublishUpdate(ctx context.Context, id string) (*models.ArticlePayload, error) {
	return r.Domain.ArticlePublish(id)
}

// ArticleSave is the resolver for the articleSave field.
func (r *mutationResolver) ArticleSave(ctx context.Context, articleInput models.ArticleInput) (*models.ArticlePayload, error) {
	isValid := validation(ctx, articleInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.ArticleSave(ctx, articleInput)
}

// ArticleUpdate is the resolver for the articleUpdate field.
func (r *mutationResolver) ArticleUpdate(ctx context.Context, articleInput models.ArticleInputWithID) (*models.ArticlePayload, error) {
	isValid := validation(ctx, articleInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.ArticleUpdate(ctx, articleInput)
}

// CartDelete is the resolver for the cartDelete field.
func (r *mutationResolver) CartDelete(ctx context.Context, id string) (*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CartPublishUpdate is the resolver for the cartPublishUpdate field.
func (r *mutationResolver) CartPublishUpdate(ctx context.Context, id string) (*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CartSave is the resolver for the cartSave field.
func (r *mutationResolver) CartSave(ctx context.Context, cartInput models.CartInput) (*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CartUpdate is the resolver for the cartUpdate field.
func (r *mutationResolver) CartUpdate(ctx context.Context, cartInput models.CartInputWithID) (*models.CartPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalanceDelete is the resolver for the clubBalanceDelete field.
func (r *mutationResolver) ClubBalanceDelete(ctx context.Context, id string) (*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalancePublishUpdate is the resolver for the clubBalancePublishUpdate field.
func (r *mutationResolver) ClubBalancePublishUpdate(ctx context.Context, id string) (*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalanceSave is the resolver for the clubBalanceSave field.
func (r *mutationResolver) ClubBalanceSave(ctx context.Context, clubBalanceInput models.ClubBalanceInput) (*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalanceUpdate is the resolver for the clubBalanceUpdate field.
func (r *mutationResolver) ClubBalanceUpdate(ctx context.Context, clubBalanceInput models.ClubBalanceInputWithID) (*models.ClubBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByMonthDelete is the resolver for the coachPaymentByMonthDelete field.
func (r *mutationResolver) CoachPaymentByMonthDelete(ctx context.Context, id string) (*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByMonthPublishUpdate is the resolver for the coachPaymentByMonthPublishUpdate field.
func (r *mutationResolver) CoachPaymentByMonthPublishUpdate(ctx context.Context, id string) (*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByMonthSave is the resolver for the coachPaymentByMonthSave field.
func (r *mutationResolver) CoachPaymentByMonthSave(ctx context.Context, coachPaymentByMonthInput models.CoachPaymentByMonthInput) (*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByMonthUpdate is the resolver for the coachPaymentByMonthUpdate field.
func (r *mutationResolver) CoachPaymentByMonthUpdate(ctx context.Context, coachPaymentByMonthInput models.CoachPaymentByMonthInputWithID) (*models.CoachPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTeamDelete is the resolver for the coachPaymentByTeamDelete field.
func (r *mutationResolver) CoachPaymentByTeamDelete(ctx context.Context, id string) (*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTeamPublishUpdate is the resolver for the coachPaymentByTeamPublishUpdate field.
func (r *mutationResolver) CoachPaymentByTeamPublishUpdate(ctx context.Context, id string) (*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTeamSave is the resolver for the coachPaymentByTeamSave field.
func (r *mutationResolver) CoachPaymentByTeamSave(ctx context.Context, coachPaymentByTeamInput models.CoachPaymentByTeamInput) (*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTeamUpdate is the resolver for the coachPaymentByTeamUpdate field.
func (r *mutationResolver) CoachPaymentByTeamUpdate(ctx context.Context, coachPaymentByTeamInput models.CoachPaymentByTeamInputWithID) (*models.CoachPaymentByTeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTrainingDelete is the resolver for the coachPaymentByTrainingDelete field.
func (r *mutationResolver) CoachPaymentByTrainingDelete(ctx context.Context, id string) (*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTrainingPublishUpdate is the resolver for the coachPaymentByTrainingPublishUpdate field.
func (r *mutationResolver) CoachPaymentByTrainingPublishUpdate(ctx context.Context, id string) (*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTrainingSave is the resolver for the coachPaymentByTrainingSave field.
func (r *mutationResolver) CoachPaymentByTrainingSave(ctx context.Context, coachPaymentByTrainingInput models.CoachPaymentByTrainingInput) (*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentByTrainingUpdate is the resolver for the coachPaymentByTrainingUpdate field.
func (r *mutationResolver) CoachPaymentByTrainingUpdate(ctx context.Context, coachPaymentByTrainingInput models.CoachPaymentByTrainingInputWithID) (*models.CoachPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorDelete is the resolver for the creatorDelete field.
func (r *mutationResolver) CreatorDelete(ctx context.Context, id string) (*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorPublishUpdate is the resolver for the creatorPublishUpdate field.
func (r *mutationResolver) CreatorPublishUpdate(ctx context.Context, id string) (*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorSave is the resolver for the creatorSave field.
func (r *mutationResolver) CreatorSave(ctx context.Context, creatorInput models.CreatorInput) (*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatorUpdate is the resolver for the creatorUpdate field.
func (r *mutationResolver) CreatorUpdate(ctx context.Context, creatorInput models.CreatorInputWithID) (*models.CreatorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitDelete is the resolver for the kitDelete field.
func (r *mutationResolver) KitDelete(ctx context.Context, id string) (*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitPublishUpdate is the resolver for the kitPublishUpdate field.
func (r *mutationResolver) KitPublishUpdate(ctx context.Context, id string) (*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitSave is the resolver for the kitSave field.
func (r *mutationResolver) KitSave(ctx context.Context, kitInput models.KitInput) (*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// KitUpdate is the resolver for the kitUpdate field.
func (r *mutationResolver) KitUpdate(ctx context.Context, kitInput models.KitInputWithID) (*models.KitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadDelete is the resolver for the leadDelete field.
func (r *mutationResolver) LeadDelete(ctx context.Context, id string) (*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadPublishUpdate is the resolver for the leadPublishUpdate field.
func (r *mutationResolver) LeadPublishUpdate(ctx context.Context, id string) (*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadSave is the resolver for the leadSave field.
func (r *mutationResolver) LeadSave(ctx context.Context, leadInput models.LeadInput) (*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// LeadUpdate is the resolver for the leadUpdate field.
func (r *mutationResolver) LeadUpdate(ctx context.Context, leadInput models.LeadInputWithID) (*models.LeadPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostDelete is the resolver for the moneyCostDelete field.
func (r *mutationResolver) MoneyCostDelete(ctx context.Context, id string) (*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostPublishUpdate is the resolver for the moneyCostPublishUpdate field.
func (r *mutationResolver) MoneyCostPublishUpdate(ctx context.Context, id string) (*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostSave is the resolver for the moneyCostSave field.
func (r *mutationResolver) MoneyCostSave(ctx context.Context, moneyCostInput models.MoneyCostInput) (*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCostUpdate is the resolver for the moneyCostUpdate field.
func (r *mutationResolver) MoneyCostUpdate(ctx context.Context, moneyCostInput models.MoneyCostInputWithID) (*models.MoneyCostPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMoveDelete is the resolver for the moneyMoveDelete field.
func (r *mutationResolver) MoneyMoveDelete(ctx context.Context, id string) (*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMovePublishUpdate is the resolver for the moneyMovePublishUpdate field.
func (r *mutationResolver) MoneyMovePublishUpdate(ctx context.Context, id string) (*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMoveSave is the resolver for the moneyMoveSave field.
func (r *mutationResolver) MoneyMoveSave(ctx context.Context, moneyMoveInput models.MoneyMoveInput) (*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMoveUpdate is the resolver for the moneyMoveUpdate field.
func (r *mutationResolver) MoneyMoveUpdate(ctx context.Context, moneyMoveInput models.MoneyMoveInputWithID) (*models.MoneyMovePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrderDelete is the resolver for the orderDelete field.
func (r *mutationResolver) OrderDelete(ctx context.Context, id string) (*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrderPublishUpdate is the resolver for the orderPublishUpdate field.
func (r *mutationResolver) OrderPublishUpdate(ctx context.Context, id string) (*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrderSave is the resolver for the orderSave field.
func (r *mutationResolver) OrderSave(ctx context.Context, orderInput models.OrderInput) (*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// OrderUpdate is the resolver for the orderUpdate field.
func (r *mutationResolver) OrderUpdate(ctx context.Context, orderInput models.OrderInputWithID) (*models.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// PlaceDelete is the resolver for the placeDelete field.
func (r *mutationResolver) PlaceDelete(ctx context.Context, id string) (*models.PlacePayload, error) {
	_, err := r.Domain.PlaceDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.PlacePayload{
		ID:    id,
		Place: nil,
	}, nil
}

// PlacePublishUpdate is the resolver for the placePublishUpdate field.
func (r *mutationResolver) PlacePublishUpdate(ctx context.Context, id string) (*models.PlacePayload, error) {
	return r.Domain.PlacePublish(id)
}

// PlaceSave is the resolver for the placeSave field.
func (r *mutationResolver) PlaceSave(ctx context.Context, placeInput models.PlaceInput) (*models.PlacePayload, error) {
	isValid := validation(ctx, placeInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.PlaceSave(placeInput)
}

// PlaceUpdate is the resolver for the placeUpdate field.
func (r *mutationResolver) PlaceUpdate(ctx context.Context, placeInput models.PlaceInputWithID) (*models.PlacePayload, error) {
	isValid := validation(ctx, placeInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.PlaceUpdate(ctx, placeInput)
}

// Refresh is the resolver for the refresh field.
func (r *mutationResolver) Refresh(ctx context.Context, phone string, token string) (*models.Token, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthDelete is the resolver for the rentPaymentByMonthDelete field.
func (r *mutationResolver) RentPaymentByMonthDelete(ctx context.Context, id string) (*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthPublishUpdate is the resolver for the rentPaymentByMonthPublishUpdate field.
func (r *mutationResolver) RentPaymentByMonthPublishUpdate(ctx context.Context, id string) (*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthSave is the resolver for the rentPaymentByMonthSave field.
func (r *mutationResolver) RentPaymentByMonthSave(ctx context.Context, rentPaymentInput models.RentPaymentByMonthInput) (*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByMonthUpdate is the resolver for the rentPaymentByMonthUpdate field.
func (r *mutationResolver) RentPaymentByMonthUpdate(ctx context.Context, rentPaymentInput models.RentPaymentByMonthInputWithID) (*models.RentPaymentByMonthPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingDelete is the resolver for the rentPaymentByTrainingDelete field.
func (r *mutationResolver) RentPaymentByTrainingDelete(ctx context.Context, id string) (*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingPublishUpdate is the resolver for the rentPaymentByTrainingPublishUpdate field.
func (r *mutationResolver) RentPaymentByTrainingPublishUpdate(ctx context.Context, id string) (*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingSave is the resolver for the rentPaymentByTrainingSave field.
func (r *mutationResolver) RentPaymentByTrainingSave(ctx context.Context, rentPaymentInput models.RentPaymentByTrainingInput) (*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentByTrainingUpdate is the resolver for the rentPaymentByTrainingUpdate field.
func (r *mutationResolver) RentPaymentByTrainingUpdate(ctx context.Context, rentPaymentInput models.RentPaymentByTrainingInputWithID) (*models.RentPaymentByTrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StadiumDelete is the resolver for the stadiumDelete field.
func (r *mutationResolver) StadiumDelete(ctx context.Context, id string) (*models.StadiumPayload, error) {
	_, err := r.Domain.ArticleDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.StadiumPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// StadiumPublishUpdate is the resolver for the stadiumPublishUpdate field.
func (r *mutationResolver) StadiumPublishUpdate(ctx context.Context, id string) (*models.StadiumPayload, error) {
	return r.Domain.StadiumPublish(id)
}

// StadiumSave is the resolver for the stadiumSave field.
func (r *mutationResolver) StadiumSave(ctx context.Context, stadiumInput models.StadiumInput) (*models.StadiumPayload, error) {
	isValid := validation(ctx, stadiumInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.StadiumSave(ctx, stadiumInput)
}

// StadiumUpdate is the resolver for the stadiumUpdate field.
func (r *mutationResolver) StadiumUpdate(ctx context.Context, stadiumInput models.StadiumInputWithID) (*models.StadiumPayload, error) {
	isValid := validation(ctx, stadiumInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.StadiumUpdate(ctx, stadiumInput)
}

// StaffDelete is the resolver for the staffDelete field.
func (r *mutationResolver) StaffDelete(ctx context.Context, id string) (*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffPublishUpdate is the resolver for the staffPublishUpdate field.
func (r *mutationResolver) StaffPublishUpdate(ctx context.Context, id string) (*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffSave is the resolver for the staffSave field.
func (r *mutationResolver) StaffSave(ctx context.Context, staffInput models.StaffInput) (*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StaffUpdate is the resolver for the staffUpdate field.
func (r *mutationResolver) StaffUpdate(ctx context.Context, staffInput models.StaffInputWithID) (*models.StaffPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitDelete is the resolver for the studentVisitDelete field.
func (r *mutationResolver) StudentVisitDelete(ctx context.Context, id string) (*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitPublishUpdate is the resolver for the studentVisitPublishUpdate field.
func (r *mutationResolver) StudentVisitPublishUpdate(ctx context.Context, id string) (*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitSave is the resolver for the studentVisitSave field.
func (r *mutationResolver) StudentVisitSave(ctx context.Context, studentVisitInput models.StudentVisitInput) (*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisitUpdate is the resolver for the studentVisitUpdate field.
func (r *mutationResolver) StudentVisitUpdate(ctx context.Context, studentVisitInput models.StudentVisitInputWithID) (*models.StudentVisitPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentDelete is the resolver for the studentDelete field.
func (r *mutationResolver) StudentDelete(ctx context.Context, id string) (*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentPublishUpdate is the resolver for the studentPublishUpdate field.
func (r *mutationResolver) StudentPublishUpdate(ctx context.Context, id string) (*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentSave is the resolver for the studentSave field.
func (r *mutationResolver) StudentSave(ctx context.Context, studentInput models.StudentInput) (*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentUpdate is the resolver for the studentUpdate field.
func (r *mutationResolver) StudentUpdate(ctx context.Context, studentInput models.StudentInputWithID) (*models.StudentPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TaskDelete is the resolver for the taskDelete field.
func (r *mutationResolver) TaskDelete(ctx context.Context, id string) (*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TaskPublishUpdate is the resolver for the taskPublishUpdate field.
func (r *mutationResolver) TaskPublishUpdate(ctx context.Context, id string) (*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TaskSave is the resolver for the taskSave field.
func (r *mutationResolver) TaskSave(ctx context.Context, taskInput models.TaskInput) (*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TaskUpdate is the resolver for the taskUpdate field.
func (r *mutationResolver) TaskUpdate(ctx context.Context, taskInput models.TaskInputWithID) (*models.TaskPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalanceDelete is the resolver for the teamBalanceDelete field.
func (r *mutationResolver) TeamBalanceDelete(ctx context.Context, id string) (*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalancePublishUpdate is the resolver for the teamBalancePublishUpdate field.
func (r *mutationResolver) TeamBalancePublishUpdate(ctx context.Context, id string) (*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalanceSave is the resolver for the teamBalanceSave field.
func (r *mutationResolver) TeamBalanceSave(ctx context.Context, teamBalanceInput models.TeamBalanceInput) (*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamBalanceUpdate is the resolver for the teamBalanceUpdate field.
func (r *mutationResolver) TeamBalanceUpdate(ctx context.Context, teamBalanceInput models.TeamBalanceInputWithID) (*models.TeamBalancePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamDelete is the resolver for the teamDelete field.
func (r *mutationResolver) TeamDelete(ctx context.Context, id string) (*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamPublishUpdate is the resolver for the teamPublishUpdate field.
func (r *mutationResolver) TeamPublishUpdate(ctx context.Context, id string) (*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamSave is the resolver for the teamSave field.
func (r *mutationResolver) TeamSave(ctx context.Context, teamInput models.TeamInput) (*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TeamUpdate is the resolver for the teamUpdate field.
func (r *mutationResolver) TeamUpdate(ctx context.Context, teamInput models.TeamInputWithID) (*models.TeamPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDayDelete is the resolver for the trainingDayDelete field.
func (r *mutationResolver) TrainingDayDelete(ctx context.Context, id string) (*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDayPublishUpdate is the resolver for the trainingDayPublishUpdate field.
func (r *mutationResolver) TrainingDayPublishUpdate(ctx context.Context, id string) (*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDaySave is the resolver for the trainingDaySave field.
func (r *mutationResolver) TrainingDaySave(ctx context.Context, trainingDayInput models.TrainingDayInput) (*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDayUpdate is the resolver for the trainingDayUpdate field.
func (r *mutationResolver) TrainingDayUpdate(ctx context.Context, trainingDayInput models.TrainingDayInputWithID) (*models.TrainingDayPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingDelete is the resolver for the trainingDelete field.
func (r *mutationResolver) TrainingDelete(ctx context.Context, id string) (*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingPublishUpdate is the resolver for the trainingPublishUpdate field.
func (r *mutationResolver) TrainingPublishUpdate(ctx context.Context, id string) (*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingSave is the resolver for the trainingSave field.
func (r *mutationResolver) TrainingSave(ctx context.Context, trainingInput models.TrainingInput) (*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingUpdate is the resolver for the trainingUpdate field.
func (r *mutationResolver) TrainingUpdate(ctx context.Context, trainingInput models.TrainingInputWithID) (*models.TrainingPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserDelete is the resolver for the userDelete field.
func (r *mutationResolver) UserDelete(ctx context.Context, id string) (*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserPublishUpdate is the resolver for the userPublishUpdate field.
func (r *mutationResolver) UserPublishUpdate(ctx context.Context, id string) (*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserSave is the resolver for the userSave field.
func (r *mutationResolver) UserSave(ctx context.Context, userInput models.UserInput) (*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserUpdate is the resolver for the userUpdate field.
func (r *mutationResolver) UserUpdate(ctx context.Context, userInput models.UserInputWithID) (*models.UserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, id string) (*models.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context, filter *models.ArticleFilter, first *int, after *string, last *int, before *string) (*models.ArticleConnection, error) {
	return r.Domain.ArticlesRepo.GetArticles(filter, first, last, after, before)
}

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context, id string) (*models.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

// Carts is the resolver for the carts field.
func (r *queryResolver) Carts(ctx context.Context, filter *models.CartFilter, first *int, after *string, last *int, before *string) (*models.CartConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalance is the resolver for the clubBalance field.
func (r *queryResolver) ClubBalance(ctx context.Context, id string) (*models.ClubBalance, error) {
	panic(fmt.Errorf("not implemented"))
}

// ClubBalances is the resolver for the clubBalances field.
func (r *queryResolver) ClubBalances(ctx context.Context, filter *models.ClubBalanceFilter, first *int, after *string, last *int, before *string) (*models.ClubBalanceConnection, error) {
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
func (r *queryResolver) CoachPaymentsByMonth(ctx context.Context, filter *models.CoachPaymentByMonthFilter, first *int, after *string, last *int, before *string) (*models.CoachPaymentByMonthConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTeam is the resolver for the coachPaymentsByTeam field.
func (r *queryResolver) CoachPaymentsByTeam(ctx context.Context, filter *models.CoachPaymentByTeamFilter, first *int, after *string, last *int, before *string) (*models.CoachPaymentByTeamConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoachPaymentsByTraining is the resolver for the coachPaymentsByTraining field.
func (r *queryResolver) CoachPaymentsByTraining(ctx context.Context, filter *models.CoachPaymentByTrainingFilter, first *int, after *string, last *int, before *string) (*models.CoachPaymentByTrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Creator is the resolver for the creator field.
func (r *queryResolver) Creator(ctx context.Context, id *string) (*models.Creator, error) {
	panic(fmt.Errorf("not implemented"))
}

// Creators is the resolver for the creators field.
func (r *queryResolver) Creators(ctx context.Context, filter *models.CreatorFilter, first *int, after *string, last *int, before *string) (*models.CreatorConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Kit is the resolver for the kit field.
func (r *queryResolver) Kit(ctx context.Context, id string) (*models.Kit, error) {
	panic(fmt.Errorf("not implemented"))
}

// Kits is the resolver for the kits field.
func (r *queryResolver) Kits(ctx context.Context, filter *models.KitFilter, first *int, after *string, last *int, before *string) (*models.KitConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Lead is the resolver for the lead field.
func (r *queryResolver) Lead(ctx context.Context, id string) (*models.Lead, error) {
	panic(fmt.Errorf("not implemented"))
}

// Leads is the resolver for the leads field.
func (r *queryResolver) Leads(ctx context.Context, filter *models.LeadFilter, first *int, after *string, last *int, before *string) (*models.LeadConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCost is the resolver for the moneyCost field.
func (r *queryResolver) MoneyCost(ctx context.Context, id string) (*models.MoneyCost, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyCosts is the resolver for the moneyCosts field.
func (r *queryResolver) MoneyCosts(ctx context.Context, filter *models.MoneyCostFilter, first *int, after *string, last *int, before *string) (*models.MoneyCostConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMove is the resolver for the moneyMove field.
func (r *queryResolver) MoneyMove(ctx context.Context, id string) (*models.MoneyMove, error) {
	panic(fmt.Errorf("not implemented"))
}

// MoneyMoves is the resolver for the moneyMoves field.
func (r *queryResolver) MoneyMoves(ctx context.Context, filter *models.MoneyMoveFilter, first *int, after *string, last *int, before *string) (*models.MoneyMoveConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, filter *models.OrderFilter, first *int, after *string, last *int, before *string) (*models.OrderConnection, error) {
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
func (r *queryResolver) RentPaymentsByMonth(ctx context.Context, filter *models.RentPaymentByMonthFilter, first *int, after *string, last *int, before *string) (*models.RentPaymentByMonthConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// RentPaymentsByTraining is the resolver for the rentPaymentsByTraining field.
func (r *queryResolver) RentPaymentsByTraining(ctx context.Context, filter *models.RentPaymentByTrainingFilter, first *int, after *string, last *int, before *string) (*models.RentPaymentByTrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Stadium is the resolver for the stadium field.
func (r *queryResolver) Stadium(ctx context.Context, id string) (*models.Stadium, error) {
	panic(fmt.Errorf("not implemented"))
}

// Stadiums is the resolver for the stadiums field.
func (r *queryResolver) Stadiums(ctx context.Context, filter *models.StadiumFilter, first *int, after *string, last *int, before *string) (*models.StadiumConnection, error) {
	return r.Domain.StadiumsRepo.GetStadiums(filter, first, last, after, before)
}

// Staff is the resolver for the staff field.
func (r *queryResolver) Staff(ctx context.Context, filter *models.StaffFilter, first *int, after *string, last *int, before *string) (*models.StaffConnection, error) {
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
func (r *queryResolver) StudentVisits(ctx context.Context, filter *models.StudentVisitFilter, first *int, after *string, last *int, before *string) (*models.StudentVisitConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Students is the resolver for the students field.
func (r *queryResolver) Students(ctx context.Context, filter *models.StudentFilter, first *int, after *string, last *int, before *string) (*models.StudentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*models.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// NearestStudentBirthdays is the resolver for the nearestStudentBirthdays field.
func (r *queryResolver) NearestStudentBirthdays(ctx context.Context, filter *models.StudentFilter, first *int, after *string, last *int, before *string) (*models.StudentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// NearestStaffBirthdays is the resolver for the nearestStaffBirthdays field.
func (r *queryResolver) NearestStaffBirthdays(ctx context.Context, filter *models.StaffFilter, first *int, after *string, last *int, before *string) (*models.StaffConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// UnPayedStudents is the resolver for the unPayedStudents field.
func (r *queryResolver) UnPayedStudents(ctx context.Context, filter *models.StudentFilter, first *int, after *string, last *int, before *string) (*models.StudentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// CurrentTasks is the resolver for the currentTasks field.
func (r *queryResolver) CurrentTasks(ctx context.Context, filter *models.TaskFilter, first *int, after *string, last *int, before *string) (*models.TaskConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// TimeTable is the resolver for the timeTable field.
func (r *queryResolver) TimeTable(ctx context.Context, filter *models.TrainingFilter, first *int, after *string, last *int, before *string) (*models.TrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// TrainingsByDay is the resolver for the trainingsByDay field.
func (r *queryResolver) TrainingsByDay(ctx context.Context, filter *models.TrainingFilter, first *int, after *string, last *int, before *string) (*models.TrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, filter *models.TaskFilter, first *int, after *string, last *int, before *string) (*models.TaskConnection, error) {
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
func (r *queryResolver) TeamBalances(ctx context.Context, filter *models.TeamBalanceFilter, first *int, after *string, last *int, before *string) (*models.TeamBalanceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Teams is the resolver for the teams field.
func (r *queryResolver) Teams(ctx context.Context, filter *models.TeamFilter, first *int, after *string, last *int, before *string) (*models.TeamConnection, error) {
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
func (r *queryResolver) TrainingDays(ctx context.Context, filter *models.TrainingDayFilter, first *int, after *string, last *int, before *string) (*models.TrainingDayConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Trainings is the resolver for the trainings field.
func (r *queryResolver) Trainings(ctx context.Context, filter *models.TrainingFilter, first *int, after *string, last *int, before *string) (*models.TrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.Domain.UsersRepo.GetUserByID(id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, filter *models.UserFilter, first *int, after *string, last *int, before *string) (*models.UserConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Place is the resolver for the place field.
func (r *stadiumResolver) Place(ctx context.Context, obj *models.Stadium) (*models.Place, error) {
	return For(ctx).PlaceLoader.Load(*obj.PlaceID)
}

// Article returns generated.ArticleResolver implementation.
func (r *Resolver) Article() generated.ArticleResolver { return &articleResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Stadium returns generated.StadiumResolver implementation.
func (r *Resolver) Stadium() generated.StadiumResolver { return &stadiumResolver{r} }

type articleResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type stadiumResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	ErrInput = errors.New("inputs errors")
)
