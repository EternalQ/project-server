package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/eternalq/project-server/internal/api/models"
	"github.com/eternalq/project-server/internal/api/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type server struct {
	store  *store.Store
	router *gin.Engine
}

var (
	adminID = 16
)

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
	s.router.LoadHTMLGlob("/home/eternal/go/src/github.com/project-server/internal/api/server/templates/*.html")
	s.router.Static("/assets", "/home/eternal/go/src/github.com/project-server/internal/api/server/templates/assets")
	api := s.router.Group("/api")

	user := api.Group("/user")
	user.POST("/register", s.handleReg)
	user.POST("/login", s.handleLogin)
	user.GET("/posts/:id", s.handleUserPosts)
	user.GET("/albums/:id", s.handleUserAlbums)

	post := api.Group("/post")
	post.Use(s.checkAuth)
	post.GET("/all", s.handlePosts)
	post.POST("/create", s.handleCreatePost)
	post.GET("/delete/:id", s.handleDeletePost)
	post.GET("/find/:tag", s.handleFindPost)
	post.POST("/addcomment", s.handleCreatePostComment)
	post.GET("/getcomments/:id", s.handlePostComments)

	album := api.Group("/album")
	album.Use(s.checkAuth)
	album.POST("/create", s.handleCreateAlbum)
	album.GET("/delete/:id", s.handleDeleteAlbum)
	album.GET("/find/:id", s.handleFindAlbum)
	album.POST("/addphoto", s.handleAddPhoto)
	album.POST("/removephoto", s.handleRemovePhoto)

	admin := api.Group("/admin")
	admin.GET("/", s.handleAdminIndex)
	admin.POST("/", s.handleAdminAuth)
	admin.Use(s.chechAdmin)
	admin.GET("/users", s.handleAdminUsers)
	admin.GET("/posts", s.handleAdminPosts)
	admin.POST("/posts", s.handleAdminAddPost)
	admin.GET("/deletepost/:id", s.handleAdminDeletePost)
}

func (s *server) handleAdminDeletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pp, err := s.store.Post.Delete(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, pp)
}

func (s *server) handleAdminAddPost(ctx *gin.Context) {
	p := &models.Post{
		UserID:    16,
		CreatedAt: time.Now().Round(time.Minute),
	}

	p.Text, _ = ctx.GetPostForm("text")
	p.PhotoURL, _ = ctx.GetPostForm("photo_url")

	if err := s.store.Post.Create(p); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	ctx.Redirect(http.StatusFound, "/api/admin/posts")
}

func (s *server) handleAdminPosts(ctx *gin.Context) {
	pp, err := s.store.Post.GetLasts(100, 1)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "posts.html", pp)
}

func (s *server) handleAdminUsers(ctx *gin.Context) {
	uu, err := s.store.User.GetAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "users.html", uu)
}

func (s *server) handleAdminAuth(ctx *gin.Context) {
	login := ctx.Request.FormValue("login")
	password := ctx.Request.FormValue("password")
	if login != "admin" || password != "admin" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	IsAdminAuthorized = true

	time.AfterFunc(time.Hour, func() {
		IsAdminAuthorized = false
	})
	ctx.Redirect(http.StatusFound, "/api/admin/users")
}

func (s *server) handleAdminIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

var IsAdminAuthorized bool = false

func (s *server) chechAdmin(ctx *gin.Context) {
	if !IsAdminAuthorized {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (s *server) handleRemovePhoto(ctx *gin.Context) {
	req := &struct {
		ID       int    `json:"id"`
		PhotoUrl string `json:"photo_url"`
	}{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	a, err := s.store.Album.RemovePhoto(req.ID, req.PhotoUrl)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, a)
}

func (s *server) handleAddPhoto(ctx *gin.Context) {
	req := &struct {
		ID       int    `json:"id"`
		PhotoUrl string `json:"photo_url"`
	}{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	a, err := s.store.Album.AddPhoto(req.ID, req.PhotoUrl)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, a)
}

func (s *server) handleFindAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	a, err := s.store.Album.FindByID(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, a)
}

func (s *server) handleDeleteAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	aa, err := s.store.Album.Delete(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, aa)
}

func (s *server) handleCreateAlbum(ctx *gin.Context) {
	a := &models.Album{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(a); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	a.CreatedAt = time.Now().Round(time.Minute)

	if err := s.store.Album.Create(a); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, a)
}

func (s *server) handlePostComments(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cc, err := s.store.Comment.FindByPostID(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, cc)
}

func (s *server) handleCreatePostComment(ctx *gin.Context) {
	c := &models.Comment{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(c); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	c.CreatedAt = time.Now().Round(time.Minute)

	if err := s.store.Comment.Create(c); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, c)
}

func (s *server) handleFindPost(ctx *gin.Context) {
	tag := ctx.Param("tag")

	pp, err := s.store.Post.FindByTag(tag)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, pp)
}

func (s *server) handleDeletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pp, err := s.store.Post.Delete(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, pp)
}

func (s *server) handleCreatePost(ctx *gin.Context) {
	p := &models.Post{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(p); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	p.CreatedAt = time.Now().Round(time.Minute)

	if err := s.store.Post.Create(p); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, p)
}

func (s *server) handlePosts(ctx *gin.Context) {
	pp, err := s.store.Post.GetLasts(100, 1)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, pp)
}

func (s *server) checkAuth(ctx *gin.Context) {
	tokenstr := ctx.GetHeader("Authorization")[7:]
	if err := models.ValidateToken(tokenstr); err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}
}

func (s *server) handleUserAlbums(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	aa, err := s.store.User.GetAlbums(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, aa)
}

func (s *server) handleUserPosts(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pp, err := s.store.User.GetPosts(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, pp)
}

func (s *server) handleLogin(ctx *gin.Context) {
	req := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
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
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (s *server) handleReg(ctx *gin.Context) {
	u := &models.User{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(u); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	if err := s.store.User.Create(u); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := u.NewToken(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, u)
}
