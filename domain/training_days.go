package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func trainingDaysInputToTrainingDay(ctx context.Context, input *models.TrainingDayInput) *models.TrainingDay {
	return &models.TrainingDay{
		Day:       input.Day,
		Published: input.Published,
		StadiumID: input.StadiumID,
		TeamID:    input.TeamID,
		Time:      input.Time,
	}
}
func trainingDaysInputWithIdToTrainingDay(ctx context.Context, trainingDays *models.TrainingDay, input *models.TrainingDayInput) *models.TrainingDay {
	return &models.TrainingDay{
		ID:        trainingDays.ID,
		Day:       input.Day,
		Published: input.Published,
		StadiumID: input.StadiumID,
		TeamID:    input.TeamID,
		Time:      input.Time,
	}
}

func (d *Domain) TrainingDaySave(ctx context.Context, input models.TrainingDayInput) (*models.TrainingDayPayload, error) {
	return d.TrainingDaysRepo.TrainingDaysSave(trainingDaysInputToTrainingDay(ctx, &input))
}

func (d *Domain) TrainingDayUpdate(ctx context.Context, trainingDaysInput models.TrainingDayInputWithID) (*models.TrainingDayPayload, error) {

	trainingDays, err := d.TrainingDaysRepo.GetTrainingDaysByID(trainingDaysInput.ID)

	trainingDays = trainingDaysInputWithIdToTrainingDay(ctx, trainingDays, trainingDaysInput.Input)
	trainingDaysPayload, err := d.TrainingDaysRepo.TrainingDaysUpdate(trainingDays)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return trainingDaysPayload, nil
}

func (d *Domain) TrainingDayPublish(id string) (*models.TrainingDayPayload, error) {
	trainingDays, err := d.TrainingDaysRepo.GetTrainingDaysByID(id)
	if err != nil || trainingDays == nil {
		return nil, errors.New("филиала не существует")
	}

	trainingDaysPayload, err := d.TrainingDaysRepo.TrainingDaysPublish(trainingDays)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return trainingDaysPayload, nil
}

func (d *Domain) TrainingDayRestore(id string) (*models.TrainingDayPayload, error) {
	trainingDays, err := d.TrainingDaysRepo.GetTrainingDaysByID(id)
	if err != nil || trainingDays == nil {
		return nil, errors.New("филиала не существует")
	}

	trainingDaysPayload, err := d.TrainingDaysRepo.TrainingDaysRestore(trainingDays)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return trainingDaysPayload, nil
}

// TrainingDayDelete is the resolver for the trainingDaysDelete field.
func (d *Domain) TrainingDayDelete(id string) (bool, error) {
	trainingDays, err := d.TrainingDaysRepo.GetTrainingDaysByID(id)
	if err != nil || trainingDays == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.TrainingDaysRepo.TrainingDaysDelete(trainingDays)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
