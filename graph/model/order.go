package models

import "time"

type Order struct {
	tableName   struct{}    `pg:"orders,discard_unknown_columns"`
	CartID      string      `json:"cartId"`
	CreatorID   string      `json:"creatorId"`
	FileName    *string     `json:"fileName"`
	ID          string      `json:"id"`
	OrderStatus OrderStatus `json:"orderStatus"`
	Published   bool        `json:"published"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
}
