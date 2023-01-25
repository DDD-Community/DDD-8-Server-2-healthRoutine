// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package exercise

import (
	"context"

	"github.com/google/uuid"
)

const create = `-- name: Create :exec
INSERT INTO health(id, user_id, exercise_id, weight, reps, ` + "`" + `set` + "`" + `, created_at) VALUES (?,?,?,?,?,?,?)
`

type CreateParams struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	ExerciseID int64
	Weight     int32
	Reps       int32
	Set        int32
	CreatedAt  int64
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) error {
	_, err := q.exec(ctx, q.createStmt, create,
		arg.ID,
		arg.UserID,
		arg.ExerciseID,
		arg.Weight,
		arg.Reps,
		arg.Set,
		arg.CreatedAt,
	)
	return err
}

const createExercise = `-- name: CreateExercise :exec
INSERT INTO exercise(id, subject, category_id, user_id) VALUES (?,?,?,?)
`

type CreateExerciseParams struct {
	ID         int64
	Subject    string
	CategoryID int64
	UserID     *uuid.UUID
}

func (q *Queries) CreateExercise(ctx context.Context, arg CreateExerciseParams) error {
	_, err := q.exec(ctx, q.createExerciseStmt, createExercise,
		arg.ID,
		arg.Subject,
		arg.CategoryID,
		arg.UserID,
	)
	return err
}

const deleteExercise = `-- name: DeleteExercise :exec
DELETE FROM exercise
WHERE id = ? AND user_id = ?
`

type DeleteExerciseParams struct {
	ID     int64
	UserID *uuid.UUID
}

func (q *Queries) DeleteExercise(ctx context.Context, arg DeleteExerciseParams) error {
	_, err := q.exec(ctx, q.deleteExerciseStmt, deleteExercise, arg.ID, arg.UserID)
	return err
}

const deleteHealth = `-- name: DeleteHealth :exec
DELETE FROM health
WHERE id = ?
`

func (q *Queries) DeleteHealth(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteHealthStmt, deleteHealth, id)
	return err
}

const fetchByDateTime = `-- name: FetchByDateTime :many
SELECT
    COUNT(exercise_id) AS counts,
    DATE_FORMAT(FROM_UNIXTIME(created_at/1000), '%d') AS day
FROM health
WHERE user_id = ? AND created_at BETWEEN ? AND ?
GROUP BY day
ORDER BY day
`

type FetchByDateTimeParams struct {
	UserID      uuid.UUID
	CreatedAt   int64
	CreatedAt_2 int64
}

type FetchByDateTimeRow struct {
	Counts int64
	Day    string
}

func (q *Queries) FetchByDateTime(ctx context.Context, arg FetchByDateTimeParams) ([]FetchByDateTimeRow, error) {
	rows, err := q.query(ctx, q.fetchByDateTimeStmt, fetchByDateTime, arg.UserID, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchByDateTimeRow
	for rows.Next() {
		var i FetchByDateTimeRow
		if err := rows.Scan(&i.Counts, &i.Day); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchCategories = `-- name: FetchCategories :many
SELECT id, subject FROM exercise_category
`

func (q *Queries) FetchCategories(ctx context.Context) ([]ExerciseCategory, error) {
	rows, err := q.query(ctx, q.fetchCategoriesStmt, fetchCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ExerciseCategory
	for rows.Next() {
		var i ExerciseCategory
		if err := rows.Scan(&i.ID, &i.Subject); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchExerciseByCategoryId = `-- name: FetchExerciseByCategoryId :many
SELECT id, subject, category_id, user_id FROM exercise
WHERE category_id = ?
LIMIT 8
`

func (q *Queries) FetchExerciseByCategoryId(ctx context.Context, categoryID int64) ([]Exercise, error) {
	rows, err := q.query(ctx, q.fetchExerciseByCategoryIdStmt, fetchExerciseByCategoryId, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exercise
	for rows.Next() {
		var i Exercise
		if err := rows.Scan(
			&i.ID,
			&i.Subject,
			&i.CategoryID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchTodayExerciseByUserId = `-- name: FetchTodayExerciseByUserId :many
SELECT h.id, h.user_id, h.exercise_id, h.weight, h.reps, h.` + "`" + `set` + "`" + `, h.created_at,
       ec.subject AS category_subject,
       e.subject AS exercise_subject
FROM health h
         INNER JOIN exercise e ON h.exercise_id = e.id
         INNER JOIN exercise_category ec ON e.category_id = ec.id
WHERE h.user_id = ? AND h.created_at BETWEEN ? AND ?
ORDER BY h.created_at
`

type FetchTodayExerciseByUserIdParams struct {
	UserID      uuid.UUID
	CreatedAt   int64
	CreatedAt_2 int64
}

type FetchTodayExerciseByUserIdRow struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	ExerciseID      int64
	Weight          int32
	Reps            int32
	Set             int32
	CreatedAt       int64
	CategorySubject string
	ExerciseSubject string
}

func (q *Queries) FetchTodayExerciseByUserId(ctx context.Context, arg FetchTodayExerciseByUserIdParams) ([]FetchTodayExerciseByUserIdRow, error) {
	rows, err := q.query(ctx, q.fetchTodayExerciseByUserIdStmt, fetchTodayExerciseByUserId, arg.UserID, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchTodayExerciseByUserIdRow
	for rows.Next() {
		var i FetchTodayExerciseByUserIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ExerciseID,
			&i.Weight,
			&i.Reps,
			&i.Set,
			&i.CreatedAt,
			&i.CategorySubject,
			&i.ExerciseSubject,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExerciseById = `-- name: GetExerciseById :one
SELECT id, subject, category_id, user_id FROM exercise
WHERE id = ?
`

func (q *Queries) GetExerciseById(ctx context.Context, id int64) (Exercise, error) {
	row := q.queryRow(ctx, q.getExerciseByIdStmt, getExerciseById, id)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.Subject,
		&i.CategoryID,
		&i.UserID,
	)
	return i, err
}

const getTodayExerciseCount = `-- name: GetTodayExerciseCount :one
SELECT COUNT(exercise_id) AS count FROM health
WHERE user_id = ? AND created_at BETWEEN ? AND ?
`

type GetTodayExerciseCountParams struct {
	UserID      uuid.UUID
	CreatedAt   int64
	CreatedAt_2 int64
}

func (q *Queries) GetTodayExerciseCount(ctx context.Context, arg GetTodayExerciseCountParams) (int64, error) {
	row := q.queryRow(ctx, q.getTodayExerciseCountStmt, getTodayExerciseCount, arg.UserID, arg.CreatedAt, arg.CreatedAt_2)
	var count int64
	err := row.Scan(&count)
	return count, err
}
