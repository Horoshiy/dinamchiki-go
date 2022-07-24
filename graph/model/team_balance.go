package models

import "time"

type TeamBalance struct {
	tableName struct{}   `pg:"team_balances,discard_unknown_columns"`
	Date      time.Time  `json:"date"`
	ID        string     `json:"id"`
	Published bool       `json:"published"`
	Rent      int        `json:"rent"`
	Salary    int        `json:"salary"`
	Sum       int        `json:"sum"`
	TeamID    string     `json:"teamId"`
	Tickets   int        `json:"tickets"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}
