package api

import (
	"fmt"

	db "dog-recommend/db/sqlc"
	"dog-recommend/token"
	"dog-recommend/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// router.POST("/user", server.createUser)
	// router.POST("/users/login", server.loginUser)

	server.router = router
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
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
