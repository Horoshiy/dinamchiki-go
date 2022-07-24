package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func rentPaymentByTrainingInputToRentPaymentByTraining(ctx context.Context, input *models.RentPaymentByTrainingInput) *models.RentPaymentByTraining {
	return &models.RentPaymentByTraining{
		Description: input.Description,
		Published:   input.Published,
		StadiumID:   input.StadiumID,
		Sum:         input.Sum,
		TrainingIds: input.TrainingIds,
	}
}
func rentPaymentByTrainingInputWithIdToRentPaymentByTraining(ctx context.Context, rentPaymentByTraining *models.RentPaymentByTraining, input *models.RentPaymentByTrainingInput) *models.RentPaymentByTraining {
	return &models.RentPaymentByTraining{
		ID:          rentPaymentByTraining.ID,
		Description: input.Description,
		Published:   input.Published,
		StadiumID:   input.StadiumID,
		Sum:         input.Sum,
		TrainingIds: input.TrainingIds,
	}
}

func (d *Domain) RentPaymentByTrainingSave(ctx context.Context, input models.RentPaymentByTrainingInput) (*models.RentPaymentByTrainingPayload, error) {
	return d.RentPaymentByTrainingRepo.RentPaymentByTrainingSave(rentPaymentByTrainingInputToRentPaymentByTraining(ctx, &input))
}

func (d *Domain) RentPaymentByTrainingUpdate(ctx context.Context, rentPaymentByTrainingInput models.RentPaymentByTrainingInputWithID) (*models.RentPaymentByTrainingPayload, error) {

	rentPaymentByTraining, err := d.RentPaymentByTrainingRepo.GetRentPaymentByTrainingByID(rentPaymentByTrainingInput.ID)

	rentPaymentByTraining = rentPaymentByTrainingInputWithIdToRentPaymentByTraining(ctx, rentPaymentByTraining, rentPaymentByTrainingInput.Input)
	rentPaymentByTrainingPayload, err := d.RentPaymentByTrainingRepo.RentPaymentByTrainingUpdate(rentPaymentByTraining)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return rentPaymentByTrainingPayload, nil
}

func (d *Domain) RentPaymentByTrainingPublish(id string) (*models.RentPaymentByTrainingPayload, error) {
	rentPaymentByTraining, err := d.RentPaymentByTrainingRepo.GetRentPaymentByTrainingByID(id)
	if err != nil || rentPaymentByTraining == nil {
		return nil, errors.New("филиала не существует")
	}

	rentPaymentByTrainingPayload, err := d.RentPaymentByTrainingRepo.RentPaymentByTrainingPublish(rentPaymentByTraining)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return rentPaymentByTrainingPayload, nil
}

func (d *Domain) RentPaymentByTrainingRestore(id string) (*models.RentPaymentByTrainingPayload, error) {
	rentPaymentByTraining, err := d.RentPaymentByTrainingRepo.GetRentPaymentByTrainingByID(id)
	if err != nil || rentPaymentByTraining == nil {
		return nil, errors.New("филиала не существует")
	}

	rentPaymentByTrainingPayload, err := d.RentPaymentByTrainingRepo.RentPaymentByTrainingRestore(rentPaymentByTraining)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return rentPaymentByTrainingPayload, nil
}

// RentPaymentByTrainingDelete is the resolver for the rentPaymentByTrainingDelete field.
func (d *Domain) RentPaymentByTrainingDelete(id string) (bool, error) {
	rentPaymentByTraining, err := d.RentPaymentByTrainingRepo.GetRentPaymentByTrainingByID(id)
	if err != nil || rentPaymentByTraining == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.RentPaymentByTrainingRepo.RentPaymentByTrainingDelete(rentPaymentByTraining)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
