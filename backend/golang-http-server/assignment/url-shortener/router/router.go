package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {
	// setup
	r := gin.Default()

	// handler
	r.GET("/:path", urlHandler.Get)            // set handler
	r.POST("/", urlHandler.Create)             // data / json body
	r.POST("/custom", urlHandler.CreateCustom) // data / json body

	// goggle.com => localhost:8080/short
	// google.com => localhost:8080/sdadsad

	// return

	return r
}
