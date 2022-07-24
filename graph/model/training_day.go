package models

import "time"

type TrainingDay struct {
	tableName struct{}   `pg:"training_days,discard_unknown_columns"`
	Day       *DayOfWeek `json:"day"`
	ID        string     `json:"id"`
	Published bool       `json:"published"`
	StadiumID *string    `json:"stadiumId"`
	TeamID    string     `json:"teamId"`
	Time      *time.Time `json:"time"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}
