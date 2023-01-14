package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	db "dog-recommend/db/sqlc"
	"dog-recommend/token"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type DogWithSignedUrl struct {
	Dog      db.Dog `json:"dog"`
	PhotoUrl string `json:"photo_url"`
}

type listDogsRequest struct {
	Offset int32 `form:"offset" binding:"min=0"`
	Limit  int32 `form:"limit" binding:"required,min=5"`
}

// @Summary	List dogs
// @Tags 	Dog
// @ID 		list-dog
// @Accept 	json
// @Produce json
// @Param 	data body listDogsRequest true "Request body"
// @Success 200 {array} []db.Dog
// @Router 	/dogs [get]
func (server *Server) listDogs(ctx *gin.Context) {
	var req listDogsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetDogsParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	dogs, err := server.store.GetDogs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create a new slice, Assign signed urls for each dog and append dog to new slice
	var dogsWithSignedUrl []DogWithSignedUrl
	for _, dog := range dogs {
		photo_url, err := server.awsMaker.CreateS3GetSignedUrl(fmt.Sprintf("%d.%s", dog.ID, dog.ImageType))
		if err != nil {
			fmt.Println("Cannot get signed url", err)
		}
		url_dog := DogWithSignedUrl{
			Dog:      dog,
			PhotoUrl: photo_url,
		}
		dogsWithSignedUrl = append(dogsWithSignedUrl, url_dog)
	}

	ctx.JSON(http.StatusOK, dogsWithSignedUrl)
}

type getDogRequest struct {
	DogID int64 `uri:"id" binding:"required,min=1"`
}

// @Summary	Get dog by ID
// @Tags 	Dog
// @ID 		get-dog
// @Accept 	json
// @Produce json
// @Param 	data body getDogRequest true "Request body"
// @Success 200 {object} db.Dog
// @Router 	/dog/:id [get]
func (server *Server) getDog(ctx *gin.Context) {
	var req getDogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	dog, err := server.store.GetDog(ctx, req.DogID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, dog)
}

type listDogRecommendationsRequest struct {
	DogID  int64 `form:"dog_id" binding:"required,min=1"`
	Offset int32 `form:"offset" binding:"min=0"`
	Limit  int32 `form:"limit" binding:"required,min=5"`
}

type SimilarDogWithSignedUrl struct {
	Dog      db.GetSimilarDogsRow `json:"dog"`
	PhotoUrl string               `json:"photo_url"`
}

// @Summary	List dog recommendations
// @Tags 	Dog
// @ID 		list-dog-recommendations
// @Accept 	json
// @Produce json
// @Param 	data body listDogRecommendationsRequest true "Request body"
// @Success 200 {array} []db.Dog
// @Router 	/recommend/dogs [get]
func (server *Server) listDogRecommendations(ctx *gin.Context) {
	var req listDogRecommendationsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetSimilarDogsParams{
		ID:     req.DogID,
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	dogs, err := server.store.GetSimilarDogs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create a new slice, Assign signed urls for each dog and append dog to new slice
	var dogsWithSignedUrl []SimilarDogWithSignedUrl
	for _, dog := range dogs {
		photo_url, err := server.awsMaker.CreateS3GetSignedUrl(fmt.Sprintf("%d.%s", dog.ID, dog.ImageType))
		if err != nil {
			fmt.Println("Cannot get signed url", err)
		}
		url_dog := SimilarDogWithSignedUrl{
			Dog:      dog,
			PhotoUrl: photo_url,
		}
		dogsWithSignedUrl = append(dogsWithSignedUrl, url_dog)
	}

	ctx.JSON(http.StatusOK, dogsWithSignedUrl)
}

type createDogRequest struct {
	Name      string `json:"name" binding:"required"`
	Breed     string `json:"breed" binding:"required"`
	BirthYear int32  `json:"birth_year" binding:"required"`
	ImageType string `json:"image_type" binding:"required"`
	Message   string `json:"message"`
}

type createDogResponse struct {
	Dog       db.Dog `json:"dog"`
	UploadUrl string `json:"upload_url"`
}

// @Summary	Create a new dog
// @Tags 	Dog
// @ID 		create-dog
// @Accept 	json
// @Produce json
// @Param 	data body createDogRequest true "Request body"
// @Success 200 {object} createDogResponse
// @Router 	/dog [post]
func (server *Server) createDog(ctx *gin.Context) {
	var req createDogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)

	arg := db.CreateDogParams{
		Name:      req.Name,
		Breed:     req.Breed,
		ImageType: req.ImageType,
		OwnerID: sql.NullInt64{
			Int64: authPayload.UserID,
			Valid: true,
		},
		Message: sql.NullString{
			String: req.Message,
			Valid:  req.Message != "",
		},
	}

	dog, err := server.store.CreateDog(ctx, arg)
	if err != nil {
		if _, ok := err.(*pq.Error); ok {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	uploadUrl, err := server.awsMaker.CreateS3UploadSignedUrl(fmt.Sprintf("%d.%s", dog.ID, req.ImageType))
	if err != nil {
		_ = server.store.DeleteDog(ctx, db.DeleteDogParams{
			ID: dog.ID,
			OwnerID: sql.NullInt64{
				Int64: authPayload.UserID,
				Valid: true,
			},
		})
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createDogResponse{
		Dog:       dog,
		UploadUrl: uploadUrl,
	}
	ctx.JSON(http.StatusOK, response)

	time.AfterFunc(5*time.Second, func() {
		dogJson, _ := server.awsMaker.ReadJsonFile("output-dog-recommend-012023", fmt.Sprintf("%d.%s.json", dog.ID, req.ImageType))
		var filteredLabels []string
		for _, label := range dogJson.Labels {
			if strings.Contains(strings.Join(label.Parents, " "), "Dog") {
				filteredLabels = append(filteredLabels, label.Name)
			}
		}
		fmt.Println("Filtered labels", filteredLabels)
		_, err := server.store.UpdateDogLabels(ctx, db.UpdateDogLabelsParams{
			ID:     dog.ID,
			Labels: filteredLabels,
			OwnerID: sql.NullInt64{
				Int64: authPayload.UserID,
				Valid: true,
			},
		})
		if err != nil {
			fmt.Println("Cannot update dog labels", err)
		}
	})
}

type updateDogLabelsRequest struct {
	DogID  int64    `json:"dog_id" binding:"required,min=1"`
	Labels []string `json:"labels" binding:"required"`
}

// @Summary	Update a dog's labels
// @Tags 	Dog
// @ID 		update-dog-labels
// @Accept 	json
// @Produce json
// @Param 	data body updateDogLabelsRequest true "Request body"
// @Success 200 {object} db.Dog
// @Router 	/dog/labels [put]
func (server *Server) updateDogLabels(ctx *gin.Context) {
	var req updateDogLabelsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)

	arg := db.UpdateDogLabelsParams{
		ID:     req.DogID,
		Labels: req.Labels,
		OwnerID: sql.NullInt64{
			Int64: authPayload.UserID,
			Valid: true,
		},
	}

	dog, err := server.store.UpdateDogLabels(ctx, arg)
	if err != nil {
		if _, ok := err.(*pq.Error); ok {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dog)
}

type deleteDogRequest struct {
	DogID int64 `uri:"id" binding:"required,min=1"`
}

// @Summary	Delete dog by ID
// @Tags 	Dog
// @ID 		get-dog
// @Accept 	json
// @Produce json
// @Param 	data body deleteDogRequest true "Request body"
// @Success 200 {object} db.Dog
// @Router 	/dog/:id [delete]
func (server *Server) deleteDog(ctx *gin.Context) {
	var req deleteDogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authPayloadKey).(*token.Payload)

	err := server.store.DeleteDog(ctx, db.DeleteDogParams{
		ID: req.DogID,
		OwnerID: sql.NullInt64{
			Int64: authPayload.UserID,
			Valid: true,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
