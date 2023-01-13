-- name: CreateDog :one
INSERT INTO dogs 
  (
    owner_id,
    image,
    name,
    breed,
    birth_year,
    message
  )
VALUES 
  ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: UpdateDogLabels :one
UPDATE dogs
SET labels = $1
WHERE id = $2
RETURNING *;

-- name: GetDog :one
SELECT *
FROM dogs
WHERE id = $1
LIMIT 1;

-- name: GetOwnedDogs :many
SELECT *
FROM dogs
WHERE owner_id = $1
LIMIT $1
OFFSET $2;

-- name: GetSimilarDogs :many
SELECT
  id,
  image,
  name,
  birth_year,
  breed,
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
  END,
  similarity DESC
LIMIT $2
OFFSET $3;