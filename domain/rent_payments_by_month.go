package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func rentPaymentByMonthInputToRentPaymentByMonth(ctx context.Context, input *models.RentPaymentByMonthInput) *models.RentPaymentByMonth {
	return &models.RentPaymentByMonth{
		Description: input.Description,
		Month:       input.Month,
		PaymentDate: &input.PaymentDate,
		Published:   input.Published,
		StadiumID:   input.StadiumID,
		Sum:         input.Sum,
	}
}
func rentPaymentByMonthInputWithIdToRentPaymentByMonth(ctx context.Context, rentPaymentByMonth *models.RentPaymentByMonth, input *models.RentPaymentByMonthInput) *models.RentPaymentByMonth {
	return &models.RentPaymentByMonth{
		ID:          rentPaymentByMonth.ID,
		Description: input.Description,
		Month:       input.Month,
		PaymentDate: &input.PaymentDate,
		Published:   input.Published,
		StadiumID:   input.StadiumID,
		Sum:         input.Sum,
	}
}

func (d *Domain) RentPaymentByMonthSave(ctx context.Context, input models.RentPaymentByMonthInput) (*models.RentPaymentByMonthPayload, error) {
	return d.RentPaymentByMonthRepo.RentPaymentByMonthSave(rentPaymentByMonthInputToRentPaymentByMonth(ctx, &input))
}

func (d *Domain) RentPaymentByMonthUpdate(ctx context.Context, rentPaymentByMonthInput models.RentPaymentByMonthInputWithID) (*models.RentPaymentByMonthPayload, error) {

	rentPaymentByMonth, err := d.RentPaymentByMonthRepo.GetRentPaymentByMonthByID(rentPaymentByMonthInput.ID)

	rentPaymentByMonth = rentPaymentByMonthInputWithIdToRentPaymentByMonth(ctx, rentPaymentByMonth, rentPaymentByMonthInput.Input)
	rentPaymentByMonthPayload, err := d.RentPaymentByMonthRepo.RentPaymentByMonthUpdate(rentPaymentByMonth)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return rentPaymentByMonthPayload, nil
}

func (d *Domain) RentPaymentByMonthPublish(id string) (*models.RentPaymentByMonthPayload, error) {
	rentPaymentByMonth, err := d.RentPaymentByMonthRepo.GetRentPaymentByMonthByID(id)
	if err != nil || rentPaymentByMonth == nil {
		return nil, errors.New("филиала не существует")
	}

	rentPaymentByMonthPayload, err := d.RentPaymentByMonthRepo.RentPaymentByMonthPublish(rentPaymentByMonth)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return rentPaymentByMonthPayload, nil
}

func (d *Domain) RentPaymentByMonthRestore(id string) (*models.RentPaymentByMonthPayload, error) {
	rentPaymentByMonth, err := d.RentPaymentByMonthRepo.GetRentPaymentByMonthByID(id)
	if err != nil || rentPaymentByMonth == nil {
		return nil, errors.New("филиала не существует")
	}

	rentPaymentByMonthPayload, err := d.RentPaymentByMonthRepo.RentPaymentByMonthRestore(rentPaymentByMonth)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return rentPaymentByMonthPayload, nil
}

// RentPaymentByMonthDelete is the resolver for the rentPaymentByMonthDelete field.
func (d *Domain) RentPaymentByMonthDelete(id string) (bool, error) {
	rentPaymentByMonth, err := d.RentPaymentByMonthRepo.GetRentPaymentByMonthByID(id)
	if err != nil || rentPaymentByMonth == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.RentPaymentByMonthRepo.RentPaymentByMonthDelete(rentPaymentByMonth)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
