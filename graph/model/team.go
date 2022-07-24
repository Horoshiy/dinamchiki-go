package models

import "time"

type Team struct {
	tableName   struct{}   `pg:"teams,discard_unknown_columns"`
	ID          string     `json:"id"`
	Ages        []Age      `json:"ages" pg:",array"`
	CoachIds    []string   `json:"coachIds" pg:",array"`
	HeadCoachID *string    `json:"headCoachId"`
	Name        string     `json:"name"`
	PlaceID     string     `json:"placeId"`
	Published   bool       `json:"published"`
	Writable    bool       `json:"writable"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
