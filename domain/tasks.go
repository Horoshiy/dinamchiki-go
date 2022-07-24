package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func taskInputToTask(ctx context.Context, input *models.TaskInput) *models.Task {
	return &models.Task{
		AuthorID:    input.AuthorID,
		Description: input.Description,
		EndTime:     input.EndTime,
		LeadIds:     input.LeadIds,
		Priority:    input.Priority,
		Published:   input.Published,
		Result:      input.Result,
		StartTime:   input.StartTime,
		StudentIds:  input.StudentIds,
		TaskStatus:  input.TaskStatus,
		Title:       input.Title,
		WorkerIds:   input.WorkerIds,
	}
}
func taskInputWithIdToTask(ctx context.Context, task *models.Task, input *models.TaskInput) *models.Task {
	return &models.Task{
		ID:          task.ID,
		AuthorID:    input.AuthorID,
		Description: input.Description,
		EndTime:     input.EndTime,
		LeadIds:     input.LeadIds,
		Priority:    input.Priority,
		Published:   input.Published,
		Result:      input.Result,
		StartTime:   input.StartTime,
		StudentIds:  input.StudentIds,
		TaskStatus:  input.TaskStatus,
		Title:       input.Title,
		WorkerIds:   input.WorkerIds,
	}
}

func (d *Domain) TaskSave(ctx context.Context, input models.TaskInput) (*models.TaskPayload, error) {
	return d.TasksRepo.TasksSave(taskInputToTask(ctx, &input))
}

func (d *Domain) TaskUpdate(ctx context.Context, taskInput models.TaskInputWithID) (*models.TaskPayload, error) {

	task, err := d.TasksRepo.GetTasksByID(taskInput.ID)
	if err != nil || task == nil {
		return nil, errors.New("такого стадиона не существует")
	}

	task = taskInputWithIdToTask(ctx, task, taskInput.Input)
	taskPayload, err := d.TasksRepo.TasksUpdate(task)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return taskPayload, nil
}

func (d *Domain) TaskPublish(id string) (*models.TaskPayload, error) {
	task, err := d.TasksRepo.GetTasksByID(id)
	if err != nil || task == nil {
		return nil, errors.New("филиала не существует")
	}

	taskPayload, err := d.TasksRepo.TasksPublish(task)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return taskPayload, nil
}

func (d *Domain) TaskRestore(id string) (*models.TaskPayload, error) {
	task, err := d.TasksRepo.GetTasksByID(id)
	if err != nil || task == nil {
		return nil, errors.New("филиала не существует")
	}

	taskPayload, err := d.TasksRepo.TasksRestore(task)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return taskPayload, nil
}

// TaskDelete is the resolver for the taskDelete field.
func (d *Domain) TaskDelete(id string) (bool, error) {
	task, err := d.TasksRepo.GetTasksByID(id)
	if err != nil || task == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.TasksRepo.TasksDelete(task)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
