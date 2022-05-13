package server

import (
	"encoding/json"
	"net/http"

	"github.com/eternalq/project-server/internal/api/models"
	"github.com/eternalq/project-server/internal/api/store"
	"github.com/gin-contrib/cors"
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
	s.router.Use(cors.Default())

	s.router.POST("/register", s.handleReg)
	s.router.POST("/login", s.handleLogin())
}

func (s *server) handleReg(ctx *gin.Context) {
	u := &models.User{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(u); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	if err := s.store.User.Create(u); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := u.NewToken(); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (s *server) handleLogin() gin.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(ctx *gin.Context) {
		req := &request{}
		if err := json.NewDecoder(ctx.Request.Body).Decode(req); err != nil {
			ctx.AbortWithError(http.StatusUnprocessableEntity, err)
			return
		}

		u, err := s.store.User.FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if err := u.NewToken(); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, u)
	}
}
