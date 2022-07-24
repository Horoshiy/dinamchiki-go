package models

import "time"

type MoneyMove struct {
	tableName   struct{}   `pg:"money_moves,discard_unknown_columns"`
	ID          string     `json:"id"`
	DateFinish  *time.Time `json:"dateFinish"`
	DatePayment *time.Time `json:"datePayment"`
	DateStart   *time.Time `json:"dateStart"`
	Description *string    `json:"description"`
	MoneyForm   *MoneyForm `json:"moneyForm"`
	OwnerID     string     `json:"ownerId"`
	Published   bool       `json:"published"`
	StudentID   string     `json:"studentId"`
	Sum         *int       `json:"sum"`
	UserID      string     `json:"userId"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
