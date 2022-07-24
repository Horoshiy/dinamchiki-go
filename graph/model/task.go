package models

import "time"

type Task struct {
	tableName   struct{}    `pg:"tasks,discard_unknown_columns"`
	AuthorID    *string     `json:"authorId"`
	Description *string     `json:"description"`
	EndTime     *time.Time  `json:"endTime"`
	ID          string      `json:"id"`
	LeadIds     []string    `json:"leadIds" pg:",array"`
	Priority    *Priority   `json:"priority"`
	Published   bool        `json:"published"`
	Result      *string     `json:"result"`
	StartTime   *time.Time  `json:"startTime"`
	StudentIds  []string    `json:"studentIds" pg:",array"`
	TaskStatus  *TaskStatus `json:"taskStatus"`
	Title       string      `json:"title"`
	WorkerIds   []string    `json:"workerIds" pg:",array"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
}
