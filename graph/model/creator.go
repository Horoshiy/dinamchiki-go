package models

import "time"

type Creator struct {
	tableName   struct{}   `pg:"creators,discard_unknown_columns"`
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	PassportNum *string    `json:"passportNum"`
	Phone       string     `json:"phone"`
	Published   bool       `json:"published"`
	UserID      *string    `json:"userId"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
