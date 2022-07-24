package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func stadiumInputToStadium(ctx context.Context, input *models.StadiumInput) *models.Stadium {
	return &models.Stadium{
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Name:      input.Name,
		PlaceID:   &input.PlaceID,
		Published: input.Published,
	}
}
func stadiumInputWithIdToStadium(ctx context.Context, stadium *models.Stadium, input *models.StadiumInput) *models.Stadium {
	return &models.Stadium{
		ID:        stadium.ID,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Name:      input.Name,
		PlaceID:   &input.PlaceID,
		Published: input.Published,
	}
}

func (d *Domain) StadiumSave(ctx context.Context, input models.StadiumInput) (*models.StadiumPayload, error) {

	_, err := d.StadiumsRepo.GetStadiumByName(input.Name)
	if err == nil {
		return nil, errors.New("такое название уже есть у стадиона")
	}

	return d.StadiumsRepo.StadiumSave(stadiumInputToStadium(ctx, &input))
}

func (d *Domain) StadiumUpdate(ctx context.Context, stadiumInput models.StadiumInputWithID) (*models.StadiumPayload, error) {

	stadium, err := d.StadiumsRepo.GetStadiumByID(stadiumInput.ID)
	if err != nil || stadium == nil {
		return nil, errors.New("такого стадиона не существует")
	}

	if len(stadiumInput.Input.Name) < 3 {
		return nil, errors.New("кол-во символов в названии мало")
	}

	stadium = stadiumInputWithIdToStadium(ctx, stadium, stadiumInput.Input)
	stadiumPayload, err := d.StadiumsRepo.StadiumUpdate(stadium)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return stadiumPayload, nil
}

func (d *Domain) StadiumPublish(id string) (*models.StadiumPayload, error) {
	stadium, err := d.StadiumsRepo.GetStadiumByID(id)
	if err != nil || stadium == nil {
		return nil, errors.New("филиала не существует")
	}

	stadiumPayload, err := d.StadiumsRepo.StadiumPublish(stadium)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return stadiumPayload, nil
}

func (d *Domain) StadiumRestore(id string) (*models.StadiumPayload, error) {
	stadium, err := d.StadiumsRepo.GetStadiumByID(id)
	if err != nil || stadium == nil {
		return nil, errors.New("филиала не существует")
	}

	stadiumPayload, err := d.StadiumsRepo.StadiumRestore(stadium)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return stadiumPayload, nil
}

// StadiumDelete is the resolver for the stadiumDelete field.
func (d *Domain) StadiumDelete(id string) (bool, error) {
	stadium, err := d.StadiumsRepo.GetStadiumByID(id)
	if err != nil || stadium == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.StadiumsRepo.StadiumDelete(stadium)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
