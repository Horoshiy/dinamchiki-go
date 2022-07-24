package models

import "time"

type Training struct {
	tableName   struct{}   `pg:"trainings,discard_unknown_columns"`
	ID          string     `json:"id"`
	CoachIds    []string   `json:"coachIds" pg:",array"`
	HeadCoachID *string    `json:"headCoachId"`
	Published   bool       `json:"published"`
	StadiumID   *string    `json:"stadiumId"`
	TeamID      string     `json:"teamId"`
	Squad       *Team      `pg:"-"`
	Time        *time.Time `json:"time"`
	Visits      int        `json:"visits"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
