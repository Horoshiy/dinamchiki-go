package models

import "time"

type Lead struct {
	tableName   struct{}    `pg:"leads,discard_unknown_columns"`
	Description *string     `json:"description"`
	ID          string      `json:"id"`
	Name        *string     `json:"name"`
	NextVisitID *string     `json:"nextVisitId"`
	Phone       string      `json:"phone"`
	Published   bool        `json:"published"`
	Source      *LeadSource `json:"source"`
	Status      *LeadStatus `json:"status"`
	StudentIds  []string    `json:"studentIds" pg:",array"`
	TeamID      *string     `json:"teamId"`
	YearBorn    *int        `json:"yearBorn"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
}
