-- name: CreateDog :one
INSERT INTO dogs 
  (
    owner_id,
    name,
    breed,
    birth_year,
    image_type,
    message
  )
VALUES 
  ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: UpdateDogLabels :one
UPDATE dogs
SET labels = $1
WHERE id = $2 AND owner_id = $3
RETURNING *;

-- name: GetDog :one
SELECT *
FROM dogs
WHERE id = $1
LIMIT 1;

-- name: GetOwnedDogs :many
SELECT *
FROM dogs
WHERE owner_id = $1;

-- name: GetDogs :many
SELECT *
FROM dogs
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: GetSimilarDogs :many
SELECT
  id,
  name,
  birth_year,
  breed,
  image_type,
  message,
  labels
FROM dogs
WHERE
  (
    labels && (
      SELECT labels
      FROM dogs
      WHERE dogs.id = $1
    )
    OR breed = (
      SELECT breed
      FROM dogs
      WHERE dogs.id = $1
    )
  )
ORDER BY
  CASE
    WHEN labels && (
      SELECT labels
      FROM dogs
      WHERE dogs.id = $1
    ) THEN 1
    ELSE 2
  END DESC,
  id DESC
LIMIT $2
OFFSET $3;

-- name: DeleteDog :exec
DELETE FROM dogs
WHERE id = $1 AND owner_id = $2;