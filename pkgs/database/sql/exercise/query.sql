-- name: Create :exec
INSERT INTO health(id, user_id, exercise_id, weight, reps, `set`, created_at) VALUES (?,?,?,?,?,?,?);

-- name: CreateExercise :exec
INSERT INTO exercise(id, subject, category_id, user_id) VALUES (?,?,?,?);

-- name: FetchByDateTime :many
SELECT
    COUNT(exercise_id) AS counts,
    DATE_FORMAT(FROM_UNIXTIME(created_at/1000), '%d') AS day
FROM health
WHERE user_id = ? AND created_at BETWEEN ? AND ?
GROUP BY day
ORDER BY day;

-- name: GetTodayExerciseCount :one
SELECT COUNT(exercise_id) AS count FROM health
WHERE user_id = ? AND created_at BETWEEN ? AND ?;

-- name: FetchCategories :many
SELECT * FROM exercise_category;

-- name: FetchExerciseByCategoryId :many
SELECT * FROM exercise
WHERE category_id = ? AND (user_id = ? OR user_id IS NULL)
ORDER BY user_id IS NULL DESC;

-- name: FetchTodayExerciseByUserId :many
SELECT h.*,
       ec.subject AS category_subject,
       e.subject AS exercise_subject
FROM health h
         INNER JOIN exercise e ON h.exercise_id = e.id
         INNER JOIN exercise_category ec ON e.category_id = ec.id
WHERE h.user_id = ? AND h.created_at BETWEEN ? AND ?
ORDER BY h.created_at;

-- name: GetExerciseById :one
SELECT * FROM exercise
WHERE id = ?;

-- name: DeleteExercise :exec
DELETE FROM exercise
WHERE id = ? AND user_id = ?;

-- name: DeleteHealth :exec
DELETE FROM health
WHERE id = ?;

-- name: GetWaterByUserId :one
SELECT * FROM water
WHERE user_id = ? AND created_at BETWEEN ? AND ?;

-- name: CreateOrUpdateWater :exec
INSERT INTO water(user_id, capacity, unit, `date`, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE capacity = ? , updated_at = ?;

-- name: CountExerciseHistoryByUserId :one
SELECT COUNT(user_id) AS exercise_count FROM health
WHERE user_id = ?;

-- name: CountDrinkHistoryByUserId :one
SELECT COUNT(user_id) AS drink_count FROM water
WHERE user_id = ?;

