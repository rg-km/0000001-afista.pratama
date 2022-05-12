package handlers

import (
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
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
	path := c.Param("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "path is required",
		})
	}

	url, err := h.repo.Get(path)
	if err != nil {
		// internal server error
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}

	c.Redirect(http.StatusFound, url.LongURL)
}

func (h *URLHandler) Create(c *gin.Context) {
	var url *entity.URL

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
	}

	url, err := h.repo.Create(url.LongURL)
	if err != nil {
		// internal server error
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": url,
	})
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	var url *entity.URL

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
	}

	url, err := h.repo.CreateCustom(url.LongURL, url.ShortURL)
	if err != nil {
		// internal server error
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{"data": url})
}
