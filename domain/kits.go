package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func kitInputToKit(ctx context.Context, input *models.KitInput) *models.Kit {
	return &models.Kit{
		FileName:  input.FileName,
		Name:      input.Name,
		Number:    input.Number,
		Price:     input.Price,
		Published: input.Published,
		Quantity:  input.Quantity,
		Size:      input.Size,
		Title:     input.Title,
	}
}
func kitInputWithIdToKit(ctx context.Context, kit *models.Kit, input *models.KitInput) *models.Kit {
	return &models.Kit{
		ID:        kit.ID,
		FileName:  input.FileName,
		Name:      input.Name,
		Number:    input.Number,
		Price:     input.Price,
		Published: input.Published,
		Quantity:  input.Quantity,
		Size:      input.Size,
		Title:     input.Title,
	}
}

func (d *Domain) KitSave(ctx context.Context, input models.KitInput) (*models.KitPayload, error) {
	return d.KitsRepo.KitsSave(kitInputToKit(ctx, &input))
}

func (d *Domain) KitUpdate(ctx context.Context, kitInput models.KitInputWithID) (*models.KitPayload, error) {

	kit, err := d.KitsRepo.GetKitsByID(kitInput.ID)
	if err != nil || kit == nil {
		return nil, errors.New("такого стадиона не существует")
	}

	kit = kitInputWithIdToKit(ctx, kit, kitInput.Input)
	kitPayload, err := d.KitsRepo.KitsUpdate(kit)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return kitPayload, nil
}

func (d *Domain) KitPublish(id string) (*models.KitPayload, error) {
	kit, err := d.KitsRepo.GetKitsByID(id)
	if err != nil || kit == nil {
		return nil, errors.New("филиала не существует")
	}

	kitPayload, err := d.KitsRepo.KitsPublish(kit)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return kitPayload, nil
}

func (d *Domain) KitRestore(id string) (*models.KitPayload, error) {
	kit, err := d.KitsRepo.GetKitsByID(id)
	if err != nil || kit == nil {
		return nil, errors.New("филиала не существует")
	}

	kitPayload, err := d.KitsRepo.KitsRestore(kit)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return kitPayload, nil
}

// KitDelete is the resolver for the kitDelete field.
func (d *Domain) KitDelete(id string) (bool, error) {
	kit, err := d.KitsRepo.GetKitsByID(id)
	if err != nil || kit == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.KitsRepo.KitsDelete(kit)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
