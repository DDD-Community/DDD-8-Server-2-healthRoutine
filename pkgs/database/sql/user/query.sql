-- name: Create :exec
INSERT INTO users (id, nickname, email, password, profile_img, created_at, updated_at) VALUES (?,?,?,?,?,?,?);

-- name: GetById :one
SELECT * FROM users
WHERE id = ?;

-- name: GetByEmail :one
SELECT * from users
WHERE email = ?;

-- name: GetNicknameById :one
SELECT nickname from users
WHERE id = ?;

-- name: CheckExistsByEmail :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = ?
           ) AS isExists;

-- name: CheckExistsByNickname :one
SELECT EXISTS(
   SELECT * FROM users
   WHERE nickname = ?
           ) AS isExists;

-- name: UpdateProfileById :exec
UPDATE users
SET nickname = ?, profile_img = ?, updated_at = ?
WHERE id = ?;

-- name: CreateBadge :exec
INSERT INTO badge_users (users_id, badge_id, created_at) VALUES (?, ?, ?);

-- name: GetBadgeByUserId :many
SELECT bu.badge_id FROM badge_users bu
INNER JOIN badge b on bu.badge_id = b.id
WHERE users_id = ?
ORDER BY b.id;

-- name: GetLatestBadgeByUserId :one
SELECT b.id, subject FROM badge_users bu
    INNER JOIN badge b on bu.badge_id = b.id
WHERE users_id = ?
ORDER BY created_at LIMIT 1;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = ?