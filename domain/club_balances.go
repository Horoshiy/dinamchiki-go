package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func clubBalanceInputToClubBalance(ctx context.Context, input *models.ClubBalanceInput) *models.ClubBalance {
	return &models.ClubBalance{
		Date:       input.Date,
		OtherCosts: input.OtherCosts,
		Published:  input.Published,
		Rent:       input.Rent,
		Salary:     input.Salary,
		Sum:        input.Tickets - input.Salary - input.Rent - input.OtherCosts,
		Tickets:    input.Tickets,
	}
}
func clubBalanceInputWithIdToClubBalance(ctx context.Context, clubBalance *models.ClubBalance, input *models.ClubBalanceInput) *models.ClubBalance {
	return &models.ClubBalance{
		ID:         clubBalance.ID,
		Date:       input.Date,
		OtherCosts: input.OtherCosts,
		Published:  input.Published,
		Rent:       input.Rent,
		Salary:     input.Salary,
		Sum:        input.Tickets - input.Salary - input.Rent - input.OtherCosts,
		Tickets:    input.Tickets,
	}
}

func (d *Domain) ClubBalanceSave(ctx context.Context, input models.ClubBalanceInput) (*models.ClubBalancePayload, error) {
	return d.ClubBalancesRepo.ClubBalanceSave(clubBalanceInputToClubBalance(ctx, &input))
}

func (d *Domain) ClubBalanceUpdate(ctx context.Context, clubBalanceInput models.ClubBalanceInputWithID) (*models.ClubBalancePayload, error) {

	clubBalance, err := d.ClubBalancesRepo.GetClubBalanceByID(clubBalanceInput.ID)

	clubBalance = clubBalanceInputWithIdToClubBalance(ctx, clubBalance, clubBalanceInput.Input)
	clubBalancePayload, err := d.ClubBalancesRepo.ClubBalanceUpdate(clubBalance)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return clubBalancePayload, nil
}

func (d *Domain) ClubBalancePublish(id string) (*models.ClubBalancePayload, error) {
	clubBalance, err := d.ClubBalancesRepo.GetClubBalanceByID(id)
	if err != nil || clubBalance == nil {
		return nil, errors.New("филиала не существует")
	}

	clubBalancePayload, err := d.ClubBalancesRepo.ClubBalancePublish(clubBalance)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return clubBalancePayload, nil
}

func (d *Domain) ClubBalanceRestore(id string) (*models.ClubBalancePayload, error) {
	clubBalance, err := d.ClubBalancesRepo.GetClubBalanceByID(id)
	if err != nil || clubBalance == nil {
		return nil, errors.New("филиала не существует")
	}

	clubBalancePayload, err := d.ClubBalancesRepo.ClubBalanceRestore(clubBalance)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return clubBalancePayload, nil
}

// ClubBalanceDelete is the resolver for the clubBalanceDelete field.
func (d *Domain) ClubBalanceDelete(id string) (bool, error) {
	clubBalance, err := d.ClubBalancesRepo.GetClubBalanceByID(id)
	if err != nil || clubBalance == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.ClubBalancesRepo.ClubBalanceDelete(clubBalance)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
