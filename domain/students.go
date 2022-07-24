package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func studentInputToStudent(ctx context.Context, input *models.StudentInput) *models.Student {
	return &models.Student{
		Birthday:    input.Birthday,
		CreatorIds:  input.CreatorIds,
		Name:        input.Name,
		PassportNum: input.PassportNum,
		PaymentSum:  input.PaymentSum,
		Published:   input.Published,
		TeamIds:     input.TeamIds,
	}
}
func studentInputWithIdToStudent(ctx context.Context, student *models.Student, input *models.StudentInput) *models.Student {
	return &models.Student{
		ID:          student.ID,
		Birthday:    input.Birthday,
		CreatorIds:  input.CreatorIds,
		Name:        input.Name,
		PassportNum: input.PassportNum,
		PaymentSum:  input.PaymentSum,
		Published:   input.Published,
		TeamIds:     input.TeamIds,
	}
}

func (d *Domain) StudentSave(ctx context.Context, input models.StudentInput) (*models.StudentPayload, error) {
	return d.StudentsRepo.StudentsSave(studentInputToStudent(ctx, &input))
}

func (d *Domain) StudentUpdate(ctx context.Context, studentInput models.StudentInputWithID) (*models.StudentPayload, error) {

	student, err := d.StudentsRepo.GetStudentsByID(studentInput.ID)

	student = studentInputWithIdToStudent(ctx, student, studentInput.Input)
	studentPayload, err := d.StudentsRepo.StudentsUpdate(student)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return studentPayload, nil
}

func (d *Domain) StudentPublish(id string) (*models.StudentPayload, error) {
	student, err := d.StudentsRepo.GetStudentsByID(id)
	if err != nil || student == nil {
		return nil, errors.New("филиала не существует")
	}

	studentPayload, err := d.StudentsRepo.StudentsPublish(student)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return studentPayload, nil
}

func (d *Domain) StudentRestore(id string) (*models.StudentPayload, error) {
	student, err := d.StudentsRepo.GetStudentsByID(id)
	if err != nil || student == nil {
		return nil, errors.New("филиала не существует")
	}

	studentPayload, err := d.StudentsRepo.StudentsRestore(student)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return studentPayload, nil
}

// StudentDelete is the resolver for the studentDelete field.
func (d *Domain) StudentDelete(id string) (bool, error) {
	student, err := d.StudentsRepo.GetStudentsByID(id)
	if err != nil || student == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.StudentsRepo.StudentsDelete(student)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
