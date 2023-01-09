-- name: Create :exec
INSERT INTO health(id, user_id, exercise_id, weight, `set`, `minute`, created_at) VALUES (?,?,?,?,?,?,?);

-- name: FetchByDateTime :many
SELECT h.*, e.subject FROM health h
         INNER JOIN exercise e ON h.exercise_id = e.id
         WHERE h.user_id = ? AND h.created_at BETWEEN ? AND ?
         ORDER BY h.created_at DESC;

-- name: GetTodayExerciseCount :one
SELECT COUNT(exercise_id) AS count FROM health
WHERE user_id = ? AND created_at BETWEEN ? AND ?;

-- name: FetchCategories :many
SELECT * FROM exercise_category;

-- name: FetchExerciseByCategoryId :many
SELECT * FROM exercise
WHERE category_id = ?;

-- name: FetchTodayExerciseByUserId :many
SELECT e.subject, COUNT(h.exercise_id) AS count FROM health h
    INNER JOIN exercise e ON h.exercise_id = e.id
    WHERE h.user_id = ? AND h.created_at BETWEEN ? AND ?
GROUP BY e.subject, h.exercise_id;