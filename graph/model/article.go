package models

import "time"

type Article struct {
	AuthorID    string     `json:"authorId"`
	Description string     `json:"description"`
	FileName    *string    `json:"fileName"`
	ID          string     `json:"id"`
	Published   bool       `json:"published"`
	Tags        []string   `json:"tags"`
	Title       string     `json:"title"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
}
