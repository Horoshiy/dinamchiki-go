package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func (d *Domain) CreatePlace(input models.PlaceInput) (*models.PlacePayload, error) {

	_, err := d.PlacesRepo.GetPlaceByName(input.Name)
	if err == nil {
		return nil, errors.New("Такой название уже есть у филиала")
	}

	place := &models.Place{
		Address:     input.Address,
		Description: input.Description,
		Name:        input.Name,
		OrderNumber: input.OrderNumber,
		Published:   input.Published,
	}
	return d.PlacesRepo.CreatePlace(place)
}

func (d *Domain) UpdatePlace(ctx context.Context, id string, input *models.PlaceInput) (*models.Place, error) {

	place, err := d.PlacesRepo.GetPlaceByID(id)
	if err != nil || place == nil {
		return nil, errors.New("такого филиала не существует")
	}

	if len(input.Name) < 3 {
		return nil, errors.New("кол-во символов в названии мало")
	}
	if len(input.Description) < 3 {
		return nil, errors.New("кол-во символов в описании мало")
	}
	if len(input.Address) < 3 {
		return nil, errors.New("кол-во символов в адресе мало")
	}

	place.Name = input.Name
	place.Description = input.Description
	place.Address = input.Address
	place.OrderNumber = input.OrderNumber
	place.Published = input.Published
	place, err = d.PlacesRepo.UpdatePlace(place)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return place, nil
}

func (d *Domain) DeletePlace(ctx context.Context, id string) (bool, error) {
	place, err := d.PlacesRepo.GetPlaceByID(id)
	if err != nil || place == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.PlacesRepo.DeletePlace(place)
	if err != nil {
		return false, fmt.Errorf("ошибка при удалении филиала %v", err)
	}

	return true, nil
}
