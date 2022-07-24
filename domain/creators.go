package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func creatorInputToCreator(ctx context.Context, input *models.CreatorInput) *models.Creator {
	return &models.Creator{
		Name:        input.Name,
		PassportNum: input.PassportNum,
		Phone:       input.Phone,
		Published:   input.Published,
		UserID:      input.UserID,
	}
}
func creatorInputWithIdToCreator(ctx context.Context, creator *models.Creator, input *models.CreatorInput) *models.Creator {
	return &models.Creator{
		ID:          creator.ID,
		Name:        input.Name,
		PassportNum: input.PassportNum,
		Phone:       input.Phone,
		Published:   input.Published,
		UserID:      input.UserID,
	}
}

func (d *Domain) CreatorSave(ctx context.Context, input models.CreatorInput) (*models.CreatorPayload, error) {
	return d.CreatorsRepo.CreatorSave(creatorInputToCreator(ctx, &input))
}

func (d *Domain) CreatorUpdate(ctx context.Context, creatorInput models.CreatorInputWithID) (*models.CreatorPayload, error) {

	creator, err := d.CreatorsRepo.GetCreatorByID(creatorInput.ID)

	creator = creatorInputWithIdToCreator(ctx, creator, creatorInput.Input)
	creatorPayload, err := d.CreatorsRepo.CreatorUpdate(creator)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return creatorPayload, nil
}

func (d *Domain) CreatorPublish(id string) (*models.CreatorPayload, error) {
	creator, err := d.CreatorsRepo.GetCreatorByID(id)
	if err != nil || creator == nil {
		return nil, errors.New("филиала не существует")
	}

	creatorPayload, err := d.CreatorsRepo.CreatorPublish(creator)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return creatorPayload, nil
}

func (d *Domain) CreatorRestore(id string) (*models.CreatorPayload, error) {
	creator, err := d.CreatorsRepo.GetCreatorByID(id)
	if err != nil || creator == nil {
		return nil, errors.New("филиала не существует")
	}

	creatorPayload, err := d.CreatorsRepo.CreatorRestore(creator)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return creatorPayload, nil
}

// CreatorDelete is the resolver for the creatorDelete field.
func (d *Domain) CreatorDelete(id string) (bool, error) {
	creator, err := d.CreatorsRepo.GetCreatorByID(id)
	if err != nil || creator == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.CreatorsRepo.CreatorDelete(creator)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
