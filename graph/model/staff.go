package models

import "time"

type Staff struct {
	tableName   struct{}   `pg:"staff,discard_unknown_columns"`
	ID          string     `json:"id"`
	Birthday    *time.Time `json:"birthday"`
	Department  Department `json:"department"`
	Description *string    `json:"description"`
	FileName    *string    `json:"fileName"`
	Name        string     `json:"name"`
	OrderNumber int        `json:"orderNumber"`
	PhoneNumber *string    `json:"phoneNumber"`
	Published   bool       `json:"published"`
	UserID      *string    `json:"userId"`
	Work        string     `json:"work"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
