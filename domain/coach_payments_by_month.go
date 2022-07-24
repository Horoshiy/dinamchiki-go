package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func coachPaymentByMonthInputToCoachPaymentByMonth(ctx context.Context, input *models.CoachPaymentByMonthInput) *models.CoachPaymentByMonth {
	return &models.CoachPaymentByMonth{
		CoachID:   input.CoachID,
		Date:      &input.Date,
		Published: input.Published,
		Sum:       input.Sum,
	}
}
func coachPaymentByMonthInputWithIdToCoachPaymentByMonth(ctx context.Context, coachPaymentByMonth *models.CoachPaymentByMonth, input *models.CoachPaymentByMonthInput) *models.CoachPaymentByMonth {
	return &models.CoachPaymentByMonth{
		ID:        coachPaymentByMonth.ID,
		CoachID:   input.CoachID,
		Date:      &input.Date,
		Published: input.Published,
		Sum:       input.Sum,
	}
}

func (d *Domain) CoachPaymentByMonthSave(ctx context.Context, input models.CoachPaymentByMonthInput) (*models.CoachPaymentByMonthPayload, error) {
	return d.CoachPaymentByMonthRepo.CoachPaymentByMonthSave(coachPaymentByMonthInputToCoachPaymentByMonth(ctx, &input))
}

func (d *Domain) CoachPaymentByMonthUpdate(ctx context.Context, coachPaymentByMonthInput models.CoachPaymentByMonthInputWithID) (*models.CoachPaymentByMonthPayload, error) {

	coachPaymentByMonth, err := d.CoachPaymentByMonthRepo.GetCoachPaymentByMonthByID(coachPaymentByMonthInput.ID)

	coachPaymentByMonth = coachPaymentByMonthInputWithIdToCoachPaymentByMonth(ctx, coachPaymentByMonth, coachPaymentByMonthInput.Input)
	coachPaymentByMonthPayload, err := d.CoachPaymentByMonthRepo.CoachPaymentByMonthUpdate(coachPaymentByMonth)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByMonthPayload, nil
}

func (d *Domain) CoachPaymentByMonthPublish(id string) (*models.CoachPaymentByMonthPayload, error) {
	coachPaymentByMonth, err := d.CoachPaymentByMonthRepo.GetCoachPaymentByMonthByID(id)
	if err != nil || coachPaymentByMonth == nil {
		return nil, errors.New("филиала не существует")
	}

	coachPaymentByMonthPayload, err := d.CoachPaymentByMonthRepo.CoachPaymentByMonthPublish(coachPaymentByMonth)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByMonthPayload, nil
}

func (d *Domain) CoachPaymentByMonthRestore(id string) (*models.CoachPaymentByMonthPayload, error) {
	coachPaymentByMonth, err := d.CoachPaymentByMonthRepo.GetCoachPaymentByMonthByID(id)
	if err != nil || coachPaymentByMonth == nil {
		return nil, errors.New("филиала не существует")
	}

	coachPaymentByMonthPayload, err := d.CoachPaymentByMonthRepo.CoachPaymentByMonthRestore(coachPaymentByMonth)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByMonthPayload, nil
}

// CoachPaymentByMonthDelete is the resolver for the coachPaymentByMonthDelete field.
func (d *Domain) CoachPaymentByMonthDelete(id string) (bool, error) {
	coachPaymentByMonth, err := d.CoachPaymentByMonthRepo.GetCoachPaymentByMonthByID(id)
	if err != nil || coachPaymentByMonth == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.CoachPaymentByMonthRepo.CoachPaymentByMonthDelete(coachPaymentByMonth)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
