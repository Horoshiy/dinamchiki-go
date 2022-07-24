package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func studentVisitInputToStudentVisit(ctx context.Context, input *models.StudentVisitInput) *models.StudentVisit {
	return &models.StudentVisit{
		Payed:       input.Payed,
		Published:   input.Published,
		StudentID:   input.StudentID,
		TrainingID:  input.TrainingID,
		VisitStatus: input.VisitStatus,
	}
}
func studentVisitInputWithIdToStudentVisit(ctx context.Context, studentVisit *models.StudentVisit, input *models.StudentVisitInput) *models.StudentVisit {
	return &models.StudentVisit{
		ID:          studentVisit.StudentID,
		Payed:       input.Payed,
		Published:   input.Published,
		StudentID:   input.StudentID,
		TrainingID:  input.TrainingID,
		VisitStatus: input.VisitStatus,
	}
}

func (d *Domain) StudentVisitSave(ctx context.Context, input models.StudentVisitInput) (*models.StudentVisitPayload, error) {
	return d.StudentVisitsRepo.StudentVisitSave(studentVisitInputToStudentVisit(ctx, &input))
}

func (d *Domain) StudentVisitUpdate(ctx context.Context, studentVisitInput models.StudentVisitInputWithID) (*models.StudentVisitPayload, error) {

	studentVisit, err := d.StudentVisitsRepo.GetStudentVisitByID(studentVisitInput.ID)

	studentVisit = studentVisitInputWithIdToStudentVisit(ctx, studentVisit, studentVisitInput.Input)
	studentVisitPayload, err := d.StudentVisitsRepo.StudentVisitUpdate(studentVisit)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return studentVisitPayload, nil
}

func (d *Domain) StudentVisitPublish(id string) (*models.StudentVisitPayload, error) {
	studentVisit, err := d.StudentVisitsRepo.GetStudentVisitByID(id)
	if err != nil || studentVisit == nil {
		return nil, errors.New("филиала не существует")
	}

	studentVisitPayload, err := d.StudentVisitsRepo.StudentVisitPublish(studentVisit)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return studentVisitPayload, nil
}

func (d *Domain) StudentVisitRestore(id string) (*models.StudentVisitPayload, error) {
	studentVisit, err := d.StudentVisitsRepo.GetStudentVisitByID(id)
	if err != nil || studentVisit == nil {
		return nil, errors.New("филиала не существует")
	}

	studentVisitPayload, err := d.StudentVisitsRepo.StudentVisitRestore(studentVisit)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return studentVisitPayload, nil
}

// StudentVisitDelete is the resolver for the studentVisitDelete field.
func (d *Domain) StudentVisitDelete(id string) (bool, error) {
	studentVisit, err := d.StudentVisitsRepo.GetStudentVisitByID(id)
	if err != nil || studentVisit == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.StudentVisitsRepo.StudentVisitDelete(studentVisit)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
