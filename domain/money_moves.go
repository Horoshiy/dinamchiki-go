package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func moneyMoveInputToMoneyMove(ctx context.Context, input *models.MoneyMoveInput) *models.MoneyMove {
	return &models.MoneyMove{
		DateFinish:  &input.DateFinish,
		DatePayment: &input.DatePayment,
		DateStart:   &input.DateStart,
		Description: input.Description,
		MoneyForm:   &input.MoneyForm,
		OwnerID:     input.OwnerID,
		Published:   input.Published,
		StudentID:   input.StudentID,
		Sum:         &input.Sum,
		UserID:      input.UserID,
	}
}
func moneyMoveInputWithIdToMoneyMove(ctx context.Context, moneyMove *models.MoneyMove, input *models.MoneyMoveInput) *models.MoneyMove {
	return &models.MoneyMove{
		ID:          moneyMove.StudentID,
		DateFinish:  &input.DateFinish,
		DatePayment: &input.DatePayment,
		DateStart:   &input.DateStart,
		Description: input.Description,
		MoneyForm:   &input.MoneyForm,
		OwnerID:     input.OwnerID,
		Published:   input.Published,
		StudentID:   input.StudentID,
		Sum:         &input.Sum,
		UserID:      input.UserID,
	}
}

func (d *Domain) MoneyMoveSave(ctx context.Context, input models.MoneyMoveInput) (*models.MoneyMovePayload, error) {
	return d.MoneyMoveRepo.MoneyMoveSave(moneyMoveInputToMoneyMove(ctx, &input))
}

func (d *Domain) MoneyMoveUpdate(ctx context.Context, moneyMoveInput models.MoneyMoveInputWithID) (*models.MoneyMovePayload, error) {

	moneyMove, err := d.MoneyMoveRepo.GetMoneyMoveByID(moneyMoveInput.ID)

	moneyMove = moneyMoveInputWithIdToMoneyMove(ctx, moneyMove, moneyMoveInput.Input)
	moneyMovePayload, err := d.MoneyMoveRepo.MoneyMoveUpdate(moneyMove)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return moneyMovePayload, nil
}

func (d *Domain) MoneyMovePublish(id string) (*models.MoneyMovePayload, error) {
	moneyMove, err := d.MoneyMoveRepo.GetMoneyMoveByID(id)
	if err != nil || moneyMove == nil {
		return nil, errors.New("филиала не существует")
	}

	moneyMovePayload, err := d.MoneyMoveRepo.MoneyMovePublish(moneyMove)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return moneyMovePayload, nil
}

func (d *Domain) MoneyMoveRestore(id string) (*models.MoneyMovePayload, error) {
	moneyMove, err := d.MoneyMoveRepo.GetMoneyMoveByID(id)
	if err != nil || moneyMove == nil {
		return nil, errors.New("филиала не существует")
	}

	moneyMovePayload, err := d.MoneyMoveRepo.MoneyMoveRestore(moneyMove)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return moneyMovePayload, nil
}

// MoneyMoveDelete is the resolver for the moneyMoveDelete field.
func (d *Domain) MoneyMoveDelete(id string) (bool, error) {
	moneyMove, err := d.MoneyMoveRepo.GetMoneyMoveByID(id)
	if err != nil || moneyMove == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.MoneyMoveRepo.MoneyMoveDelete(moneyMove)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
