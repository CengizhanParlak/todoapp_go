package model

import "time"

type Todo struct {
	id        int
	task      string
	category  string
	status    Status
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
	createdBy string
}
