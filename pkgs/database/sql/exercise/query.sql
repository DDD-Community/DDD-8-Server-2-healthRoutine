-- name: Create :exec
INSERT INTO health(id, user_id, exercise_id, weight, `set`, `minute`, created_at) VALUES (?,?,?,?,?,?,?);

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
WHERE category_id = ?
LIMIT 8;

-- name: FetchTodayExerciseByUserId :many
SELECT
    ec.subject,
    e.subject,
    e.id,
    SUM(h.weight) AS weight,
    SUM(h.`set`) AS `set`,
    COUNT(h.exercise_id) AS count,
    h.created_at
FROM health h
         INNER JOIN exercise e ON h.exercise_id = e.id
         INNER JOIN exercise_category ec ON e.category_id = ec.id
WHERE h.user_id = ? AND h.created_at BETWEEN ? AND ?
GROUP BY h.exercise_id;

-- name: GetExerciseById :one
SELECT * FROM exercise
WHERE id = ?;

-- name: DeleteExercise :exec
DELETE FROM exercise
WHERE id = ? AND user_id = ?;

-- name: DeleteHealth :exec
DELETE FROM health
WHERE
    user_id = ? AND
    exercise_id = ? AND
    created_at BETWEEN ? AND ?;