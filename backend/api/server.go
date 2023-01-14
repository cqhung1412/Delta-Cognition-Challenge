package api

import (
	"fmt"

	aws_util "dog-recommend/aws"
	db "dog-recommend/db/sqlc"
	"dog-recommend/token"
	"dog-recommend/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	awsMaker   aws_util.AwsMaker
	router     *gin.Engine
}

func (server *Server) setupRouter() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	router.POST("/signin", server.createUser)
	router.POST("/login", server.loginUser)
	authRoutes.GET("/user", server.getUser)

	authRoutes.POST("/dog", server.createDog)
	authRoutes.PUT("/dog/labels", server.updateDogLabels)
	router.GET("/dog/:id", server.getDog)
	router.GET("/dogs", server.listDogs)
	router.GET("/recommend/dogs", server.listDogRecommendations)
	authRoutes.DELETE("/dog/:id", server.deleteDog)

	server.router = router
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	awsMaker, err := aws_util.NewAwsMaker(config.AWSAccessKeyID, config.AWSSecretAccessKey, config.AWSRegion, config.S3Bucket)
	if err != nil {
		return nil, fmt.Errorf("cannot create AWS maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		awsMaker:   *awsMaker,
	}

	server.setupRouter()
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
