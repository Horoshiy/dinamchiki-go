package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func orderInputToOrder(ctx context.Context, input *models.OrderInput) *models.Order {
	return &models.Order{
		CartID:      input.CartID,
		CreatorID:   input.CreatorID,
		FileName:    input.FileName,
		OrderStatus: input.OrderStatus,
		Published:   input.Published,
	}
}
func orderInputWithIdToOrder(ctx context.Context, order *models.Order, input *models.OrderInput) *models.Order {
	return &models.Order{
		ID:          order.ID,
		CartID:      input.CartID,
		CreatorID:   input.CreatorID,
		FileName:    input.FileName,
		OrderStatus: input.OrderStatus,
		Published:   input.Published,
	}
}

func (d *Domain) OrderSave(ctx context.Context, input models.OrderInput) (*models.OrderPayload, error) {
	return d.OrdersRepo.OrdersSave(orderInputToOrder(ctx, &input))
}

func (d *Domain) OrderUpdate(ctx context.Context, orderInput models.OrderInputWithID) (*models.OrderPayload, error) {

	order, err := d.OrdersRepo.GetOrdersByID(orderInput.ID)
	if err != nil || order == nil {
		return nil, errors.New("такого стадиона не существует")
	}

	order = orderInputWithIdToOrder(ctx, order, orderInput.Input)
	orderPayload, err := d.OrdersRepo.OrdersUpdate(order)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return orderPayload, nil
}

func (d *Domain) OrderPublish(id string) (*models.OrderPayload, error) {
	order, err := d.OrdersRepo.GetOrdersByID(id)
	if err != nil || order == nil {
		return nil, errors.New("филиала не существует")
	}

	orderPayload, err := d.OrdersRepo.OrdersPublish(order)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return orderPayload, nil
}

func (d *Domain) OrderRestore(id string) (*models.OrderPayload, error) {
	order, err := d.OrdersRepo.GetOrdersByID(id)
	if err != nil || order == nil {
		return nil, errors.New("филиала не существует")
	}

	orderPayload, err := d.OrdersRepo.OrdersRestore(order)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return orderPayload, nil
}

// OrderDelete is the resolver for the orderDelete field.
func (d *Domain) OrderDelete(id string) (bool, error) {
	order, err := d.OrdersRepo.GetOrdersByID(id)
	if err != nil || order == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.OrdersRepo.OrdersDelete(order)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
