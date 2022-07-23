package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func placeInputToPlace(input *models.PlaceInput) *models.Place {
	return &models.Place{
		Address:     input.Address,
		Description: input.Description,
		Name:        input.Name,
		OrderNumber: input.OrderNumber,
		Published:   input.Published,
	}
}
func placeInputWithIdToPlace(place *models.Place, input *models.PlaceInput) *models.Place {
	return &models.Place{
		ID:          place.ID,
		Address:     input.Address,
		Description: input.Description,
		Name:        input.Name,
		OrderNumber: input.OrderNumber,
		Published:   input.Published,
	}
}
func (d *Domain) PlaceSave(input models.PlaceInput) (*models.PlacePayload, error) {

	_, err := d.PlacesRepo.GetPlaceByName(input.Name)
	if err == nil {
		return nil, errors.New("такое название уже есть у филиала")
	}

	return d.PlacesRepo.PlaceSave(placeInputToPlace(&input))
}

func (d *Domain) PlaceUpdate(ctx context.Context, placeInput models.PlaceInputWithID) (*models.PlacePayload, error) {

	place, err := d.PlacesRepo.GetPlaceByID(placeInput.ID)
	if err != nil || place == nil {
		return nil, errors.New("такого филиала не существует")
	}

	if len(placeInput.Input.Name) < 3 {
		return nil, errors.New("кол-во символов в названии мало")
	}
	if len(placeInput.Input.Description) < 3 {
		return nil, errors.New("кол-во символов в описании мало")
	}
	if len(placeInput.Input.Address) < 3 {
		return nil, errors.New("кол-во символов в адресе мало")
	}

	place = placeInputWithIdToPlace(place, placeInput.Input)
	placePayload, err := d.PlacesRepo.PlaceUpdate(place)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return placePayload, nil
}

func (d *Domain) PlacePublish(id string) (*models.PlacePayload, error) {
	place, err := d.PlacesRepo.GetPlaceByID(id)
	if err != nil || place == nil {
		return nil, errors.New("филиала не существует")
	}

	placePayload, err := d.PlacesRepo.PlacePublish(place)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return placePayload, nil
}

func (d *Domain) PlaceRestore(id string) (*models.PlacePayload, error) {
	place, err := d.PlacesRepo.GetPlaceByID(id)
	if err != nil || place == nil {
		return nil, errors.New("филиала не существует")
	}

	placePayload, err := d.PlacesRepo.PlaceRestore(place)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return placePayload, nil
}

// PlaceDelete is the resolver for the placeDelete field.
func (d *Domain) PlaceDelete(id string) (bool, error) {
	place, err := d.PlacesRepo.GetPlaceByID(id)
	if err != nil || place == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.PlacesRepo.PlaceDelete(place)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
