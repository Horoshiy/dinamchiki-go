package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func staffInputToStaff(ctx context.Context, input *models.StaffInput) *models.Staff {
	return &models.Staff{
		Birthday:    input.Birthday,
		Department:  input.Department,
		Description: input.Description,
		FileName:    input.FileName,
		Name:        input.Name,
		OrderNumber: input.OrderNumber,
		PhoneNumber: input.PhoneNumber,
		Published:   input.Published,
		UserId:      input.UserID,
		Work:        input.Work,
	}
}
func staffInputWithIdToStaff(ctx context.Context, staff *models.Staff, input *models.StaffInput) *models.Staff {
	return &models.Staff{
		ID:          staff.ID,
		Birthday:    input.Birthday,
		Department:  input.Department,
		Description: input.Description,
		FileName:    input.FileName,
		Name:        input.Name,
		OrderNumber: input.OrderNumber,
		PhoneNumber: input.PhoneNumber,
		Published:   input.Published,
		UserId:      input.UserID,
		Work:        input.Work,
	}
}

func (d *Domain) StaffSave(ctx context.Context, input models.StaffInput) (*models.StaffPayload, error) {

	_, err := d.StaffRepo.GetStaffByName(input.Name)
	if err == nil {
		return nil, errors.New("такое название уже есть у стадиона")
	}

	return d.StaffRepo.StaffSave(staffInputToStaff(ctx, &input))
}

func (d *Domain) StaffUpdate(ctx context.Context, staffInput models.StaffInputWithID) (*models.StaffPayload, error) {

	staff, err := d.StaffRepo.GetStaffByID(staffInput.ID)
	if err != nil || staff == nil {
		return nil, errors.New("такого стадиона не существует")
	}

	if len(staffInput.Input.Name) < 3 {
		return nil, errors.New("кол-во символов в названии мало")
	}

	staff = staffInputWithIdToStaff(ctx, staff, staffInput.Input)
	staffPayload, err := d.StaffRepo.StaffUpdate(staff)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return staffPayload, nil
}

func (d *Domain) StaffPublish(id string) (*models.StaffPayload, error) {
	staff, err := d.StaffRepo.GetStaffByID(id)
	if err != nil || staff == nil {
		return nil, errors.New("филиала не существует")
	}

	staffPayload, err := d.StaffRepo.StaffPublish(staff)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return staffPayload, nil
}

func (d *Domain) StaffRestore(id string) (*models.StaffPayload, error) {
	staff, err := d.StaffRepo.GetStaffByID(id)
	if err != nil || staff == nil {
		return nil, errors.New("филиала не существует")
	}

	staffPayload, err := d.StaffRepo.StaffRestore(staff)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return staffPayload, nil
}

// StaffDelete is the resolver for the staffDelete field.
func (d *Domain) StaffDelete(id string) (bool, error) {
	staff, err := d.StaffRepo.GetStaffByID(id)
	if err != nil || staff == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.StaffRepo.StaffDelete(staff)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
