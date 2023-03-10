// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type Dog struct {
	ID        int64          `json:"id"`
	OwnerID   sql.NullInt64  `json:"owner_id"`
	Name      string         `json:"name"`
	Breed     string         `json:"breed"`
	BirthYear int32          `json:"birth_year"`
	ImageType string         `json:"image_type"`
	Message   sql.NullString `json:"message"`
	// get from rekognition
	Labels    []string  `json:"labels"`
	CreatedAt time.Time `json:"created_at"`
}

type Favorite struct {
	UserID    int64     `json:"user_id"`
	DogID     int64     `json:"dog_id"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID             int64     `json:"id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
}
