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

// Kits is the resolver for the kits field.
func (r *cartResolver) Kits(ctx context.Context, obj *models.Cart) ([]*models.Kit, error) {
	return r.Domain.CartsRepo.GetKitsFoCart(obj)
}

// Student is the resolver for the student field.
func (r *cartResolver) Student(ctx context.Context, obj *models.Cart) (*models.Student, error) {
	return For(ctx).StudentLoader.Load(obj.StudentID)
}

// Coach is the resolver for the coach field.
func (r *coachPaymentByMonthResolver) Coach(ctx context.Context, obj *models.CoachPaymentByMonth) (*models.Staff, error) {
	return For(ctx).StaffLoader.Load(obj.CoachID)
}

// Coach is the resolver for the coach field.
func (r *coachPaymentByTeamResolver) Coach(ctx context.Context, obj *models.CoachPaymentByTeam) (*models.Staff, error) {
	return For(ctx).StaffLoader.Load(obj.CoachID)
}

// Team is the resolver for the team field.
func (r *coachPaymentByTeamResolver) Team(ctx context.Context, obj *models.CoachPaymentByTeam) (*models.Team, error) {
	return For(ctx).TeamLoader.Load(*obj.TeamID)
}

// Coach is the resolver for the coach field.
func (r *coachPaymentByTrainingResolver) Coach(ctx context.Context, obj *models.CoachPaymentByTraining) (*models.Staff, error) {
	return For(ctx).StaffLoader.Load(obj.CoachID)
}

// Training is the resolver for the training field.
func (r *coachPaymentByTrainingResolver) Training(ctx context.Context, obj *models.CoachPaymentByTraining) (*models.Training, error) {
	return For(ctx).TrainingLoader.Load(*obj.TrainingID)
}

// User is the resolver for the user field.
func (r *creatorResolver) User(ctx context.Context, obj *models.Creator) (*models.User, error) {
	return For(ctx).UserLoader.Load(*obj.UserID)
}

// NextVisit is the resolver for the nextVisit field.
func (r *leadResolver) NextVisit(ctx context.Context, obj *models.Lead) (*models.Training, error) {
	return For(ctx).TrainingLoader.Load(*obj.NextVisitID)
}

// Students is the resolver for the students field.
func (r *leadResolver) Students(ctx context.Context, obj *models.Lead) ([]*models.Student, error) {
	return r.Domain.LeadsRepo.GetStudentsForLead(obj)
}

// Team is the resolver for the team field.
func (r *leadResolver) Team(ctx context.Context, obj *models.Lead) (*models.Team, error) {
	return For(ctx).TeamLoader.Load(*obj.TeamID)
}

// Staff is the resolver for the staff field.
func (r *moneyCostResolver) Staff(ctx context.Context, obj *models.MoneyCost) (*models.Staff, error) {
	return For(ctx).StaffLoader.Load(obj.StaffID)
}

// Owner is the resolver for the owner field.
func (r *moneyMoveResolver) Owner(ctx context.Context, obj *models.MoneyMove) (*models.Staff, error) {
	return For(ctx).StaffLoader.Load(obj.OwnerID)
}

// Student is the resolver for the student field.
func (r *moneyMoveResolver) Student(ctx context.Context, obj *models.MoneyMove) (*models.Student, error) {
	return For(ctx).StudentLoader.Load(obj.StudentID)
}

// User is the resolver for the user field.
func (r *moneyMoveResolver) User(ctx context.Context, obj *models.MoneyMove) (*models.User, error) {
	return For(ctx).UserLoader.Load(obj.UserID)
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
	_, err := r.Domain.CartDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.CartPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// CartPublishUpdate is the resolver for the cartPublishUpdate field.
func (r *mutationResolver) CartPublishUpdate(ctx context.Context, id string) (*models.CartPayload, error) {
	return r.Domain.CartPublish(id)
}

// CartSave is the resolver for the cartSave field.
func (r *mutationResolver) CartSave(ctx context.Context, cartInput models.CartInput) (*models.CartPayload, error) {
	isValid := validation(ctx, cartInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CartSave(ctx, cartInput)
}

// CartUpdate is the resolver for the cartUpdate field.
func (r *mutationResolver) CartUpdate(ctx context.Context, cartInput models.CartInputWithID) (*models.CartPayload, error) {
	isValid := validation(ctx, cartInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CartUpdate(ctx, cartInput)
}

// ClubBalanceDelete is the resolver for the clubBalanceDelete field.
func (r *mutationResolver) ClubBalanceDelete(ctx context.Context, id string) (*models.ClubBalancePayload, error) {
	_, err := r.Domain.ClubBalanceDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.ClubBalancePayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// ClubBalancePublishUpdate is the resolver for the clubBalancePublishUpdate field.
func (r *mutationResolver) ClubBalancePublishUpdate(ctx context.Context, id string) (*models.ClubBalancePayload, error) {
	return r.Domain.ClubBalancePublish(id)
}

// ClubBalanceSave is the resolver for the clubBalanceSave field.
func (r *mutationResolver) ClubBalanceSave(ctx context.Context, clubBalanceInput models.ClubBalanceInput) (*models.ClubBalancePayload, error) {
	isValid := validation(ctx, clubBalanceInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.ClubBalanceSave(ctx, clubBalanceInput)
}

// ClubBalanceUpdate is the resolver for the clubBalanceUpdate field.
func (r *mutationResolver) ClubBalanceUpdate(ctx context.Context, clubBalanceInput models.ClubBalanceInputWithID) (*models.ClubBalancePayload, error) {
	isValid := validation(ctx, clubBalanceInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.ClubBalanceUpdate(ctx, clubBalanceInput)
}

// CoachPaymentByMonthDelete is the resolver for the coachPaymentByMonthDelete field.
func (r *mutationResolver) CoachPaymentByMonthDelete(ctx context.Context, id string) (*models.CoachPaymentByMonthPayload, error) {
	_, err := r.Domain.CoachPaymentByMonthDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.CoachPaymentByMonthPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// CoachPaymentByMonthPublishUpdate is the resolver for the coachPaymentByMonthPublishUpdate field.
func (r *mutationResolver) CoachPaymentByMonthPublishUpdate(ctx context.Context, id string) (*models.CoachPaymentByMonthPayload, error) {
	return r.Domain.CoachPaymentByMonthPublish(id)
}

// CoachPaymentByMonthSave is the resolver for the coachPaymentByMonthSave field.
func (r *mutationResolver) CoachPaymentByMonthSave(ctx context.Context, coachPaymentByMonthInput models.CoachPaymentByMonthInput) (*models.CoachPaymentByMonthPayload, error) {
	isValid := validation(ctx, coachPaymentByMonthInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CoachPaymentByMonthSave(ctx, coachPaymentByMonthInput)
}

// CoachPaymentByMonthUpdate is the resolver for the coachPaymentByMonthUpdate field.
func (r *mutationResolver) CoachPaymentByMonthUpdate(ctx context.Context, coachPaymentByMonthInput models.CoachPaymentByMonthInputWithID) (*models.CoachPaymentByMonthPayload, error) {
	isValid := validation(ctx, coachPaymentByMonthInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CoachPaymentByMonthUpdate(ctx, coachPaymentByMonthInput)
}

// CoachPaymentByTeamDelete is the resolver for the coachPaymentByTeamDelete field.
func (r *mutationResolver) CoachPaymentByTeamDelete(ctx context.Context, id string) (*models.CoachPaymentByTeamPayload, error) {
	_, err := r.Domain.CoachPaymentByTeamDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.CoachPaymentByTeamPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// CoachPaymentByTeamPublishUpdate is the resolver for the coachPaymentByTeamPublishUpdate field.
func (r *mutationResolver) CoachPaymentByTeamPublishUpdate(ctx context.Context, id string) (*models.CoachPaymentByTeamPayload, error) {
	return r.Domain.CoachPaymentByTeamPublish(id)
}

// CoachPaymentByTeamSave is the resolver for the coachPaymentByTeamSave field.
func (r *mutationResolver) CoachPaymentByTeamSave(ctx context.Context, coachPaymentByTeamInput models.CoachPaymentByTeamInput) (*models.CoachPaymentByTeamPayload, error) {
	isValid := validation(ctx, coachPaymentByTeamInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CoachPaymentByTeamSave(ctx, coachPaymentByTeamInput)
}

// CoachPaymentByTeamUpdate is the resolver for the coachPaymentByTeamUpdate field.
func (r *mutationResolver) CoachPaymentByTeamUpdate(ctx context.Context, coachPaymentByTeamInput models.CoachPaymentByTeamInputWithID) (*models.CoachPaymentByTeamPayload, error) {
	isValid := validation(ctx, coachPaymentByTeamInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CoachPaymentByTeamUpdate(ctx, coachPaymentByTeamInput)
}

// CoachPaymentByTrainingDelete is the resolver for the coachPaymentByTrainingDelete field.
func (r *mutationResolver) CoachPaymentByTrainingDelete(ctx context.Context, id string) (*models.CoachPaymentByTrainingPayload, error) {
	_, err := r.Domain.CoachPaymentByTrainingDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.CoachPaymentByTrainingPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// CoachPaymentByTrainingPublishUpdate is the resolver for the coachPaymentByTrainingPublishUpdate field.
func (r *mutationResolver) CoachPaymentByTrainingPublishUpdate(ctx context.Context, id string) (*models.CoachPaymentByTrainingPayload, error) {
	return r.Domain.CoachPaymentByTrainingPublish(id)
}

// CoachPaymentByTrainingSave is the resolver for the coachPaymentByTrainingSave field.
func (r *mutationResolver) CoachPaymentByTrainingSave(ctx context.Context, coachPaymentByTrainingInput models.CoachPaymentByTrainingInput) (*models.CoachPaymentByTrainingPayload, error) {
	isValid := validation(ctx, coachPaymentByTrainingInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CoachPaymentByTrainingSave(ctx, coachPaymentByTrainingInput)
}

// CoachPaymentByTrainingUpdate is the resolver for the coachPaymentByTrainingUpdate field.
func (r *mutationResolver) CoachPaymentByTrainingUpdate(ctx context.Context, coachPaymentByTrainingInput models.CoachPaymentByTrainingInputWithID) (*models.CoachPaymentByTrainingPayload, error) {
	isValid := validation(ctx, coachPaymentByTrainingInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CoachPaymentByTrainingUpdate(ctx, coachPaymentByTrainingInput)
}

// CreatorDelete is the resolver for the creatorDelete field.
func (r *mutationResolver) CreatorDelete(ctx context.Context, id string) (*models.CreatorPayload, error) {
	_, err := r.Domain.CreatorDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.CreatorPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// CreatorPublishUpdate is the resolver for the creatorPublishUpdate field.
func (r *mutationResolver) CreatorPublishUpdate(ctx context.Context, id string) (*models.CreatorPayload, error) {
	return r.Domain.CreatorPublish(id)
}

// CreatorSave is the resolver for the creatorSave field.
func (r *mutationResolver) CreatorSave(ctx context.Context, creatorInput models.CreatorInput) (*models.CreatorPayload, error) {
	isValid := validation(ctx, creatorInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CreatorSave(ctx, creatorInput)
}

// CreatorUpdate is the resolver for the creatorUpdate field.
func (r *mutationResolver) CreatorUpdate(ctx context.Context, creatorInput models.CreatorInputWithID) (*models.CreatorPayload, error) {
	isValid := validation(ctx, creatorInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.CreatorUpdate(ctx, creatorInput)
}

// KitDelete is the resolver for the kitDelete field.
func (r *mutationResolver) KitDelete(ctx context.Context, id string) (*models.KitPayload, error) {
	_, err := r.Domain.KitDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.KitPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// KitPublishUpdate is the resolver for the kitPublishUpdate field.
func (r *mutationResolver) KitPublishUpdate(ctx context.Context, id string) (*models.KitPayload, error) {
	return r.Domain.KitPublish(id)
}

// KitSave is the resolver for the kitSave field.
func (r *mutationResolver) KitSave(ctx context.Context, kitInput models.KitInput) (*models.KitPayload, error) {
	isValid := validation(ctx, kitInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.KitSave(ctx, kitInput)
}

// KitUpdate is the resolver for the kitUpdate field.
func (r *mutationResolver) KitUpdate(ctx context.Context, kitInput models.KitInputWithID) (*models.KitPayload, error) {
	isValid := validation(ctx, kitInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.KitUpdate(ctx, kitInput)
}

// LeadDelete is the resolver for the leadDelete field.
func (r *mutationResolver) LeadDelete(ctx context.Context, id string) (*models.LeadPayload, error) {
	_, err := r.Domain.LeadDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.LeadPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// LeadPublishUpdate is the resolver for the leadPublishUpdate field.
func (r *mutationResolver) LeadPublishUpdate(ctx context.Context, id string) (*models.LeadPayload, error) {
	return r.Domain.LeadPublish(id)
}

// LeadSave is the resolver for the leadSave field.
func (r *mutationResolver) LeadSave(ctx context.Context, leadInput models.LeadInput) (*models.LeadPayload, error) {
	isValid := validation(ctx, leadInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.LeadSave(ctx, leadInput)
}

// LeadUpdate is the resolver for the leadUpdate field.
func (r *mutationResolver) LeadUpdate(ctx context.Context, leadInput models.LeadInputWithID) (*models.LeadPayload, error) {
	isValid := validation(ctx, leadInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.LeadUpdate(ctx, leadInput)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}

	return r.Domain.Login(ctx, input)
}

// MoneyCostDelete is the resolver for the moneyCostDelete field.
func (r *mutationResolver) MoneyCostDelete(ctx context.Context, id string) (*models.MoneyCostPayload, error) {
	_, err := r.Domain.MoneyCostDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.MoneyCostPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// MoneyCostPublishUpdate is the resolver for the moneyCostPublishUpdate field.
func (r *mutationResolver) MoneyCostPublishUpdate(ctx context.Context, id string) (*models.MoneyCostPayload, error) {
	return r.Domain.MoneyCostPublish(id)
}

// MoneyCostSave is the resolver for the moneyCostSave field.
func (r *mutationResolver) MoneyCostSave(ctx context.Context, moneyCostInput models.MoneyCostInput) (*models.MoneyCostPayload, error) {
	isValid := validation(ctx, moneyCostInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.MoneyCostSave(ctx, moneyCostInput)
}

// MoneyCostUpdate is the resolver for the moneyCostUpdate field.
func (r *mutationResolver) MoneyCostUpdate(ctx context.Context, moneyCostInput models.MoneyCostInputWithID) (*models.MoneyCostPayload, error) {
	isValid := validation(ctx, moneyCostInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.MoneyCostUpdate(ctx, moneyCostInput)
}

// MoneyMoveDelete is the resolver for the moneyMoveDelete field.
func (r *mutationResolver) MoneyMoveDelete(ctx context.Context, id string) (*models.MoneyMovePayload, error) {
	_, err := r.Domain.MoneyMoveDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.MoneyMovePayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// MoneyMovePublishUpdate is the resolver for the moneyMovePublishUpdate field.
func (r *mutationResolver) MoneyMovePublishUpdate(ctx context.Context, id string) (*models.MoneyMovePayload, error) {
	return r.Domain.MoneyMovePublish(id)
}

// MoneyMoveSave is the resolver for the moneyMoveSave field.
func (r *mutationResolver) MoneyMoveSave(ctx context.Context, moneyMoveInput models.MoneyMoveInput) (*models.MoneyMovePayload, error) {
	isValid := validation(ctx, moneyMoveInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.MoneyMoveSave(ctx, moneyMoveInput)
}

// MoneyMoveUpdate is the resolver for the moneyMoveUpdate field.
func (r *mutationResolver) MoneyMoveUpdate(ctx context.Context, moneyMoveInput models.MoneyMoveInputWithID) (*models.MoneyMovePayload, error) {
	isValid := validation(ctx, moneyMoveInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.MoneyMoveUpdate(ctx, moneyMoveInput)
}

// OrderDelete is the resolver for the orderDelete field.
func (r *mutationResolver) OrderDelete(ctx context.Context, id string) (*models.OrderPayload, error) {
	_, err := r.Domain.OrderDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.OrderPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// OrderPublishUpdate is the resolver for the orderPublishUpdate field.
func (r *mutationResolver) OrderPublishUpdate(ctx context.Context, id string) (*models.OrderPayload, error) {
	return r.Domain.OrderPublish(id)
}

// OrderSave is the resolver for the orderSave field.
func (r *mutationResolver) OrderSave(ctx context.Context, orderInput models.OrderInput) (*models.OrderPayload, error) {
	isValid := validation(ctx, orderInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.OrderSave(ctx, orderInput)
}

// OrderUpdate is the resolver for the orderUpdate field.
func (r *mutationResolver) OrderUpdate(ctx context.Context, orderInput models.OrderInputWithID) (*models.OrderPayload, error) {
	isValid := validation(ctx, orderInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.OrderUpdate(ctx, orderInput)
}

// PlaceDelete is the resolver for the placeDelete field.
func (r *mutationResolver) PlaceDelete(ctx context.Context, id string) (*models.PlacePayload, error) {
	_, err := r.Domain.PlaceDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.PlacePayload{
		RecordID: id,
		Record:   nil,
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

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}

	return r.Domain.Register(ctx, input)
}

// RentPaymentByMonthDelete is the resolver for the rentPaymentByMonthDelete field.
func (r *mutationResolver) RentPaymentByMonthDelete(ctx context.Context, id string) (*models.RentPaymentByMonthPayload, error) {
	_, err := r.Domain.RentPaymentByMonthDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.RentPaymentByMonthPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// RentPaymentByMonthPublishUpdate is the resolver for the rentPaymentByMonthPublishUpdate field.
func (r *mutationResolver) RentPaymentByMonthPublishUpdate(ctx context.Context, id string) (*models.RentPaymentByMonthPayload, error) {
	return r.Domain.RentPaymentByMonthPublish(id)
}

// RentPaymentByMonthSave is the resolver for the rentPaymentByMonthSave field.
func (r *mutationResolver) RentPaymentByMonthSave(ctx context.Context, rentPaymentInput models.RentPaymentByMonthInput) (*models.RentPaymentByMonthPayload, error) {
	isValid := validation(ctx, rentPaymentInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.RentPaymentByMonthSave(ctx, rentPaymentInput)
}

// RentPaymentByMonthUpdate is the resolver for the rentPaymentByMonthUpdate field.
func (r *mutationResolver) RentPaymentByMonthUpdate(ctx context.Context, rentPaymentInput models.RentPaymentByMonthInputWithID) (*models.RentPaymentByMonthPayload, error) {
	isValid := validation(ctx, rentPaymentInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.RentPaymentByMonthUpdate(ctx, rentPaymentInput)
}

// RentPaymentByTrainingDelete is the resolver for the rentPaymentByTrainingDelete field.
func (r *mutationResolver) RentPaymentByTrainingDelete(ctx context.Context, id string) (*models.RentPaymentByTrainingPayload, error) {
	_, err := r.Domain.RentPaymentByTrainingDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.RentPaymentByTrainingPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// RentPaymentByTrainingPublishUpdate is the resolver for the rentPaymentByTrainingPublishUpdate field.
func (r *mutationResolver) RentPaymentByTrainingPublishUpdate(ctx context.Context, id string) (*models.RentPaymentByTrainingPayload, error) {
	return r.Domain.RentPaymentByTrainingPublish(id)
}

// RentPaymentByTrainingSave is the resolver for the rentPaymentByTrainingSave field.
func (r *mutationResolver) RentPaymentByTrainingSave(ctx context.Context, rentPaymentInput models.RentPaymentByTrainingInput) (*models.RentPaymentByTrainingPayload, error) {
	isValid := validation(ctx, rentPaymentInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.RentPaymentByTrainingSave(ctx, rentPaymentInput)
}

// RentPaymentByTrainingUpdate is the resolver for the rentPaymentByTrainingUpdate field.
func (r *mutationResolver) RentPaymentByTrainingUpdate(ctx context.Context, rentPaymentInput models.RentPaymentByTrainingInputWithID) (*models.RentPaymentByTrainingPayload, error) {
	isValid := validation(ctx, rentPaymentInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.RentPaymentByTrainingUpdate(ctx, rentPaymentInput)
}

// StadiumDelete is the resolver for the stadiumDelete field.
func (r *mutationResolver) StadiumDelete(ctx context.Context, id string) (*models.StadiumPayload, error) {
	_, err := r.Domain.StadiumDelete(id)
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
	_, err := r.Domain.StaffDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.StaffPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// StaffPublishUpdate is the resolver for the staffPublishUpdate field.
func (r *mutationResolver) StaffPublishUpdate(ctx context.Context, id string) (*models.StaffPayload, error) {
	return r.Domain.StaffPublish(id)
}

// StaffSave is the resolver for the staffSave field.
func (r *mutationResolver) StaffSave(ctx context.Context, staffInput models.StaffInput) (*models.StaffPayload, error) {
	isValid := validation(ctx, staffInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.StaffSave(ctx, staffInput)
}

// StaffUpdate is the resolver for the staffUpdate field.
func (r *mutationResolver) StaffUpdate(ctx context.Context, staffInput models.StaffInputWithID) (*models.StaffPayload, error) {
	isValid := validation(ctx, staffInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.StaffUpdate(ctx, staffInput)
}

// StudentDelete is the resolver for the studentDelete field.
func (r *mutationResolver) StudentDelete(ctx context.Context, id string) (*models.StudentPayload, error) {
	_, err := r.Domain.StudentDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.StudentPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// StudentPublishUpdate is the resolver for the studentPublishUpdate field.
func (r *mutationResolver) StudentPublishUpdate(ctx context.Context, id string) (*models.StudentPayload, error) {
	return r.Domain.StudentPublish(id)
}

// StudentSave is the resolver for the studentSave field.
func (r *mutationResolver) StudentSave(ctx context.Context, studentInput models.StudentInput) (*models.StudentPayload, error) {
	isValid := validation(ctx, studentInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.StudentSave(ctx, studentInput)
}

// StudentUpdate is the resolver for the studentUpdate field.
func (r *mutationResolver) StudentUpdate(ctx context.Context, studentInput models.StudentInputWithID) (*models.StudentPayload, error) {
	isValid := validation(ctx, studentInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.StudentUpdate(ctx, studentInput)
}

// StudentVisitDelete is the resolver for the studentVisitDelete field.
func (r *mutationResolver) StudentVisitDelete(ctx context.Context, id string) (*models.StudentVisitPayload, error) {
	_, err := r.Domain.StudentVisitDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.StudentVisitPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// StudentVisitPublishUpdate is the resolver for the studentVisitPublishUpdate field.
func (r *mutationResolver) StudentVisitPublishUpdate(ctx context.Context, id string) (*models.StudentVisitPayload, error) {
	return r.Domain.StudentVisitPublish(id)
}

// StudentVisitSave is the resolver for the studentVisitSave field.
func (r *mutationResolver) StudentVisitSave(ctx context.Context, studentVisitInput models.StudentVisitInput) (*models.StudentVisitPayload, error) {
	isValid := validation(ctx, studentVisitInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.StudentVisitSave(ctx, studentVisitInput)
}

// StudentVisitUpdate is the resolver for the studentVisitUpdate field.
func (r *mutationResolver) StudentVisitUpdate(ctx context.Context, studentVisitInput models.StudentVisitInputWithID) (*models.StudentVisitPayload, error) {
	isValid := validation(ctx, studentVisitInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.StudentVisitUpdate(ctx, studentVisitInput)
}

// TaskDelete is the resolver for the taskDelete field.
func (r *mutationResolver) TaskDelete(ctx context.Context, id string) (*models.TaskPayload, error) {
	_, err := r.Domain.TaskDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.TaskPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// TaskPublishUpdate is the resolver for the taskPublishUpdate field.
func (r *mutationResolver) TaskPublishUpdate(ctx context.Context, id string) (*models.TaskPayload, error) {
	return r.Domain.TaskPublish(id)
}

// TaskSave is the resolver for the taskSave field.
func (r *mutationResolver) TaskSave(ctx context.Context, taskInput models.TaskInput) (*models.TaskPayload, error) {
	isValid := validation(ctx, taskInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TaskSave(ctx, taskInput)
}

// TaskUpdate is the resolver for the taskUpdate field.
func (r *mutationResolver) TaskUpdate(ctx context.Context, taskInput models.TaskInputWithID) (*models.TaskPayload, error) {
	isValid := validation(ctx, taskInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TaskUpdate(ctx, taskInput)
}

// TeamBalanceDelete is the resolver for the teamBalanceDelete field.
func (r *mutationResolver) TeamBalanceDelete(ctx context.Context, id string) (*models.TeamBalancePayload, error) {
	_, err := r.Domain.TeamBalanceDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.TeamBalancePayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// TeamBalancePublishUpdate is the resolver for the teamBalancePublishUpdate field.
func (r *mutationResolver) TeamBalancePublishUpdate(ctx context.Context, id string) (*models.TeamBalancePayload, error) {
	return r.Domain.TeamBalancePublish(id)
}

// TeamBalanceSave is the resolver for the teamBalanceSave field.
func (r *mutationResolver) TeamBalanceSave(ctx context.Context, teamBalanceInput models.TeamBalanceInput) (*models.TeamBalancePayload, error) {
	isValid := validation(ctx, teamBalanceInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TeamBalanceSave(ctx, teamBalanceInput)
}

// TeamBalanceUpdate is the resolver for the teamBalanceUpdate field.
func (r *mutationResolver) TeamBalanceUpdate(ctx context.Context, teamBalanceInput models.TeamBalanceInputWithID) (*models.TeamBalancePayload, error) {
	isValid := validation(ctx, teamBalanceInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TeamBalanceUpdate(ctx, teamBalanceInput)
}

// TeamDelete is the resolver for the teamDelete field.
func (r *mutationResolver) TeamDelete(ctx context.Context, id string) (*models.TeamPayload, error) {
	_, err := r.Domain.TeamDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.TeamPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// TeamPublishUpdate is the resolver for the teamPublishUpdate field.
func (r *mutationResolver) TeamPublishUpdate(ctx context.Context, id string) (*models.TeamPayload, error) {
	return r.Domain.TeamPublish(id)
}

// TeamSave is the resolver for the teamSave field.
func (r *mutationResolver) TeamSave(ctx context.Context, teamInput models.TeamInput) (*models.TeamPayload, error) {
	isValid := validation(ctx, teamInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TeamSave(ctx, teamInput)
}

// TeamUpdate is the resolver for the teamUpdate field.
func (r *mutationResolver) TeamUpdate(ctx context.Context, teamInput models.TeamInputWithID) (*models.TeamPayload, error) {
	isValid := validation(ctx, teamInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TeamUpdate(ctx, teamInput)
}

// TrainingDayDelete is the resolver for the trainingDayDelete field.
func (r *mutationResolver) TrainingDayDelete(ctx context.Context, id string) (*models.TrainingDayPayload, error) {
	_, err := r.Domain.TrainingDayDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.TrainingDayPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// TrainingDayPublishUpdate is the resolver for the trainingDayPublishUpdate field.
func (r *mutationResolver) TrainingDayPublishUpdate(ctx context.Context, id string) (*models.TrainingDayPayload, error) {
	return r.Domain.TrainingDayPublish(id)
}

// TrainingDaySave is the resolver for the trainingDaySave field.
func (r *mutationResolver) TrainingDaySave(ctx context.Context, trainingDayInput models.TrainingDayInput) (*models.TrainingDayPayload, error) {
	isValid := validation(ctx, trainingDayInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TrainingDaySave(ctx, trainingDayInput)
}

// TrainingDayUpdate is the resolver for the trainingDayUpdate field.
func (r *mutationResolver) TrainingDayUpdate(ctx context.Context, trainingDayInput models.TrainingDayInputWithID) (*models.TrainingDayPayload, error) {
	isValid := validation(ctx, trainingDayInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TrainingDayUpdate(ctx, trainingDayInput)
}

// TrainingDelete is the resolver for the trainingDelete field.
func (r *mutationResolver) TrainingDelete(ctx context.Context, id string) (*models.TrainingPayload, error) {
	_, err := r.Domain.TrainingDelete(id)
	if err != nil {
		return nil, err
	}
	return &models.TrainingPayload{
		RecordID: id,
		Record:   nil,
	}, nil
}

// TrainingPublishUpdate is the resolver for the trainingPublishUpdate field.
func (r *mutationResolver) TrainingPublishUpdate(ctx context.Context, id string) (*models.TrainingPayload, error) {
	return r.Domain.TrainingPublish(id)
}

// TrainingSave is the resolver for the trainingSave field.
func (r *mutationResolver) TrainingSave(ctx context.Context, trainingInput models.TrainingInput) (*models.TrainingPayload, error) {
	isValid := validation(ctx, trainingInput)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TrainingSave(ctx, trainingInput)
}

// TrainingUpdate is the resolver for the trainingUpdate field.
func (r *mutationResolver) TrainingUpdate(ctx context.Context, trainingInput models.TrainingInputWithID) (*models.TrainingPayload, error) {
	isValid := validation(ctx, trainingInput.Input)
	if !isValid {
		return nil, ErrInput
	}
	return r.Domain.TrainingUpdate(ctx, trainingInput)
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

// Cart is the resolver for the cart field.
func (r *orderResolver) Cart(ctx context.Context, obj *models.Order) (*models.Cart, error) {
	return For(ctx).CartLoader.Load(obj.CartID)
}

// Creator is the resolver for the creator field.
func (r *orderResolver) Creator(ctx context.Context, obj *models.Order) (*models.Creator, error) {
	return For(ctx).CreatorLoader.Load(obj.CreatorID)
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, id string) (*models.Article, error) {
	return r.Domain.ArticlesRepo.GetArticleByID(id)

}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context, after *string, before *string, filter *models.ArticleFilter, first *int, last *int) (*models.ArticleConnection, error) {
	return r.Domain.ArticlesRepo.GetArticles(filter, first, last, after, before)
}

// Cart is the resolver for the cart field.
func (r *queryResolver) Cart(ctx context.Context, id string) (*models.Cart, error) {
	return r.Domain.CartsRepo.GetCartsByID(id)
}

// CartAll is the resolver for the cartAll field.
func (r *queryResolver) CartAll(ctx context.Context) ([]*models.CartDto, error) {
	panic(fmt.Errorf("not implemented"))
}

// Carts is the resolver for the carts field.
func (r *queryResolver) Carts(ctx context.Context, after *string, before *string, filter *models.CartFilter, first *int, last *int) (*models.CartConnection, error) {
	return r.Domain.CartsRepo.GetCarts(filter, first, last, after, before)
}

// ClubBalance is the resolver for the clubBalance field.
func (r *queryResolver) ClubBalance(ctx context.Context, id string) (*models.ClubBalance, error) {
	return r.Domain.ClubBalancesRepo.GetClubBalanceByID(id)
}

// ClubBalances is the resolver for the clubBalances field.
func (r *queryResolver) ClubBalances(ctx context.Context, after *string, before *string, filter *models.ClubBalanceFilter, first *int, last *int) (*models.ClubBalanceConnection, error) {
	return r.Domain.ClubBalancesRepo.GetClubBalances(filter, first, last, after, before)
}

// CoachPaymentByMonth is the resolver for the coachPaymentByMonth field.
func (r *queryResolver) CoachPaymentByMonth(ctx context.Context, id string) (*models.CoachPaymentByMonth, error) {
	return r.Domain.CoachPaymentByMonthRepo.GetCoachPaymentByMonthByID(id)
}

// CoachPaymentByTeam is the resolver for the coachPaymentByTeam field.
func (r *queryResolver) CoachPaymentByTeam(ctx context.Context, id string) (*models.CoachPaymentByTeam, error) {
	return r.Domain.CoachPaymentByTeamRepo.GetCoachPaymentByTeamByID(id)
}

// CoachPaymentByTraining is the resolver for the coachPaymentByTraining field.
func (r *queryResolver) CoachPaymentByTraining(ctx context.Context, id string) (*models.CoachPaymentByTraining, error) {
	return r.Domain.CoachPaymentByTrainingRepo.GetCoachPaymentByTrainingByID(id)
}

// CoachPaymentsByMonth is the resolver for the coachPaymentsByMonth field.
func (r *queryResolver) CoachPaymentsByMonth(ctx context.Context, after *string, before *string, filter *models.CoachPaymentByMonthFilter, first *int, last *int) (*models.CoachPaymentByMonthConnection, error) {
	return r.Domain.CoachPaymentByMonthRepo.GetCoachPaymentByMonth(filter, first, last, after, before)
}

// CoachPaymentsByTeam is the resolver for the coachPaymentsByTeam field.
func (r *queryResolver) CoachPaymentsByTeam(ctx context.Context, after *string, before *string, filter *models.CoachPaymentByTeamFilter, first *int, last *int) (*models.CoachPaymentByTeamConnection, error) {
	return r.Domain.CoachPaymentByTeamRepo.GetCoachPaymentByTeam(filter, first, last, after, before)
}

// CoachPaymentsByTraining is the resolver for the coachPaymentsByTraining field.
func (r *queryResolver) CoachPaymentsByTraining(ctx context.Context, after *string, before *string, filter *models.CoachPaymentByTrainingFilter, first *int, last *int) (*models.CoachPaymentByTrainingConnection, error) {
	return r.Domain.CoachPaymentByTrainingRepo.GetCoachPaymentByTraining(filter, first, last, after, before)
}

// Creator is the resolver for the creator field.
func (r *queryResolver) Creator(ctx context.Context, id *string) (*models.Creator, error) {
	return r.Domain.CreatorsRepo.GetCreatorByID(*id)
}

// CreatorAll is the resolver for the creatorAll field.
func (r *queryResolver) CreatorAll(ctx context.Context) ([]*models.CreatorDto, error) {
	panic(fmt.Errorf("not implemented"))
}

// Creators is the resolver for the creators field.
func (r *queryResolver) Creators(ctx context.Context, after *string, before *string, filter *models.CreatorFilter, first *int, last *int) (*models.CreatorConnection, error) {
	return r.Domain.CreatorsRepo.GetCreators(filter, first, last, after, before)
}

// CurrentTasks is the resolver for the currentTasks field.
func (r *queryResolver) CurrentTasks(ctx context.Context, after *string, before *string, filter *models.TaskFilter, first *int, last *int) (*models.TaskConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Kit is the resolver for the kit field.
func (r *queryResolver) Kit(ctx context.Context, id string) (*models.Kit, error) {
	return r.Domain.KitsRepo.GetKitsByID(id)
}

// KitAll is the resolver for the kitAll field.
func (r *queryResolver) KitAll(ctx context.Context) ([]*models.KitDto, error) {
	panic(fmt.Errorf("not implemented"))
}

// Kits is the resolver for the kits field.
func (r *queryResolver) Kits(ctx context.Context, after *string, before *string, filter *models.KitFilter, first *int, last *int) (*models.KitConnection, error) {
	return r.Domain.KitsRepo.GetKits(filter, first, last, after, before)
}

// Lead is the resolver for the lead field.
func (r *queryResolver) Lead(ctx context.Context, id string) (*models.Lead, error) {
	return r.Domain.LeadsRepo.GetLeadsByID(id)
}

// LeadAll is the resolver for the leadAll field.
func (r *queryResolver) LeadAll(ctx context.Context) ([]*models.LeadDto, error) {
	panic(fmt.Errorf("not implemented"))
}

// Leads is the resolver for the leads field.
func (r *queryResolver) Leads(ctx context.Context, after *string, before *string, filter *models.LeadFilter, first *int, last *int) (*models.LeadConnection, error) {
	return r.Domain.LeadsRepo.GetLeads(filter, first, last, after, before)
}

// MoneyCost is the resolver for the moneyCost field.
func (r *queryResolver) MoneyCost(ctx context.Context, id string) (*models.MoneyCost, error) {
	return r.Domain.MoneyCostRepo.GetMoneyCostByID(id)
}

// MoneyCosts is the resolver for the moneyCosts field.
func (r *queryResolver) MoneyCosts(ctx context.Context, after *string, before *string, filter *models.MoneyCostFilter, first *int, last *int) (*models.MoneyCostConnection, error) {
	return r.Domain.MoneyCostRepo.GetMoneyCosts(filter, first, last, after, before)
}

// MoneyMove is the resolver for the moneyMove field.
func (r *queryResolver) MoneyMove(ctx context.Context, id string) (*models.MoneyMove, error) {
	return r.Domain.MoneyMoveRepo.GetMoneyMoveByID(id)
}

// MoneyMoves is the resolver for the moneyMoves field.
func (r *queryResolver) MoneyMoves(ctx context.Context, after *string, before *string, filter *models.MoneyMoveFilter, first *int, last *int) (*models.MoneyMoveConnection, error) {
	return r.Domain.MoneyMoveRepo.GetMoneyMoves(filter, first, last, after, before)
}

// NearestStaffBirthdays is the resolver for the nearestStaffBirthdays field.
func (r *queryResolver) NearestStaffBirthdays(ctx context.Context, after *string, before *string, filter *models.StaffFilter, first *int, last *int) (*models.StaffConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// NearestStudentBirthdays is the resolver for the nearestStudentBirthdays field.
func (r *queryResolver) NearestStudentBirthdays(ctx context.Context, after *string, before *string, filter *models.StudentFilter, first *int, last *int) (*models.StudentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	return r.Domain.OrdersRepo.GetOrdersByID(id)
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, after *string, before *string, filter *models.OrderFilter, first *int, last *int) (*models.OrderConnection, error) {
	return r.Domain.OrdersRepo.GetOrders(filter, first, last, after, before)
}

// Place is the resolver for the place field.
func (r *queryResolver) Place(ctx context.Context, id string) (*models.Place, error) {
	return r.Domain.PlacesRepo.GetPlaceByID(id)
}

// PlaceAll is the resolver for the placeAll field.
func (r *queryResolver) PlaceAll(ctx context.Context) ([]*models.PlaceDto, error) {
	return r.Domain.PlacesRepo.All()
}

// Places is the resolver for the places field.
func (r *queryResolver) Places(ctx context.Context, after *string, before *string, filter *models.PlaceFilter, first *int, last *int) (*models.PlaceConnection, error) {
	return r.Domain.PlacesRepo.GetPlaces(filter, first, last, after, before)
}

// RentPaymentByMonth is the resolver for the rentPaymentByMonth field.
func (r *queryResolver) RentPaymentByMonth(ctx context.Context, id string) (*models.RentPaymentByMonth, error) {
	return r.Domain.RentPaymentByMonthRepo.GetRentPaymentByMonthByID(id)
}

// RentPaymentByTraining is the resolver for the rentPaymentByTraining field.
func (r *queryResolver) RentPaymentByTraining(ctx context.Context, id string) (*models.RentPaymentByTraining, error) {
	return r.Domain.RentPaymentByTrainingRepo.GetRentPaymentByTrainingByID(id)
}

// RentPaymentsByMonth is the resolver for the rentPaymentsByMonth field.
func (r *queryResolver) RentPaymentsByMonth(ctx context.Context, after *string, before *string, filter *models.RentPaymentByMonthFilter, first *int, last *int) (*models.RentPaymentByMonthConnection, error) {
	return r.Domain.RentPaymentByMonthRepo.GetRentPaymentByMonth(filter, first, last, after, before)
}

// RentPaymentsByTraining is the resolver for the rentPaymentsByTraining field.
func (r *queryResolver) RentPaymentsByTraining(ctx context.Context, after *string, before *string, filter *models.RentPaymentByTrainingFilter, first *int, last *int) (*models.RentPaymentByTrainingConnection, error) {
	return r.Domain.RentPaymentByTrainingRepo.GetRentPaymentByTraining(filter, first, last, after, before)
}

// Stadium is the resolver for the stadium field.
func (r *queryResolver) Stadium(ctx context.Context, id string) (*models.Stadium, error) {
	return r.Domain.StadiumsRepo.GetStadiumByID(id)
}

// StadiumAll is the resolver for the stadiumAll field.
func (r *queryResolver) StadiumAll(ctx context.Context) ([]*models.StadiumDto, error) {
	return r.Domain.StadiumsRepo.All()
}

// Stadiums is the resolver for the stadiums field.
func (r *queryResolver) Stadiums(ctx context.Context, after *string, before *string, filter *models.StadiumFilter, first *int, last *int) (*models.StadiumConnection, error) {
	return r.Domain.StadiumsRepo.GetStadiums(filter, first, last, after, before)
}

// Staff is the resolver for the staff field.
func (r *queryResolver) Staff(ctx context.Context, after *string, before *string, filter *models.StaffFilter, first *int, last *int) (*models.StaffConnection, error) {
	return r.Domain.StaffRepo.GetStaff(filter, first, last, after, before)
}

// StaffAll is the resolver for the staffAll field.
func (r *queryResolver) StaffAll(ctx context.Context) ([]*models.StaffDto, error) {
	return r.Domain.StaffRepo.All()
}

// StaffPerson is the resolver for the staffPerson field.
func (r *queryResolver) StaffPerson(ctx context.Context, id string) (*models.Staff, error) {
	return r.Domain.StaffRepo.GetStaffByID(id)
}

// Student is the resolver for the student field.
func (r *queryResolver) Student(ctx context.Context, id string) (*models.Student, error) {
	return r.Domain.StudentsRepo.GetStudentsByID(id)
}

// StudentAll is the resolver for the studentAll field.
func (r *queryResolver) StudentAll(ctx context.Context) ([]*models.StudentDto, error) {
	panic(fmt.Errorf("not implemented"))
}

// StudentVisit is the resolver for the studentVisit field.
func (r *queryResolver) StudentVisit(ctx context.Context, id string) (*models.StudentVisit, error) {
	return r.Domain.StudentVisitsRepo.GetStudentVisitByID(id)
}

// StudentVisits is the resolver for the studentVisits field.
func (r *queryResolver) StudentVisits(ctx context.Context, after *string, before *string, filter *models.StudentVisitFilter, first *int, last *int) (*models.StudentVisitConnection, error) {
	return r.Domain.StudentVisitsRepo.GetStudentVisits(filter, first, last, after, before)
}

// Students is the resolver for the students field.
func (r *queryResolver) Students(ctx context.Context, after *string, before *string, filter *models.StudentFilter, first *int, last *int) (*models.StudentConnection, error) {
	return r.Domain.StudentsRepo.GetStudents(filter, first, last, after, before)
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*models.Task, error) {
	return r.Domain.TasksRepo.GetTasksByID(id)
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, after *string, before *string, filter *models.TaskFilter, first *int, last *int) (*models.TaskConnection, error) {
	return r.Domain.TasksRepo.GetTasks(filter, first, last, after, before)
}

// Team is the resolver for the team field.
func (r *queryResolver) Team(ctx context.Context, id string) (*models.Team, error) {
	return r.Domain.TeamsRepo.GetTeamsByID(id)
}

// TeamAll is the resolver for the teamAll field.
func (r *queryResolver) TeamAll(ctx context.Context) ([]*models.TeamDto, error) {
	return r.Domain.TeamsRepo.All()
}

// TeamBalance is the resolver for the teamBalance field.
func (r *queryResolver) TeamBalance(ctx context.Context, id string) (*models.TeamBalance, error) {
	return r.Domain.TeamBalancesRepo.GetTeamBalanceByID(id)
}

// TeamBalances is the resolver for the teamBalances field.
func (r *queryResolver) TeamBalances(ctx context.Context, after *string, before *string, filter *models.TeamBalanceFilter, first *int, last *int) (*models.TeamBalanceConnection, error) {
	return r.Domain.TeamBalancesRepo.GetTeamBalances(filter, first, last, after, before)
}

// Teams is the resolver for the teams field.
func (r *queryResolver) Teams(ctx context.Context, after *string, before *string, filter *models.TeamFilter, first *int, last *int) (*models.TeamConnection, error) {
	return r.Domain.TeamsRepo.GetTeams(filter, first, last, after, before)
}

// TimeTable is the resolver for the timeTable field.
func (r *queryResolver) TimeTable(ctx context.Context, after *string, before *string, filter *models.TrainingFilter, first *int, last *int) (*models.TrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Training is the resolver for the training field.
func (r *queryResolver) Training(ctx context.Context, id string) (*models.Training, error) {
	return r.Domain.TrainingsRepo.GetTrainingsByID(id)
}

// TrainingAll is the resolver for the trainingAll field.
func (r *queryResolver) TrainingAll(ctx context.Context) ([]*models.TrainingDto, error) {
	return r.Domain.TrainingsRepo.All()
}

// TrainingDay is the resolver for the trainingDay field.
func (r *queryResolver) TrainingDay(ctx context.Context, id string) (*models.TrainingDay, error) {
	return r.Domain.TrainingDaysRepo.GetTrainingDaysByID(id)
}

// TrainingDays is the resolver for the trainingDays field.
func (r *queryResolver) TrainingDays(ctx context.Context, after *string, before *string, filter *models.TrainingDayFilter, first *int, last *int) (*models.TrainingDayConnection, error) {
	return r.Domain.TrainingDaysRepo.GetTrainingDays(filter, first, last, after, before)
}

// Trainings is the resolver for the trainings field.
func (r *queryResolver) Trainings(ctx context.Context, after *string, before *string, filter *models.TrainingFilter, first *int, last *int) (*models.TrainingConnection, error) {
	return r.Domain.TrainingsRepo.GetTrainings(filter, first, last, after, before)
}

// TrainingsByDay is the resolver for the trainingsByDay field.
func (r *queryResolver) TrainingsByDay(ctx context.Context, after *string, before *string, filter *models.TrainingFilter, first *int, last *int) (*models.TrainingConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// UnPayedStudents is the resolver for the unPayedStudents field.
func (r *queryResolver) UnPayedStudents(ctx context.Context, after *string, before *string, filter *models.StudentFilter, first *int, last *int) (*models.StudentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.Domain.UsersRepo.GetUserByID(id)
}

// UserAll is the resolver for the userAll field.
func (r *queryResolver) UserAll(ctx context.Context) ([]*models.UserDto, error) {
	return r.Domain.UsersRepo.GetUsersAll()
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *string, before *string, filter *models.UserFilter, first *int, last *int) (*models.UserConnection, error) {
	return r.Domain.UsersRepo.GetUsers(filter, first, last, after, before)
}

// Stadium is the resolver for the stadium field.
func (r *rentPaymentByMonthResolver) Stadium(ctx context.Context, obj *models.RentPaymentByMonth) (*models.Stadium, error) {
	return For(ctx).StadiumLoader.Load(obj.StadiumID)
}

// Stadium is the resolver for the stadium field.
func (r *rentPaymentByTrainingResolver) Stadium(ctx context.Context, obj *models.RentPaymentByTraining) (*models.Stadium, error) {
	return For(ctx).StadiumLoader.Load(obj.StadiumID)
}

// Trainings is the resolver for the trainings field.
func (r *rentPaymentByTrainingResolver) Trainings(ctx context.Context, obj *models.RentPaymentByTraining) ([]*models.Training, error) {
	return r.Domain.RentPaymentByTrainingRepo.GetTrainingsForRentPaymentByTraining(obj)
}

// Place is the resolver for the place field.
func (r *stadiumResolver) Place(ctx context.Context, obj *models.Stadium) (*models.Place, error) {
	return For(ctx).PlaceLoader.Load(*obj.PlaceID)
}

// User is the resolver for the user field.
func (r *staffResolver) User(ctx context.Context, obj *models.Staff) (*models.User, error) {
	return For(ctx).UserLoader.Load(*obj.UserId)
}

// Creators is the resolver for the creators field.
func (r *studentResolver) Creators(ctx context.Context, obj *models.Student) ([]*models.Creator, error) {
	return r.Domain.StudentsRepo.GetCreatorsForStudent(obj)
}

// Teams is the resolver for the teams field.
func (r *studentResolver) Teams(ctx context.Context, obj *models.Student) ([]*models.Team, error) {
	return r.Domain.StudentsRepo.GetTeamsForStudent(obj)
}

// Student is the resolver for the student field.
func (r *studentVisitResolver) Student(ctx context.Context, obj *models.StudentVisit) (*models.Student, error) {
	return For(ctx).StudentLoader.Load(obj.StudentID)
}

// Training is the resolver for the training field.
func (r *studentVisitResolver) Training(ctx context.Context, obj *models.StudentVisit) (*models.Training, error) {
	return For(ctx).TrainingLoader.Load(obj.TrainingID)
}

// Author is the resolver for the author field.
func (r *taskResolver) Author(ctx context.Context, obj *models.Task) (*models.User, error) {
	return For(ctx).UserLoader.Load(*obj.AuthorID)
}

// Leads is the resolver for the leads field.
func (r *taskResolver) Leads(ctx context.Context, obj *models.Task) ([]*models.Lead, error) {
	return r.Domain.TasksRepo.GetLeadsForTask(obj)
}

// Students is the resolver for the students field.
func (r *taskResolver) Students(ctx context.Context, obj *models.Task) ([]*models.Student, error) {
	return r.Domain.TasksRepo.GetStudentsForTask(obj)
}

// Workers is the resolver for the workers field.
func (r *taskResolver) Workers(ctx context.Context, obj *models.Task) ([]*models.Staff, error) {
	return r.Domain.TasksRepo.GetWorkersForTask(obj)
}

// Coaches is the resolver for the coaches field.
func (r *teamResolver) Coaches(ctx context.Context, obj *models.Team) ([]*models.Staff, error) {
	return r.Domain.TeamsRepo.GetCoachesForTeam(obj)
}

// HeadCoach is the resolver for the headCoach field.
func (r *teamResolver) HeadCoach(ctx context.Context, obj *models.Team) (*models.Staff, error) {
	return For(ctx).StaffLoader.Load(*obj.HeadCoachID)
}

// Place is the resolver for the place field.
func (r *teamResolver) Place(ctx context.Context, obj *models.Team) (*models.Place, error) {
	return For(ctx).PlaceLoader.Load(obj.PlaceID)
}

// Team is the resolver for the team field.
func (r *teamBalanceResolver) Team(ctx context.Context, obj *models.TeamBalance) (*models.Team, error) {
	return For(ctx).TeamLoader.Load(obj.TeamID)
}

// Coaches is the resolver for the coaches field.
func (r *trainingResolver) Coaches(ctx context.Context, obj *models.Training) ([]*models.Staff, error) {
	return r.Domain.TrainingsRepo.GetCoachesForTeam(obj)
}

// HeadCoach is the resolver for the headCoach field.
func (r *trainingResolver) HeadCoach(ctx context.Context, obj *models.Training) (*models.Staff, error) {
	return For(ctx).StaffLoader.Load(*obj.HeadCoachID)
}

// Stadium is the resolver for the stadium field.
func (r *trainingResolver) Stadium(ctx context.Context, obj *models.Training) (*models.Stadium, error) {
	return For(ctx).StadiumLoader.Load(*obj.StadiumID)
}

// Team is the resolver for the team field.
func (r *trainingResolver) Team(ctx context.Context, obj *models.Training) (*models.Team, error) {
	return For(ctx).TeamLoader.Load(obj.TeamID)
}

// Stadium is the resolver for the stadium field.
func (r *trainingDayResolver) Stadium(ctx context.Context, obj *models.TrainingDay) (*models.Stadium, error) {
	return For(ctx).StadiumLoader.Load(*obj.StadiumID)
}

// Team is the resolver for the team field.
func (r *trainingDayResolver) Team(ctx context.Context, obj *models.TrainingDay) (*models.Team, error) {
	return For(ctx).TeamLoader.Load(obj.TeamID)
}

// Article returns generated.ArticleResolver implementation.
func (r *Resolver) Article() generated.ArticleResolver { return &articleResolver{r} }

// Cart returns generated.CartResolver implementation.
func (r *Resolver) Cart() generated.CartResolver { return &cartResolver{r} }

// CoachPaymentByMonth returns generated.CoachPaymentByMonthResolver implementation.
func (r *Resolver) CoachPaymentByMonth() generated.CoachPaymentByMonthResolver {
	return &coachPaymentByMonthResolver{r}
}

// CoachPaymentByTeam returns generated.CoachPaymentByTeamResolver implementation.
func (r *Resolver) CoachPaymentByTeam() generated.CoachPaymentByTeamResolver {
	return &coachPaymentByTeamResolver{r}
}

// CoachPaymentByTraining returns generated.CoachPaymentByTrainingResolver implementation.
func (r *Resolver) CoachPaymentByTraining() generated.CoachPaymentByTrainingResolver {
	return &coachPaymentByTrainingResolver{r}
}

// Creator returns generated.CreatorResolver implementation.
func (r *Resolver) Creator() generated.CreatorResolver { return &creatorResolver{r} }

// Lead returns generated.LeadResolver implementation.
func (r *Resolver) Lead() generated.LeadResolver { return &leadResolver{r} }

// MoneyCost returns generated.MoneyCostResolver implementation.
func (r *Resolver) MoneyCost() generated.MoneyCostResolver { return &moneyCostResolver{r} }

// MoneyMove returns generated.MoneyMoveResolver implementation.
func (r *Resolver) MoneyMove() generated.MoneyMoveResolver { return &moneyMoveResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// RentPaymentByMonth returns generated.RentPaymentByMonthResolver implementation.
func (r *Resolver) RentPaymentByMonth() generated.RentPaymentByMonthResolver {
	return &rentPaymentByMonthResolver{r}
}

// RentPaymentByTraining returns generated.RentPaymentByTrainingResolver implementation.
func (r *Resolver) RentPaymentByTraining() generated.RentPaymentByTrainingResolver {
	return &rentPaymentByTrainingResolver{r}
}

// Stadium returns generated.StadiumResolver implementation.
func (r *Resolver) Stadium() generated.StadiumResolver { return &stadiumResolver{r} }

// Staff returns generated.StaffResolver implementation.
func (r *Resolver) Staff() generated.StaffResolver { return &staffResolver{r} }

// Student returns generated.StudentResolver implementation.
func (r *Resolver) Student() generated.StudentResolver { return &studentResolver{r} }

// StudentVisit returns generated.StudentVisitResolver implementation.
func (r *Resolver) StudentVisit() generated.StudentVisitResolver { return &studentVisitResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

// TeamBalance returns generated.TeamBalanceResolver implementation.
func (r *Resolver) TeamBalance() generated.TeamBalanceResolver { return &teamBalanceResolver{r} }

// Training returns generated.TrainingResolver implementation.
func (r *Resolver) Training() generated.TrainingResolver { return &trainingResolver{r} }

// TrainingDay returns generated.TrainingDayResolver implementation.
func (r *Resolver) TrainingDay() generated.TrainingDayResolver { return &trainingDayResolver{r} }

type articleResolver struct{ *Resolver }
type cartResolver struct{ *Resolver }
type coachPaymentByMonthResolver struct{ *Resolver }
type coachPaymentByTeamResolver struct{ *Resolver }
type coachPaymentByTrainingResolver struct{ *Resolver }
type creatorResolver struct{ *Resolver }
type leadResolver struct{ *Resolver }
type moneyCostResolver struct{ *Resolver }
type moneyMoveResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type rentPaymentByMonthResolver struct{ *Resolver }
type rentPaymentByTrainingResolver struct{ *Resolver }
type stadiumResolver struct{ *Resolver }
type staffResolver struct{ *Resolver }
type studentResolver struct{ *Resolver }
type studentVisitResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
type teamResolver struct{ *Resolver }
type teamBalanceResolver struct{ *Resolver }
type trainingResolver struct{ *Resolver }
type trainingDayResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	ErrInput = errors.New("inputs errors")
)
