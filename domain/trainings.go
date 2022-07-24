package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func trainingInputToTraining(ctx context.Context, input *models.TrainingInput) *models.Training {
	return &models.Training{
		CoachIds:    input.CoachIds,
		HeadCoachID: input.HeadCoachID,
		Published:   input.Published,
		StadiumID:   input.StadiumID,
		TeamID:      input.TeamID,
		Time:        &input.Time,
		Visits:      input.Visits,
	}
}
func trainingInputWithIdToTraining(ctx context.Context, training *models.Training, input *models.TrainingInput) *models.Training {
	return &models.Training{
		ID:          training.ID,
		CoachIds:    input.CoachIds,
		HeadCoachID: input.HeadCoachID,
		Published:   input.Published,
		StadiumID:   input.StadiumID,
		TeamID:      input.TeamID,
		Time:        &input.Time,
		Visits:      input.Visits,
	}
}

func (d *Domain) TrainingSave(ctx context.Context, input models.TrainingInput) (*models.TrainingPayload, error) {
	return d.TrainingsRepo.TrainingsSave(trainingInputToTraining(ctx, &input))
}

func (d *Domain) TrainingUpdate(ctx context.Context, trainingInput models.TrainingInputWithID) (*models.TrainingPayload, error) {

	training, err := d.TrainingsRepo.GetTrainingsByID(trainingInput.ID)

	training = trainingInputWithIdToTraining(ctx, training, trainingInput.Input)
	trainingPayload, err := d.TrainingsRepo.TrainingsUpdate(training)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return trainingPayload, nil
}

func (d *Domain) TrainingPublish(id string) (*models.TrainingPayload, error) {
	training, err := d.TrainingsRepo.GetTrainingsByID(id)
	if err != nil || training == nil {
		return nil, errors.New("филиала не существует")
	}

	trainingPayload, err := d.TrainingsRepo.TrainingsPublish(training)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return trainingPayload, nil
}

func (d *Domain) TrainingRestore(id string) (*models.TrainingPayload, error) {
	training, err := d.TrainingsRepo.GetTrainingsByID(id)
	if err != nil || training == nil {
		return nil, errors.New("филиала не существует")
	}

	trainingPayload, err := d.TrainingsRepo.TrainingsRestore(training)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return trainingPayload, nil
}

// TrainingDelete is the resolver for the trainingDelete field.
func (d *Domain) TrainingDelete(id string) (bool, error) {
	training, err := d.TrainingsRepo.GetTrainingsByID(id)
	if err != nil || training == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.TrainingsRepo.TrainingsDelete(training)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
