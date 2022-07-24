package models

import "time"

type ClubBalance struct {
	tableName  struct{}   `pg:"club_balances,discard_unknown_columns"`
	Date       time.Time  `json:"date"`
	ID         string     `json:"id"`
	OtherCosts int        `json:"otherCosts"`
	Published  bool       `json:"published"`
	Rent       int        `json:"rent"`
	Salary     int        `json:"salary"`
	Sum        int        `json:"sum"`
	Tickets    int        `json:"tickets"`
	DeletedAt  *time.Time `json:"-" pg:",soft_delete"`
}
