package models

import "time"

type Stadium struct {
	tableName struct{}   `pg:"stadiums,discard_unknown_columns"`
	ID        string     `json:"id"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Name      string     `json:"name"`
	PlaceID   *string    `json:"placeId"`
	Published bool       `json:"published"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}
