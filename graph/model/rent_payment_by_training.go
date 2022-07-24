package models

import "time"

type RentPaymentByTraining struct {
	tableName   struct{}   `pg:"rent_payments_by_training,discard_unknown_columns"`
	Description *string    `json:"description"`
	ID          string     `json:"id"`
	Published   bool       `json:"published"`
	StadiumID   string     `json:"stadiumId"`
	Sum         int        `json:"sum"`
	TrainingIds []string   `json:"trainingIds" pg:",array"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
