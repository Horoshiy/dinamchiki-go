package models

import "time"

type Cart struct {
	tableName struct{}   `pg:"carts,discard_unknown_columns"`
	ID        string     `json:"id"`
	KitIds    []string   `json:"kitIds" pg:",array"`
	Published bool       `json:"published"`
	StudentID string     `json:"studentId"`
	Sum       int        `json:"sum"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}
