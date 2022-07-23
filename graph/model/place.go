package models

import (
	"time"
)

type Place struct {
	tableName   struct{}   `pg:",discard_unknown_columns"`
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Address     string     `json:"address"`
	Published   bool       `json:"published"`
	OrderNumber int        `json:"orderNumber"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
