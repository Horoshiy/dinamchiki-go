package models

import "time"

type RentPaymentByMonth struct {
	tableName   struct{}   `pg:"rent_payments_by_month,discard_unknown_columns"`
	ID          string     `json:"id"`
	Description *string    `json:"description"`
	Month       time.Time  `json:"month"`
	PaymentDate *time.Time `json:"paymentDate"`
	Published   bool       `json:"published"`
	StadiumID   string     `json:"stadiumId"`
	Sum         int        `json:"sum"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
