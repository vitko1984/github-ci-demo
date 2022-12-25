package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(r *gin.Engine, db *gorm.DB) {
	r.LoadHTMLGlob("templates/**/*.html")
	r.Use(connectDatabase(db))
	r.GET("/products/", homePageHandler)
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/products/")
	})
}
