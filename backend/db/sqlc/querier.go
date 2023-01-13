// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	CreateDog(ctx context.Context, arg CreateDogParams) (Dog, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteFavorite(ctx context.Context, arg DeleteFavoriteParams) error
	GetDog(ctx context.Context, id int64) (Dog, error)
	GetFavoritesByDog(ctx context.Context, dogID int64) ([]Favorite, error)
	GetFavoritesByUser(ctx context.Context, userID int64) ([]Favorite, error)
	GetOwnedDogs(ctx context.Context, arg GetOwnedDogsParams) ([]Dog, error)
	GetSimilarDogs(ctx context.Context, arg GetSimilarDogsParams) ([]GetSimilarDogsRow, error)
	GetUser(ctx context.Context, email string) (User, error)
	UpdateDogLabels(ctx context.Context, arg UpdateDogLabelsParams) (Dog, error)
}

var _ Querier = (*Queries)(nil)
