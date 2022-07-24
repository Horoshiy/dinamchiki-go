package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func coachPaymentByTrainingInputToCoachPaymentByTraining(ctx context.Context, input *models.CoachPaymentByTrainingInput) *models.CoachPaymentByTraining {
	return &models.CoachPaymentByTraining{
		CoachID:    input.CoachID,
		Published:  input.Published,
		Sum:        &input.Sum,
		TrainingID: &input.TrainingID,
	}
}
func coachPaymentByTrainingInputWithIdToCoachPaymentByTraining(ctx context.Context, coachPaymentByTraining *models.CoachPaymentByTraining, input *models.CoachPaymentByTrainingInput) *models.CoachPaymentByTraining {
	return &models.CoachPaymentByTraining{
		ID:         coachPaymentByTraining.ID,
		CoachID:    input.CoachID,
		Published:  input.Published,
		Sum:        &input.Sum,
		TrainingID: &input.TrainingID,
	}
}

func (d *Domain) CoachPaymentByTrainingSave(ctx context.Context, input models.CoachPaymentByTrainingInput) (*models.CoachPaymentByTrainingPayload, error) {
	return d.CoachPaymentByTrainingRepo.CoachPaymentByTrainingSave(coachPaymentByTrainingInputToCoachPaymentByTraining(ctx, &input))
}

func (d *Domain) CoachPaymentByTrainingUpdate(ctx context.Context, coachPaymentByTrainingInput models.CoachPaymentByTrainingInputWithID) (*models.CoachPaymentByTrainingPayload, error) {

	coachPaymentByTraining, err := d.CoachPaymentByTrainingRepo.GetCoachPaymentByTrainingByID(coachPaymentByTrainingInput.ID)

	coachPaymentByTraining = coachPaymentByTrainingInputWithIdToCoachPaymentByTraining(ctx, coachPaymentByTraining, coachPaymentByTrainingInput.Input)
	coachPaymentByTrainingPayload, err := d.CoachPaymentByTrainingRepo.CoachPaymentByTrainingUpdate(coachPaymentByTraining)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByTrainingPayload, nil
}

func (d *Domain) CoachPaymentByTrainingPublish(id string) (*models.CoachPaymentByTrainingPayload, error) {
	coachPaymentByTraining, err := d.CoachPaymentByTrainingRepo.GetCoachPaymentByTrainingByID(id)
	if err != nil || coachPaymentByTraining == nil {
		return nil, errors.New("филиала не существует")
	}

	coachPaymentByTrainingPayload, err := d.CoachPaymentByTrainingRepo.CoachPaymentByTrainingPublish(coachPaymentByTraining)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByTrainingPayload, nil
}

func (d *Domain) CoachPaymentByTrainingRestore(id string) (*models.CoachPaymentByTrainingPayload, error) {
	coachPaymentByTraining, err := d.CoachPaymentByTrainingRepo.GetCoachPaymentByTrainingByID(id)
	if err != nil || coachPaymentByTraining == nil {
		return nil, errors.New("филиала не существует")
	}

	coachPaymentByTrainingPayload, err := d.CoachPaymentByTrainingRepo.CoachPaymentByTrainingRestore(coachPaymentByTraining)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return coachPaymentByTrainingPayload, nil
}

// CoachPaymentByTrainingDelete is the resolver for the coachPaymentByTrainingDelete field.
func (d *Domain) CoachPaymentByTrainingDelete(id string) (bool, error) {
	coachPaymentByTraining, err := d.CoachPaymentByTrainingRepo.GetCoachPaymentByTrainingByID(id)
	if err != nil || coachPaymentByTraining == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.CoachPaymentByTrainingRepo.CoachPaymentByTrainingDelete(coachPaymentByTraining)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
