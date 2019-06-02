package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server starts a simple server on the given port
func Server(port int) {
	router := gin.Default()

	loadAllServerResources(router)
	configureRoutes(router)

	fmt.Printf("Started server on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}

func configureRoutes(router *gin.Engine) {

	posts := []map[string]string{
		{"Title": "Hitgub"},
		{"Title": "ItUcket"},
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "landing.gohtml", posts)
	})
}

// loadAllServerResources will load all necessary resources for the server
func loadAllServerResources(router *gin.Engine) {
	loadStaticFiles(router, "./static")
	loadTemplates(router, "templates/*")
}

func loadTemplates(router *gin.Engine, templateDir string) {
	router.LoadHTMLGlob(templateDir)
}

func loadStaticFiles(router *gin.Engine, staticDir string) {
	router.Static("/static", staticDir)
}
