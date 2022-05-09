package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// dynamic parameter route
// "/:name"
// "/*name"

// query

func main() {
	// setup
	r := setupRouter()
	r.Run(":8000")

}

func setupRouter() *gin.Engine {
	router := gin.Default()
	// handler
	// root route "/"
	RootRoute(router)

	MainRoute(router)

	// running
	return router
}

func RootRoute(router *gin.Engine) *gin.Engine {
	router.GET("/", HandlerHelloWorld)

	// hello-name/afis/pratama

	// afis
	router.GET("/hello-name/:name", HandlerHello)

	// dynamic routing with wildcard
	// "/afis/pratama/jaya/1"

	router.GET("/hello/*name", HandlerHello)

	router.GET("/hay", HandlerHay)

	return router

}

func MainRoute(router *gin.Engine) *gin.Engine {
	// postForm
	router.POST("/form-post", HandlerFormPost)

	router.GET("/json-output", HandlerJSONStruct)

	// no route / when route not found
	router.NoRoute(HandlerNoRoute)

	v1 := router.Group("/v1")
	{
		subV1 := v1.Group("/sub")
		{
			subV1.GET("/hello", HandlerHelloWorld)
			subV1.GET("/hay", HandlerHay)
		}
	}

	// bindingJSON
	router.POST("/login", HandlerLogin)
	router.POST("/bind-login", HandlerLoginBinding)

	return router
}

type Login struct {
	// tagging : binding
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HandlerLoginBinding(c *gin.Context) {
	var loginUser Login

	if err := c.BindJSON(&loginUser); err != nil {
		// gin error handling
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Bad Request",
			"detail":  "field username & password required",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login SUCCESS",
	})
}

func HandlerLogin(c *gin.Context) {
	// ShouldBindJSON
	var loginUser Login

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		// gin error handling
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Bad Request",
			"detail":  "field username & password required",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login SUCCESS",
	})
}

type JSONOutput struct {
	Name string `json:"name"` // tagging
}

// JSON rendering from struct
func HandlerJSONStruct(c *gin.Context) {
	var output = JSONOutput{
		Name: "Pratama",
	}

	// insert ke c.JSON
	c.JSON(http.StatusOK, output)
}

func HandlerNoRoute(c *gin.Context) {
	// response code harus benar
	c.JSON(http.StatusNotFound, gin.H{
		"status":  "404",
		"message": "Not Found",
	})
}

func HandlerFormPost(c *gin.Context) {
	// PostForm
	msg := c.PostForm("message")
	name := c.DefaultPostForm("name", "Stranger")

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"name":    name,
		"method":  "POST",
	})
}

func HandlerHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

// localhost:8000/hay?firstname=afis&lastname=pratama

func HandlerHay(c *gin.Context) {
	fn := c.Query("firstname")
	ln := c.DefaultQuery("lastname", "Stranger")

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + fn + " " + ln,
	})
}

func HandlerHello(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
	})
}
