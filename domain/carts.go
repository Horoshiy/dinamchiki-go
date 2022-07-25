package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func cartInputToCart(ctx context.Context, input *models.CartInput) *models.Cart {
	return &models.Cart{
		KitIds:    input.KitIds,
		Published: input.Published,
		StudentID: input.StudentID,
		Sum:       input.Sum,
	}
}
func cartInputWithIdToCart(ctx context.Context, cart *models.Cart, input *models.CartInput) *models.Cart {
	return &models.Cart{
		ID:        cart.ID,
		KitIds:    input.KitIds,
		Published: input.Published,
		StudentID: input.StudentID,
		Sum:       input.Sum,
	}
}

func (d *Domain) CartSave(ctx context.Context, input models.CartInput) (*models.CartPayload, error) {
	return d.CartsRepo.CartsSave(cartInputToCart(ctx, &input))
}

func (d *Domain) CartUpdate(ctx context.Context, cartInput models.CartInputWithID) (*models.CartPayload, error) {

	cart, err := d.CartsRepo.GetCartsByID(cartInput.ID)
	if err != nil || cart == nil {
		return nil, errors.New("такого стадиона не существует")
	}

	cart = cartInputWithIdToCart(ctx, cart, cartInput.Input)
	cartPayload, err := d.CartsRepo.CartsUpdate(cart)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return cartPayload, nil
}

func (d *Domain) CartPublish(id string) (*models.CartPayload, error) {
	cart, err := d.CartsRepo.GetCartsByID(id)
	if err != nil || cart == nil {
		return nil, errors.New("филиала не существует")
	}

	cartPayload, err := d.CartsRepo.CartsPublish(cart)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return cartPayload, nil
}

func (d *Domain) CartRestore(id string) (*models.CartPayload, error) {
	cart, err := d.CartsRepo.GetCartsByID(id)
	if err != nil || cart == nil {
		return nil, errors.New("филиала не существует")
	}

	cartPayload, err := d.CartsRepo.CartsRestore(cart)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return cartPayload, nil
}

// CartDelete is the resolver for the cartDelete field.
func (d *Domain) CartDelete(id string) (bool, error) {
	cart, err := d.CartsRepo.GetCartsByID(id)
	if err != nil || cart == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.CartsRepo.CartsDelete(cart)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
