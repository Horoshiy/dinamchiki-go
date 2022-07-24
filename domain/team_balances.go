package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func teamBalanceInputToTeamBalance(ctx context.Context, input *models.TeamBalanceInput) *models.TeamBalance {
	return &models.TeamBalance{
		Date:      input.Date,
		Published: input.Published,
		Rent:      input.Rent,
		Salary:    input.Salary,
		Sum:       input.Tickets - input.Salary - input.Rent,
		TeamID:    input.TeamID,
		Tickets:   input.Tickets,
	}
}
func teamBalanceInputWithIdToTeamBalance(ctx context.Context, teamBalance *models.TeamBalance, input *models.TeamBalanceInput) *models.TeamBalance {
	return &models.TeamBalance{
		ID:        teamBalance.ID,
		Date:      input.Date,
		Published: input.Published,
		Rent:      input.Rent,
		Salary:    input.Salary,
		Sum:       input.Tickets - input.Salary - input.Rent,
		TeamID:    input.TeamID,
		Tickets:   input.Tickets,
	}
}

func (d *Domain) TeamBalanceSave(ctx context.Context, input models.TeamBalanceInput) (*models.TeamBalancePayload, error) {
	return d.TeamBalancesRepo.TeamBalanceSave(teamBalanceInputToTeamBalance(ctx, &input))
}

func (d *Domain) TeamBalanceUpdate(ctx context.Context, teamBalanceInput models.TeamBalanceInputWithID) (*models.TeamBalancePayload, error) {

	teamBalance, err := d.TeamBalancesRepo.GetTeamBalanceByID(teamBalanceInput.ID)

	teamBalance = teamBalanceInputWithIdToTeamBalance(ctx, teamBalance, teamBalanceInput.Input)
	teamBalancePayload, err := d.TeamBalancesRepo.TeamBalanceUpdate(teamBalance)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return teamBalancePayload, nil
}

func (d *Domain) TeamBalancePublish(id string) (*models.TeamBalancePayload, error) {
	teamBalance, err := d.TeamBalancesRepo.GetTeamBalanceByID(id)
	if err != nil || teamBalance == nil {
		return nil, errors.New("филиала не существует")
	}

	teamBalancePayload, err := d.TeamBalancesRepo.TeamBalancePublish(teamBalance)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return teamBalancePayload, nil
}

func (d *Domain) TeamBalanceRestore(id string) (*models.TeamBalancePayload, error) {
	teamBalance, err := d.TeamBalancesRepo.GetTeamBalanceByID(id)
	if err != nil || teamBalance == nil {
		return nil, errors.New("филиала не существует")
	}

	teamBalancePayload, err := d.TeamBalancesRepo.TeamBalanceRestore(teamBalance)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return teamBalancePayload, nil
}

// TeamBalanceDelete is the resolver for the teamBalanceDelete field.
func (d *Domain) TeamBalanceDelete(id string) (bool, error) {
	teamBalance, err := d.TeamBalancesRepo.GetTeamBalanceByID(id)
	if err != nil || teamBalance == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.TeamBalancesRepo.TeamBalanceDelete(teamBalance)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
