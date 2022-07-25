package models

import "time"

type Kit struct {
	tableName struct{}   `pg:"kits,discard_unknown_columns"`
	FileName  *string    `json:"fileName"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Number    *int       `json:"number"`
	Price     int        `json:"price"`
	Published bool       `json:"published"`
	Quantity  *int       `json:"quantity"`
	Size      string     `json:"size"`
	Title     *string    `json:"title"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}
