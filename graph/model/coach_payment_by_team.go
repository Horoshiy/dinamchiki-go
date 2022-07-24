package models

import "time"

type CoachPaymentByTeam struct {
	tableName   struct{}          `pg:"coach_payments_by_team,discard_unknown_columns"`
	CoachID     string            `json:"coachId"`
	DateFinish  *time.Time        `json:"dateFinish"`
	DateStart   *time.Time        `json:"dateStart"`
	ID          string            `json:"id"`
	PaymentRule *CoachPaymentRule `json:"paymentRule"`
	Published   bool              `json:"published"`
	Sum         *int              `json:"sum"`
	TeamID      *string           `json:"teamId"`
	DeletedAt   *time.Time        `json:"-" pg:",soft_delete"`
}
