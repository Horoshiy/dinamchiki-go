package models

import "time"

type CoachPaymentByMonth struct {
	tableName struct{}   `pg:"coach_payments_by_month,discard_unknown_columns"`
	CoachID   string     `json:"coachId"`
	Date      *time.Time `json:"date"`
	ID        string     `json:"id"`
	Published bool       `json:"published"`
	Sum       int        `json:"sum"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}
