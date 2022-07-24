package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func moneyCostInputToMoneyCost(ctx context.Context, input *models.MoneyCostInput) *models.MoneyCost {
	return &models.MoneyCost{
		Date:        input.Date,
		Description: input.Description,
		MoneyForm:   input.MoneyForm,
		Published:   input.Published,
		StaffID:     input.StaffID,
		Sum:         input.Sum,
	}
}
func moneyCostInputWithIdToMoneyCost(ctx context.Context, moneyCost *models.MoneyCost, input *models.MoneyCostInput) *models.MoneyCost {
	return &models.MoneyCost{
		ID:          moneyCost.ID,
		Date:        input.Date,
		Description: input.Description,
		MoneyForm:   input.MoneyForm,
		Published:   input.Published,
		StaffID:     input.StaffID,
		Sum:         input.Sum,
	}
}

func (d *Domain) MoneyCostSave(ctx context.Context, input models.MoneyCostInput) (*models.MoneyCostPayload, error) {
	return d.MoneyCostRepo.MoneyCostSave(moneyCostInputToMoneyCost(ctx, &input))
}

func (d *Domain) MoneyCostUpdate(ctx context.Context, moneyCostInput models.MoneyCostInputWithID) (*models.MoneyCostPayload, error) {

	moneyCost, err := d.MoneyCostRepo.GetMoneyCostByID(moneyCostInput.ID)

	moneyCost = moneyCostInputWithIdToMoneyCost(ctx, moneyCost, moneyCostInput.Input)
	moneyCostPayload, err := d.MoneyCostRepo.MoneyCostUpdate(moneyCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return moneyCostPayload, nil
}

func (d *Domain) MoneyCostPublish(id string) (*models.MoneyCostPayload, error) {
	moneyCost, err := d.MoneyCostRepo.GetMoneyCostByID(id)
	if err != nil || moneyCost == nil {
		return nil, errors.New("филиала не существует")
	}

	moneyCostPayload, err := d.MoneyCostRepo.MoneyCostPublish(moneyCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return moneyCostPayload, nil
}

func (d *Domain) MoneyCostRestore(id string) (*models.MoneyCostPayload, error) {
	moneyCost, err := d.MoneyCostRepo.GetMoneyCostByID(id)
	if err != nil || moneyCost == nil {
		return nil, errors.New("филиала не существует")
	}

	moneyCostPayload, err := d.MoneyCostRepo.MoneyCostRestore(moneyCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return moneyCostPayload, nil
}

// MoneyCostDelete is the resolver for the moneyCostDelete field.
func (d *Domain) MoneyCostDelete(id string) (bool, error) {
	moneyCost, err := d.MoneyCostRepo.GetMoneyCostByID(id)
	if err != nil || moneyCost == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.MoneyCostRepo.MoneyCostDelete(moneyCost)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
