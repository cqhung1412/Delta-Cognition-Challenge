-- CreateFavorite :one
INSERT INTO favorite (
  user_id,
  dog_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetFavoritesByDog :many
SELECT *
FROM favorite
WHERE dog_id = $1;

-- name: GetFavoritesByUser :many
SELECT *
FROM favorite
WHERE user_id = $1;

-- name: DeleteFavorite :exec
DELETE FROM favorite
WHERE user_id = $1 AND dog_id = $2;