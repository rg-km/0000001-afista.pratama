package handlers

import (
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

// bit.ly/sdasdsadas
// -> google.com/aku/ganteng/sekali

// short : localhost:8080/
// long : localhost:8080/long/long/long

// dynamic route param

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
}
