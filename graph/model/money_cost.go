package models

import "time"

type MoneyCost struct {
	tableName   struct{}   `pg:"money_costs,discard_unknown_columns"`
	Date        time.Time  `json:"date"`
	Description string     `json:"description"`
	ID          string     `json:"id"`
	MoneyForm   MoneyForm  `json:"moneyForm"`
	Published   bool       `json:"published"`
	StaffID     string     `json:"staffId"`
	Sum         int        `json:"sum"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
