package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func coachPaymentByTeamInputToCoachPaymentByTeam(ctx context.Context, input *models.CoachPaymentByTeamInput) *models.CoachPaymentByTeam {
	return &models.CoachPaymentByTeam{
		CoachID:     input.CoachID,
		DateFinish:  &input.DateFinish,
		DateStart:   &input.DateStart,
		PaymentRule: &input.PaymentRule,
		Published:   input.Published,
		Sum:         &input.Sum,
		TeamID:      &input.TeamID,
	}
}
func coachPaymentByTeamInputWithIdToCoachPaymentByTeam(ctx context.Context, coachPaymentByTeam *models.CoachPaymentByTeam, input *models.CoachPaymentByTeamInput) *models.CoachPaymentByTeam {
	return &models.CoachPaymentByTeam{
		ID:          coachPaymentByTeam.ID,
		CoachID:     input.CoachID,
		DateFinish:  &input.DateFinish,
		DateStart:   &input.DateStart,
		PaymentRule: &input.PaymentRule,
		Published:   input.Published,
		Sum:         &input.Sum,
		TeamID:      &input.TeamID,
	}
}

func (d *Domain) CoachPaymentByTeamSave(ctx context.Context, input models.CoachPaymentByTeamInput) (*models.CoachPaymentByTeamPayload, error) {
	return d.CoachPaymentByTeamRepo.CoachPaymentByTeamSave(coachPaymentByTeamInputToCoachPaymentByTeam(ctx, &input))
}

func (d *Domain) CoachPaymentByTeamUpdate(ctx context.Context, coachPaymentByTeamInput models.CoachPaymentByTeamInputWithID) (*models.CoachPaymentByTeamPayload, error) {

	coachPaymentByTeam, err := d.CoachPaymentByTeamRepo.GetCoachPaymentByTeamByID(coachPaymentByTeamInput.ID)

	coachPaymentByTeam = coachPaymentByTeamInputWithIdToCoachPaymentByTeam(ctx, coachPaymentByTeam, coachPaymentByTeamInput.Input)
	coachPaymentByTeamPayload, err := d.CoachPaymentByTeamRepo.CoachPaymentByTeamUpdate(coachPaymentByTeam)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByTeamPayload, nil
}

func (d *Domain) CoachPaymentByTeamPublish(id string) (*models.CoachPaymentByTeamPayload, error) {
	coachPaymentByTeam, err := d.CoachPaymentByTeamRepo.GetCoachPaymentByTeamByID(id)
	if err != nil || coachPaymentByTeam == nil {
		return nil, errors.New("филиала не существует")
	}

	coachPaymentByTeamPayload, err := d.CoachPaymentByTeamRepo.CoachPaymentByTeamPublish(coachPaymentByTeam)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByTeamPayload, nil
}

func (d *Domain) CoachPaymentByTeamRestore(id string) (*models.CoachPaymentByTeamPayload, error) {
	coachPaymentByTeam, err := d.CoachPaymentByTeamRepo.GetCoachPaymentByTeamByID(id)
	if err != nil || coachPaymentByTeam == nil {
		return nil, errors.New("филиала не существует")
	}

	coachPaymentByTeamPayload, err := d.CoachPaymentByTeamRepo.CoachPaymentByTeamRestore(coachPaymentByTeam)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByTeamPayload, nil
}

// CoachPaymentByTeamDelete is the resolver for the coachPaymentByTeamDelete field.
func (d *Domain) CoachPaymentByTeamDelete(id string) (bool, error) {
	coachPaymentByTeam, err := d.CoachPaymentByTeamRepo.GetCoachPaymentByTeamByID(id)
	if err != nil || coachPaymentByTeam == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.CoachPaymentByTeamRepo.CoachPaymentByTeamDelete(coachPaymentByTeam)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
