package models

import "time"

type Student struct {
	tableName   struct{}   `pg:"students,discard_unknown_columns"`
	ID          string     `json:"id"`
	Birthday    *time.Time `json:"birthday"`
	CreatorIds  []string   `json:"creatorIds" pg:",array"`
	Name        string     `json:"name"`
	PassportNum *string    `json:"passportNum"`
	PaymentSum  *int       `json:"paymentSum"`
	Published   bool       `json:"published"`
	TeamIds     []string   `json:"teamIds" pg:",array"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
