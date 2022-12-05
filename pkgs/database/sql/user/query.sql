-- name: Create :exec
INSERT INTO users (id, nickname, email, password, profile_img, created_at, updated_at) VALUES (?,?,?,?,?,?,?);

-- name: GetByEmail :one
SELECT * from users
WHERE email = ?;

-- name: CheckExistsByEmail :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = ?
           ) AS isExist;

-- name: CheckExistsByNickname :one
SELECT EXISTS(
   SELECT * FROM users
   WHERE nickname = ?
           ) AS isExsist;