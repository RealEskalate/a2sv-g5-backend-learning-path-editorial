package model

import "time"

type Status int

const (
	Pending = iota
	InProgress
	Completed
)

type Task struct {
	ID          uint16
	Title       string
	Description string
	Deadline    time.Time
	Status      Status
}
