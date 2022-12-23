-- name: Create :exec
INSERT INTO users (id, nickname, email, password, profile_img, created_at, updated_at) VALUES (?,?,?,?,?,?,?);

-- name: GetById :one
SELECT * FROM users
WHERE id = ?;

-- name: GetByEmail :one
SELECT * from users
WHERE email = ?;

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

-- name: UpdateProfileImgById :exec
UPDATE users
SET profile_img = ?, updated_at = ?
WHERE id = ?;

-- name: UpdateNicknameById :exec
UPDATE users
SET nickname = ?, updated_at = ?
WHERE id = ?;