package models

import "time"

type StudentVisit struct {
	tableName   struct{}    `pg:"student_visits,discard_unknown_columns"`
	ID          string      `json:"id"`
	Payed       bool        `json:"payed"`
	Published   bool        `json:"published"`
	StudentID   string      `json:"studentId"`
	TrainingID  string      `json:"trainingId"`
	VisitStatus VisitStatus `json:"visitStatus"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
}
