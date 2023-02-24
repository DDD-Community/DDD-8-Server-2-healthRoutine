// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package exercise

import (
	"github.com/google/uuid"
)

type Exercise struct {
	ID         int64
	Subject    string
	CategoryID int64
	UserID     *uuid.UUID
}

type ExerciseCategory struct {
	ID      int64
	Subject string
}

type Health struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	ExerciseID int64
	Weight     int32
	Reps       int32
	Set        int32
	CreatedAt  int64
}

type Water struct {
	UserID    uuid.UUID
	Capacity  int64
	Unit      string
	Date      string
	CreatedAt int64
	UpdatedAt int64
}
