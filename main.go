package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.GET("/", defaultHandler)
	r.LoadHTMLGlob("templates/**/*.html")
}

func defaultHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "default.html", gin.H{})
}

func main() {
	r := gin.Default()
	setupRouter(r)
	err := r.Run()
	if err != nil {
		log.Fatalf("gin Run error: %s", err)
	}
}
