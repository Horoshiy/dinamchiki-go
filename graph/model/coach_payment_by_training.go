package models

import "time"

type CoachPaymentByTraining struct {
	tableName  struct{}   `pg:"coach_payments_by_training,discard_unknown_columns"`
	CoachID    string     `json:"coachId"`
	ID         string     `json:"id"`
	Published  bool       `json:"published"`
	Sum        *int       `json:"sum"`
	TrainingID *string    `json:"trainingId"`
	DeletedAt  *time.Time `json:"-" pg:",soft_delete"`
}
