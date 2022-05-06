package server

import (
	"net/http"

	"github.com/eternalq/project-server/internal/api/store"
	"github.com/gin-gonic/gin"
)

type server struct {
	store  *store.Store
	router *gin.Engine
}

func newServer(store *store.Store) *server {
	s := &server{
		store:  store,
		router: gin.Default(),
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "got it")
	})
}
