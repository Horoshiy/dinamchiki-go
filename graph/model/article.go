package models

import "time"

type Article struct {
	tableName   struct{}   `pg:",discard_unknown_columns"`
	AuthorID    string     `json:"authorId"`
	Description string     `json:"description"`
	FileName    *string    `json:"fileName"`
	ID          string     `json:"id"`
	Published   bool       `json:"published"`
	Tags        []string   `json:"tags"`
	Title       string     `json:"title"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}

func (m *Article) IsOwner(user *User) bool {
	return m.AuthorID == user.ID
}
